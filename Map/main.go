package main

import "fmt"

func main() {
	// colour := make(map[string]string)
	colour := map[string]string{
		"red":   "#FF0000",
		"blue":  "#00FF",
		"green": "#00FF00",
	}
	colour["white"] = "#ffffff"
	colour["black"] = "#000000"

	// delete(colour, "black")
	for key, value := range colour{
		fmt.Println(key+" -> "+value)
	}
}
