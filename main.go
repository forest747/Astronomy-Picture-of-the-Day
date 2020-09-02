package main

import (
	"encoding/json"
	"fmt"
	"github.com/GiterLab/urllib"
)

type Info struct {
	Date           string
	Explanation    string
	MediaType      string
	ServiceVersion string
	Title          string
	Url            string
}

func main() {
	req := urllib.Get("https://api.nasa.gov/planetary/apod?api_key=rvU2JWqSHNFizqfke1599aJG4Ax3GvKmQYXPfSld")
	req.Debug(true)
	strJson, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	var info Info

	err = json.Unmarshal([]byte(strJson), &info)
	if err != nil {
		fmt.Println("Parsing error!!")
	}

	fmt.Printf("%s",info.Date)
}
