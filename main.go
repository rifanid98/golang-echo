package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

type myHtmlHandler struct{}

func (h *myHtmlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h1>Hello World</h1>
		</body>
		</html>
	`)
}

func main() {
	// http.ListenAndServe("localhost:8080", &myHandler{})
	http.ListenAndServe("localhost:8080", &myHtmlHandler{})
}
