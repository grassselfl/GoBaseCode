package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	DOMAIN := "http://127.0.0.1:8080"
	response, err := http.Get(DOMAIN + "/api/hello")
	if err != nil {
		return
	}
	fmt.Println(response)
	fmt.Println(response.Body)
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(all))
}

func xx() {

	api := "http://127.0.0.1:8080/api/hello"
	uri, err := url.ParseRequestURI(api)
	if err != nil {
		return
	}

	values := url.Values{}
	values.Add("name", "Hello World")
	uri.RawQuery = values.Encode()

	request, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return
	}

	fmt.Println(request)

	response, err := http.DefaultClient.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	if err != nil {
		return
	}
	fmt.Println(response)

}
