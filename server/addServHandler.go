package server

import (
	"net/http"
	"strings"
)

func addHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		err := tpl.ExecuteTemplate(w, "add.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case http.MethodPost:

		name := strings.TrimSpace(r.FormValue("name"))
		url := strings.TrimSpace(r.FormValue("url"))

		if name == "" || url == "" {
			http.Error(w, "Name and URL are required", http.StatusBadRequest)
			return
		}

		config.Services = append(config.Services, Service{
			Name: name,
			URL:  url,
		})

		if err := saveConfig(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
