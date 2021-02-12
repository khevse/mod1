package main

import (
	"log"

	"github.com/khevse/mod2"
)

func main() {
	log.Println("version:", mod2.New().String())
}
