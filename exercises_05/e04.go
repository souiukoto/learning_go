package main

import "fmt"

func main() {
	anon := struct {
		level     int
		hits      int
		user      string
		data      map[string]string
		documents []string
	}{
		level: 2,
		hits:  10,
		user:  "marabagre",
		data: map[string]string{
			"image":      "s3.src.com",
			"uploadDate": "12-12-2018",
		},
		documents: []string{
			"base64data",
			"base64data",
		},
	}

	fmt.Println(anon)
}
