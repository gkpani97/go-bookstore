package controllers

import(
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gkpani97/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:= models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	bookId:= vars["bookId"]
	ID, err:= strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
    var CreateBook models.Book
    err := decoder.Decode(&CreateBook)
	if err != nil {
        panic(err)
    }
	fmt.Printf("body: %+v",CreateBook)
	book:= CreateBook.CreateBook()
	res, _:= json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	bookId:= vars["bookId"]
	ID, err:= strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	book:= models.DeleteBook(ID)
	res, _:= json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
    var UpdateBook models.Book
    err := decoder.Decode(&UpdateBook)
	if err != nil {
        panic(err)
    }

	vars:= mux.Vars(r)
	bookId:= vars["bookId"]
	ID, err:= strconv.ParseInt(bookId, 0, 0)
	if err != nil{
		fmt.Println(err)
	}

	bookDetails, db:= models.GetBookById(ID)	
	if UpdateBook.Name != ""{
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != ""{
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != ""{
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}