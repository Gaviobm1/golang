package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["Linkedin"] = "https://linkedin.com"

	delete(websites, "Google")
	fmt.Println(websites)
}
