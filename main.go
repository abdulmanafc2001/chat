package main

import "net/http"

func main() {
	apiHandler()
	http.ListenAndServe(":5000", nil)
}

func apiHandler() {
	manager := newManager()
	http.Handle("/", http.FileServer(http.Dir("./frontent")))
	http.Handle("/ws", manager.serveWS())
}
