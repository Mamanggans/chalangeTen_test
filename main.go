package main

import (
	"go-jwt/handler"
)

type MyString string

var USERNAME = "tes@gmail.com"

var PASSWORD = "123123"

func main() {

	handler.StartApp()
}
