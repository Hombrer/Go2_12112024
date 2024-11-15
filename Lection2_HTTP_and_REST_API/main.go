package main

import (
	"fmt"
	"log"
	"net/http"
)

// w - responseWriter (куда писать ответ)
// r - request (откуда брать запрос)
// Обработчик
func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello! I'm new web-server.</h1>")
}


// Функция, которая выбирает нужный обработчик, в зависимости от адреса, по которому пришел запрос
func RequestHandler(){
	http.HandleFunc("/", GetGreet) // Если придет запрос по адресу "/", то вызывай функцию GetGreet
	log.Fatal(http.ListenAndServe(":8080", nil)) // Запускаем веб сервер в режиме "слушания"
}

func main(){
	RequestHandler()
}