package concurrency

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
	"sync/atomic"
	"time"
)

type Description struct {
	Content string `json:"content"`
}

type Image struct {
	ID          string      `json:"id"`
	Owner       string      `json:"owner"`
	Secret      string      `json:"secret"`
	Server      string      `json:"server"`
	Farm        int         `json:"farm"`
	Title       string      `json:"title"`
	Public      int         `json:"ispublic"`
	Friend      int         `json:"isfriend"`
	Family      int         `json:"isfamily"`
	Description Description `json:"description"`
	Views       string      `json:"views"`
	Tags        string      `json:"tags"`
	MachineTags string      `json:"machine_tags"`
	URL         string      `json:"url_m"`
	Height      int         `json:"height_m"`
	Width       int         `json:"width_m"`
}

func Gofunc() {
	go func() {
		log.Printf("starting go func...")
		time.Sleep(500 * time.Millisecond)
		log.Println("done.")
	}()
	log.Println("main thread")

	time.Sleep(time.Second)
}

func Workers() {
	// slurp in the file
	urlData, err := ioutil.ReadFile("data/flowerurls.json")
	if err != nil {
		log.Fatal(err)
	}

	imageChan := make(chan Image)
	imgDir := "images"
	var imgCount uint64
	workerCounts := make(map[int]int)
	imgWG := sync.WaitGroup{}
	imgMutex := sync.RWMutex{}

	// make sure the image directory exists
	_, err = os.Stat(imgDir)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	if os.IsNotExist(err) {
		err = os.Mkdir(imgDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	// start 8 workers that load the files from their URLs in parallel
	for i := 0; i < 8; i++ {

		// let the WaitGroup know it should wait for one more job to be done
		imgWG.Add(1)

		go func(id int) {
			// iterate over available images in imageChan
			// a chan always returns a unique item and can be read from
			// many go funcs at the same time without mutexes
			for img := range imageChan {

				// shared variables can be read multiple times with an RLock()
				// writes are disabled during RLock()
				imgMutex.RLock()
				log.Printf("[%d] [%d] %s\n", id, imgCount, img.URL)
				imgMutex.RUnlock()

				// parse the URL to extract the path from it
				uri, err := url.Parse(img.URL)
				if err != nil {
					log.Println(err)
					continue
				}

				// retrieve the image via HTTP
				resp, err := http.Get(img.URL)
				if err != nil {
					log.Println(err)
					continue
				}

				// save the image
				imgData, err := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					log.Println(err)
					continue
				}
				err = ioutil.WriteFile(fmt.Sprintf("%s/%s", imgDir, path.Base(uri.Path)), imgData, 0644)
				if err != nil {
					log.Println(err)
				}

				// atomic can be used to modify primitive types thread-safe
				atomic.AddUint64(&imgCount, 1)

				// only modify shared data using a Mutex to prevent race conditions
				// test your binary for race conditions using something like
				// go run -race ./cmd/run -concurrency
				imgMutex.Lock()

				if _, exists := workerCounts[id]; !exists {
					workerCounts[id] = 0
				}
				workerCounts[id] += 1

				imgMutex.Unlock()
			}

			// let the WaitGroup know we're done
			imgWG.Done()
		}(i)
	}

	// parse the file line by line and send each Image to a worker
	var img Image
	scanner := bufio.NewScanner(bytes.NewBuffer(urlData))
	for scanner.Scan() {
		err := json.Unmarshal(scanner.Bytes(), &img)
		if err != nil {
			log.Println(err)
			continue
		}

		// sending the image to the image chan
		// this will block until the previous image has been read
		// from the chan by a worker
		imageChan <- img
	}

	// close the image channel so that the workers terminate
	close(imageChan)

	// wait for all workers to be done
	imgWG.Wait()

	log.Printf("%d images loaded\n", imgCount)
	for id, count := range workerCounts {
		fmt.Printf("worker %d has loaded %d images\n", id, count)
	}
}

func Buffered() {
	rand.Seed(time.Now().UnixNano())

	// create a buffered channel that can hold 2 messages
	ch := make(chan int, 2)

	// create a channel to signal that the work is done
	done := make(chan bool)

	// launch a goroutine that reads messages from the channel and processes them sequentially
	go func() {
		for i := range ch {
			log.Printf("processing %d\n", i)
			time.Sleep(time.Duration(rand.Intn(500) * int(time.Millisecond)))
		}

		// signal that the work is done
		done <- true
	}()

	ticker := time.NewTicker(time.Second)
	seconds := 0
	// try to send 50 messages to the channel
	for i := 0; i < 50; i++ {
		// select tries to read or write from multiple channels
		// the first successful read wins
		select {
		case ch <- i: // try to write to the buffered channel
			log.Printf("sent %d to chan\n", i)
		case <-ticker.C: // try to read from the time ticker
			seconds++
			log.Printf("%d seconds have passed\n", seconds)
		default: // if none of the above worked, default always fires
			log.Printf("failed to send %d to chan\n", i)
		}

		// sleep for a bit to pretend a workload
		time.Sleep(time.Duration(rand.Intn(250) * int(time.Millisecond)))
	}

	// closing the channel makes the goroutine that processes the messages break out of the for loop
	close(ch)

	// stop the ticker or it will keep running forever
	ticker.Stop()

	// wait until the goroutine is done working
	<-done
}

func Concurrency() {
	Gofunc()
	Workers()
}
