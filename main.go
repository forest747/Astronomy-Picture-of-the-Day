package main

import (
	"github.com/forest747/Astronomy-Picture-of-the-Day/rest"
	"log"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8080"))
}
