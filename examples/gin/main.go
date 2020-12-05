package main

import (
	"fmt"
	boot "github.com/ALiuGuanyan/micro-boot"
	bootflag "github.com/ALiuGuanyan/micro-boot/flag"
	"github.com/ALiuGuanyan/micro-boot/http"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

var cfg config

type config struct {
	HTTP http.HTTP
}

type server struct {
	wg sync.WaitGroup
}

func (s *server) Close() (err error) {
	s.wg.Done()
	return nil
}

func (s *server) Serve() (err error) {
	r := gin.Default()
	r.GET("/", timeout.New(
		timeout.WithTimeout(cfg.HTTP.IdleTimeout),
		timeout.WithHandler(func(c *gin.Context) {
			fmt.Fprintf(c.Writer, "Hello, Gin with Micro-Boot!")
			c.Header("status", "200")
		})),
	)

	go func() {
		s.wg.Add(1)
		r.Run(":" + cfg.HTTP.Port)
	}()

	s.wg.Wait()
	return nil
}

func (c *config) BindFlags(fs *bootflag.FlagSet)  {
	c.HTTP.BindFlags(fs)
}

func (c *config) Initialize(name string) (err error)  {
	return nil
}

func (c *config) Parse() (err error) {
	return nil
}

func main() {
	bt, err := boot.New("gin", &server{},  boot.Root{
		Start: &boot.Config{
			Configurator: &cfg,
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := bt.Execute(); err != nil {
		log.Fatal(err)
		return
	}
}

