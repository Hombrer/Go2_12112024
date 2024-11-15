package handlers

import (
	"BooksAPI/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


// GetBookById function returns book by id
func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "don't use ID parametr as uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	book, ok := models.FindBookByID(id)
	log.Println("Get book with id:", id)
	if !ok {
		writer.WriteHeader(http.StatusNotFound) // 404 error
		message := models.Message{Message: "book with that ID does not exist in database."}
		json.NewEncoder(writer).Encode(message)
	} else {
		writer.WriteHeader(http.StatusOK) // 200
		json.NewEncoder(writer).Encode(book)
	}
}

// CreateBook function creates new book an saves in database.
func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new book ...")
	var book models.Book

	// Check json file
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&book)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "provided json file is invalid."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	// Create new book and save one
	var NewBookID int = len(models.DB) + 1
	book.ID = NewBookID
	models.DB = append(models.DB, book)

	writer.WriteHeader(http.StatusCreated) // status code: 201
	json.NewEncoder(writer).Encode(book)

}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating book ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "don't use ID parametr as uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	// Check id 
	_, ok := models.FindBookByID(id)
	if !ok {
		writer.WriteHeader(http.StatusNotFound) // 404 error
		message := models.Message{Message: "book with that ID does not exist in database."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	var newBook models.Book

	// Check json file
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&newBook)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "provided json file is invalid."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	
	// Нужно заменит oldBook на newBook в DB
	res := models.UpdateBookById(id, newBook)

	if !res {
		writer.WriteHeader(http.StatusNoContent) // status code: 204
		return 
	}

	writer.WriteHeader(http.StatusOK) // status code: 200
	message := models.Message{Message: "book has successfully changed."}
	json.NewEncoder(writer).Encode(message)

}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest) // 400 error
		message := models.Message{Message: "don't use ID parametr as uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	_, ok := models.FindBookByID(id)
	
	if !ok {
		log.Println("book with id =", id, "not found.")
		writer.WriteHeader(http.StatusNotFound) // 404 error
		message := models.Message{Message: "book with that ID does not exist in database."}
		json.NewEncoder(writer).Encode(message)
		return
	} 
	
	// Delete book
	models.DeleteBookById(id)
	message := models.Message{Message: "book has successfully deleted from database."}
	json.NewEncoder(writer).Encode(message)
}
