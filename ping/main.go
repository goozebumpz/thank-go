package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	pinger := Pinger{client}
	url := "https://ya.ru"
	alive := pinger.Ping(url)
	fmt.Println("is alive =", alive)
}
