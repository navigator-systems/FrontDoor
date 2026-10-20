package server

import (
	"bytes"
	"net/http"
	"strings"
)

func rmHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		err := tpl.ExecuteTemplate(w, "remove.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case http.MethodPost:

		name := strings.TrimSpace(r.FormValue("name"))

		if name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		serviceIndex := -1
		for i, service := range config.Services {
			if service.Name == name {
				serviceIndex = i
				break
			}
		}
		if serviceIndex == -1 {
			var page bytes.Buffer
			data := struct {
				Name string
			}{Name: name}
			if err := tpl.ExecuteTemplate(&page, "not-found.html", data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write(page.Bytes())
			return
		}

		config.Services = append(config.Services[:serviceIndex], config.Services[serviceIndex+1:]...)

		if err := saveConfig(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
