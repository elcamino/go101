package httpea

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
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
	var r Foo
	gofakeit.Struct(&r)

	c.JSON(http.StatusOK, r)
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

func Http() {
	// create a webserver
	//
	srv := New(":9779")
	go srv.Run()

	time.Sleep(500 * time.Millisecond)

	// query the server
	//
	resp, err := http.Get("http://127.0.0.1:9779/rand")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var f Foo

	err = json.Unmarshal(body, &f)
	if err != nil {
		log.Fatal(err)
	}

	pretty.Println(f)
}
