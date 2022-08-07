// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"text/template"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type M map[string]interface{}

// func handlerIndex(w http.ResponseWriter, r *http.Request) {
// 	var data = M{"name": "Batman"}
// 	var tmpl = template.Must(template.ParseFiles(
// 		"views/content/listTask.html",
// 		"views/template/_header.html",
// 		"views/template/_navbar.html",
// 		"views/template/_footer.html",
// 	))

// 	var err = tmpl.ExecuteTemplate(w, "index", data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func connect() (*sql.DB, error) {
// 	db, err := sql.Open("MySql", "root:@tcp(127:0.0.1:3306)/go")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

// func main() {
// 	http.Handle("/static/",
// 		http.StripPrefix("/static/",
// 			http.FileServer(http.Dir("assets"))))
// 	http.HandleFunc("/", handlerIndex)
// 	http.HandleFunc("/index", handlerIndex)

// 	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("hello again"))
// 	})

// 	var address = "localhost:9000"
// 	fmt.Printf("server started at %s\n", address)
// 	err := http.ListenAndServe(address, nil)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
