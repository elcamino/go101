package contextdemo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type Foo struct {
	Bar           string
	Int           int
	Pointer       *int
	Name          string    `fake:"{firstname}"`  // Any available function all lowercase
	Sentence      string    `fake:"{sentence:3}"` // Can call with parameters
	RandStr       string    `fake:"{randomstring:[hello,world]}"`
	Number        string    `fake:"{number:1,10}"`       // Comma separated for multiple values
	Regex         string    `fake:"{regex:[abcdef]{5}}"` // Generate string from regex
	Skip          *string   `fake:"skip"`                // Set to "skip" to not generate data for
	Created       time.Time // Can take in a fake tag as well as a format tag
	CreatedFormat time.Time `fake:"{year}-{month}-{day}" format:"2006-01-02"`
}

type Server struct {
	addr   string
	router *gin.Engine
}

func (s *Server) getRandom(c *gin.Context) {
	log.Println("random handler started")
	defer log.Println("random handler ended")

	select {
	case <-time.After(5 * time.Second):
		var r Foo
		gofakeit.Struct(&r)

		c.JSON(http.StatusOK, r)
	case <-c.Request.Context().Done():
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "request aborted"})
	}
}

func (s *Server) Run() {
	endless.ListenAndServe(s.addr, s.router)
}

func New(addr string) *Server {
	router := gin.Default()
	s := Server{
		addr:   addr,
		router: router,
	}

	router.GET("/rand", s.getRandom)

	return &s
}

func Context() {
	// create a webserver
	//
	srv := New(":9779")
	go srv.Run()

	time.Sleep(500 * time.Millisecond)

	// cancel the request after 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// the server takes 5 seconds to answer the request
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:9779/rand", nil)
	client := &http.Client{}

	println("Do")
	resp, err := client.Do(req)
	println("Do finished")

	if err != nil {
		log.Printf("http get error: %s\n", err) // cancel caught
	} else {
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
	}

	time.Sleep(300 * time.Millisecond)
}
