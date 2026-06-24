package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

var urls = make(map[string]string)

const fileName = "urls.json"


func loadURLs() {

	file, err := os.ReadFile(fileName)

	if err != nil {
		return
	}

	json.Unmarshal(file, &urls)
}



func saveURLs() {

	data, _ := json.MarshalIndent(urls, "", "  ")

	os.WriteFile(fileName, data, 0644)

}



func generateShortCode() string {

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := ""

	for i := 0; i < 6; i++ {

		code += string(letters[rand.Intn(len(letters))])

	}

	return code
}



func shortenURL(w http.ResponseWriter, r *http.Request) {


	longURL := r.FormValue("url")


	shortCode := generateShortCode()


	urls[shortCode] = longURL


	saveURLs()


	response := map[string]string{

		"short_url":"http://localhost:8080/" + shortCode,
	}


	json.NewEncoder(w).Encode(response)

}



func redirectURL(w http.ResponseWriter, r *http.Request) {


	code := r.URL.Path[1:]


	longURL, exists := urls[code]


	if !exists {

		http.NotFound(w,r)

		return
	}


	http.Redirect(w,r,longURL,http.StatusFound)

}



func main(){


	loadURLs()


	http.HandleFunc("/shorten",shortenURL)


	http.HandleFunc("/",redirectURL)


	fmt.Println("URL Shortener running on port 8080")


	http.ListenAndServe(":8080",nil)

}