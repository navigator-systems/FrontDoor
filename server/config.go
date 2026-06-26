package server

import (
	"encoding/json"
	"os"
	"text/template"
)

type Service struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Share struct {
	Name        string `json:"name"`
	Info        string `json:"info"`
	Instruction string `json:"instruction"`
}

type Config struct {
	Listen   string    `json:"listen"`
	Services []Service `json:"services"`
	Shares   []Share   `json:"shares"`
}

type PageData struct {
	Services []Service
	Shares   []Share
}

var config Config
var tpl = template.Must(template.ParseFiles("templates/index.html"))

func loadConfig() error {
	data, err := os.ReadFile("services.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &config)
}
