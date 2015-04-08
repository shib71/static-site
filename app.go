package main

import (
  "html/template"
  "log"
  "net/http"
  "os"
  "path"
)

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static", fs))

  http.HandleFunc("/", serveTemplate)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("templates", "layout.html")
  fp := path.Join("templates", r.URL.Path)

  info, err := os.Stat(fp)
  if err != nil {
  	if os.IsNotExist(err) {
  		http.NotFound(w, r)
  		return
  	}
  }

  if info.IsDir() {
  	http.NotFound(w, r)
  	return
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
  	log.Println(err.Error())
  	http.Error(w, http.StatusText(500), 500)
  	return
  }

  err = tmpl.ExecuteTemplate(w, "layout", nil)
  if err != nil {
  	log.Println(err.Error())
  	http.Error(w, http.StatusText(500), 500)
  	return
  }
}