package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"AWS":    "https://aws.amazon.com",
	}

	fmt.Println(websites)

	fmt.Println(websites["AWS"])

	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
