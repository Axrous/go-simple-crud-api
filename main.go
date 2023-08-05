package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	fmt.Println("REST API v2.0 Mux Routers")
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequest()
}

func handleRequest()  {
	myRouter := mux.NewRouter().StrictSlash(true)

	// http.HandleFunc("/", homePage)
	// http.HandleFunc("/articles", getAllArticles)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", getAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", getArticleById)

	defer log.Fatal(http.ListenAndServe("localhost:8000", myRouter))
	
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")

}

func getAllArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: getAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func getArticleById(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
		}
	}
	
}

func createNewArticle(w http.ResponseWriter, r *http.Request)  {
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))
	

	var article Article

	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for i, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:i], Articles[i+1:]... )
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]

	var newArticle Article

	reqBody, err :=  ioutil.ReadAll(r.Body)

	if err !=nil {
		panic(err)
	}

	json.Unmarshal(reqBody, &newArticle)

	for i, article := range Articles {
		
		if article.Id == id {
			article.Id = newArticle.Id
			article.Title = newArticle.Title
			article.Desc = newArticle.Desc
			article.Content = newArticle.Content
		}

		Articles = append(Articles[:i], article)
	}

	json.NewEncoder(w).Encode(Articles)
}