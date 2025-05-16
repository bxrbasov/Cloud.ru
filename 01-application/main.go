package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

type PageData struct {
	Hostname    string
	IPAddresses []net.IP
	Author      string
}

func main() {
	http.HandleFunc("/", getWebPage)

	fmt.Println("Сервер запущен и слушает на порту 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getWebPage(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка получения имени хоста: %s", err), http.StatusInternalServerError)
		return
	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка получения IP адреса: %s", err), http.StatusInternalServerError)
		return
	}

	author := os.Getenv("AUTHOR")
	if author == "" {
		author = "Неизвестный автор"
	}

	data := PageData{
		Hostname:    hostname,
		IPAddresses: addrs,
		Author:      author,
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка загрузки шаблона: %s", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка выполнения шаблона: %s", err), http.StatusInternalServerError)
		return
	}
}
