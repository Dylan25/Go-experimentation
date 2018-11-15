// fetch prints the content found at each specified URL using io.Copy
// if 'http' is missing, it is appended, the status code is printed before each text portion
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	http_prefix := "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, http_prefix) {
			url = http_prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Status code: %v\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}