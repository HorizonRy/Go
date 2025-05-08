package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  float64  `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title:  "Inception",
		Year:   2010,
		Price:  19.99,
		Actors: []string{"Leonardo DiCaprio", "Joseph Gordon-Levitt"},
	}

	// 结构体编码为json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
		return
	}

	fmt.Printf("JSON string:%s\n", jsonStr)

	// json解码为结构体
	var myMovie Movie
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Decoded movie: %v\n", myMovie)
}
