// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
//	"io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
	        if !strings.HasPrefix( url, "http://" ) {
		   fmt.Println( "add http:// prefix." )
		   url = "http://" + url
		   //fmt.Println( url)
		}else{
		   fmt.Println( "don't add prefix." )
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		if _, err := io.Copy( os.Stdout, resp.Body ); err != nil {
		   fmt.Println( err )
		}
	}
}

//!-
