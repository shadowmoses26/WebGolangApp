package main

import (
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Показать все сниппеты"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Метод запрещен", 405)
		return

	}

	w.Write([]byte("Создать сниппет"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запускаем сервер: http://127.0.0.1:4000 ")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
