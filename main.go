package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	flag "gopkg.in/alecthomas/kingpin.v2"
)
// If you want to run the web server on a different port you can send it
// As argument
var (
	Port = flag.Flag("port", "Please specify application port").Default("8080").String()
)

func main() {

	flag.Parse()

	router := gin.Default()
	router.POST("/api/:schema", postHandler)
	log.Fatal(router.Run(fmt.Sprintf(":%s", *Port)))
}
