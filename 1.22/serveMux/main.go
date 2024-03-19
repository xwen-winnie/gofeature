package main

import (
	"fmt"
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {
	test1()
	//test2()
}

// 模式匹配
func test1() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /eddycjy/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "POST！")
	})
	mux.HandleFunc("GET /eddycjy/update", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "GET! ")
	})
	http.ListenAndServe(":8090", mux)

}

// 通配符
func test2() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "你好！")
	})

	mux.HandleFunc("/eddycjy/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "id 值为 %s", id)
	})
	mux.HandleFunc("/eddycjy/{path...}", func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("path")
		fmt.Fprintf(w, "path 值为 %s", path)
	})
	http.ListenAndServe(":8090", mux)
}
