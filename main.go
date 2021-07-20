package main

import (
	"fmt"
	"net/http"
)

func handlerHelloWorld(w http.ResponseWriter, r *http.Request) {
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

func handlerWelcome(w http.ResponseWriter, r *http.Request) {
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
			<h1>Welcome!</h1>
		</body>
		</html>
	`)
}

func main() {
	// receive functions that have http.Responewriter parameters and *http.request
	http.HandleFunc("/", handlerHelloWorld)
	// menerima http.HandlerFunc() atau variable/struct yang mengimplementasikan interface :
	// type HandlerFunc func(ResponseWriter, *Request)
	// ServeHTTP calls f(w, r).
	// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	// 	f(w, r)
	// }
	http.Handle("/welcome", http.HandlerFunc(handlerWelcome))
	http.ListenAndServe("localhost:8080", nil)
}
