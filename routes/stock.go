package routers

import (
	"bookstore_case/models"
	"bookstore_case/pkg"
	"bookstore_case/pkg/helper"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func AddStock(response http.ResponseWriter, r *http.Request) {
	var u models.JwtUserAuth
	var book models.Book

	user := r.Context().Value("data")
	tmp, _ := json.Marshal(user)
	json.Unmarshal(tmp, &u)

	json.NewDecoder(r.Body).Decode(&book)
	validate := validator.New()
	errVal := validate.Struct(book)
	if errVal != nil {
		helper.HTTPErrorHandler(response, "Book fields is empty or types wrong!", http.StatusNotAcceptable)
		return
	}

	pkg.Mongo.Insert(pkg.DbClient, book, "books")
	fmt.Fprintf(response, "Hello %s", u.Email, "Book added!->", book.Title)
}

func UpdateStock(response http.ResponseWriter, r *http.Request) {
	var book models.Book
	var u models.JwtUserAuth

	user := r.Context().Value("data")
	tmp, _ := json.Marshal(user)
	json.Unmarshal(tmp, &u)

	json.NewDecoder(r.Body).Decode(&book)
	validate := validator.New()
	errVal := validate.Struct(book)
	if errVal != nil {
		helper.HTTPErrorHandler(response, "Book fields is empty or types wrong!", http.StatusNotAcceptable)
		return
	}

	pkg.Mongo.UpdateBook(pkg.DbClient, book)
	fmt.Fprintf(response, "Hello %s", u.Email, "Book updated!->", book.Title)
}

func DeleteStock(response http.ResponseWriter, r *http.Request) {
	var book models.Book
	var u models.JwtUserAuth

	user := r.Context().Value("data")
	tmp, _ := json.Marshal(user)
	json.Unmarshal(tmp, &u)
	json.NewDecoder(r.Body).Decode(&book)

	pkg.Mongo.DeleteBook(pkg.DbClient, r.URL.Query().Get("id"))
	fmt.Fprintf(response, "Hello %s", u.Email, "Book deleted!->", book.ID)
}
