package main

import (
    "encoding/json"
    "github.com/gorilla/pat"
    "gopkg.in/mgo.v2"
    "net/http"
)

var (
    sessionMongo *mgo.Session
)

const (
    dbName            = "apiTest"
    articleCollection = "articles"
)

func init() {
    var err error
    if sessionMongo, err = mgo.Dial("localhost"); err != nil {
        panic(err)
    }
}

type Article struct {
    Titre       string `json:"title" bson:"title"`
    Description string `json:"description" bson:"description"`
    Author      string `json:"author,omitempty"`
}

func GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
    articles := make([]*Article, 0)
    if err := sessionMongo.DB(dbName).C(articleCollection).Find(nil).All(&articles); err != nil {
        w.Write([]byte(err.Error()))
        return
    }
    jsonData, err := json.MarshalIndent(articles, "", " ")
    if err != nil {
        w.Write([]byte(err.Error()))
        return
    }
    w.Write(jsonData)
}

func AddArticleHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
    router := pat.New()
    router.Get("/articles", GetArticlesHandler)
    router.Post("/articles", AddArticleHandler)
    http.ListenAndServe(":8080", router)
}
