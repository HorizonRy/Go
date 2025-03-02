package main

import (
	"log"
	"os"

	"github.com/HorizonRy/go_ws/GoInAction/Ch2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
