package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) errorPage(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	log.Printf("%d - %s", status, msg)

	data := struct {
		Status  int
		Message string
	}{
		Status:  status,
		Message: http.StatusText(status),
	}

	tmpl, err := template.ParseFiles("web/template/error.html")
	if err != nil {
		fmt.Fprintf(w, "%d - %s\n", data.Status, data.Message)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "error.html", data); err != nil {
		fmt.Fprintf(w, "%d - %s\n", data.Status, data.Message)
	}
}
