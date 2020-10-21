package main

import (
    "net/http"
		"html/template"
)

type KittensPageData struct {
    KittensPageTitle string
}

func main() {
    tmpl := template.Must(template.ParseFiles("template.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := KittensPageData{
            KittensPageTitle: "My kittens webpage!!",
        }
        tmpl.Execute(w, data)
    })

    fs := http.FileServer(http.Dir("img/"))
    http.Handle("/img/", http.StripPrefix("/img/", fs))

    http.ListenAndServe(":8080", nil)
}
