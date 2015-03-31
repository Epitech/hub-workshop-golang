package main

import (
    "github.com/gorilla/pat"
    "net/http"
    "encoding/json"
)
type Article struct {
    Titre       string
    Description string
}
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
    r.Form.Get()
    article1 := &Article{
        Titre:       "C'est la fin du monde",
        Description: "C'est une courte description de l'article",
    }
    article2 := &Article{
        Titre:       "C'est plus la fin du monde",
        Description: "C'est aussi une courte description",
    }
    articles := make([]*Article, 2)
    articles[0] = article1
    articles[1] = article2
    jsonArticles, err := json.MarshalIndent(articles, "", "  ")
    if err != nil {
        w.Write([]byte("Erreur"))
        return
    }
    w.Write(jsonArticles)
}

func main() {
    router := pat.New()
    router.Get("/articles", ArticlesHandler)
    http.ListenAndServe(":8080", router)
}
