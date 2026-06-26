package server

import (
	"log"
	"net/http"
)

func Server() {
	if err := loadConfig(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/go", redirectHandler)

	log.Println("Listening on", config.Listen)
	log.Fatal(http.ListenAndServe(config.Listen, mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		ServiceName: config.ServicePortal,
		Services:    config.Services,
		Shares:      config.Shares,
	}

	_ = tpl.Execute(w, data)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	for _, s := range config.Services {
		if s.Name == name {
			http.Redirect(w, r, s.URL, http.StatusFound)
			return
		}
	}

	http.NotFound(w, r)
}
