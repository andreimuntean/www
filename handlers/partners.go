package handlers

import "net/http"

func partners(w http.ResponseWriter, r *http.Request) {
	templates["comingsoon"].ExecuteTemplate(w, "layout", nil)
}