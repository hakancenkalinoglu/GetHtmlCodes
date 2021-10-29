package main

import (
	"fmt"
	"io/ioutil" // provides I/O utility functions.
	"net/http" // Package http provides HTTP client and server implementations.

)
// we are checking errors
func check(err error){
	if err != nil {
		panic(err)
	}
}

func main() {

	url := "https://hakancenkalinoglu.github.io/hakancenk_html/"
	fmt.Printf("HTML code of %s ...\n", url)
	// response the web site 
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html) 


	// we are writing html codes to txt 
	data := []byte(html)
	err = ioutil.WriteFile("html_codes.txt", data, 0644)
	check(err)
}