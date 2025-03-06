package main

import (
	"github.com/noahzyl/gin-ranking/routers"
)

// Test gin
func main() {
	r := routers.Router()
	r.Run(":9999") // Listen on localhost:9999
}
