package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// encrypt token dari username menggunakan bcrypt kemudian kirim ke user kedalam cookie
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		cookieName := "token"
		var creds Credentials

		// Task:  Buat JSON body diconvert menjadi credential struct & return bad request ketika terjadi kesalahan decoding json
		// TODO: answer here
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Task: Ambil password dari username yang dipakai untuk login
		// TODO: answer here
		password := users[creds.Username]

		// Task: return unauthorized jika password salah
		// TODO: answer here
		if password != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
		}

		// Task: 1. Buat token string menggunakan bcrypt dari credential username
		// 		 2. return internal server error ketika terjadi kesalahan encrypting token
		// TODO: answer here

		tokenString, err := bcrypt.GenerateFromPassword([]byte(creds.Username), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Task: Set token string kedalam cookie response
		// TODO: answer here
		cookie := http.Cookie{
			Name:    cookieName,
			Value:   string(tokenString),
			Expires: time.Now().Add(5 * time.Minute),
		}
		http.SetCookie(w, &cookie)
	})

	return mux
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
