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
	http.HandleFunc("/", handlerHelloWorld)
	http.HandleFunc("/welcome", handlerWelcome)
	http.ListenAndServe("localhost:8080", nil)
}
