package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type LoginSuccessResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

// Jwt key yang akan dipakai untuk membuat signature
var jwtKey = []byte("key")
var jwtCookieKey = "token"

// Struct claim digunakan sebagai object yang akan di encode oleh jwt
// jwt.StandardClaims ditambahkan sebagai embedded type untuk provide standart claim yang biasanya ada pada JWT
type Claims struct {
	Username string
	jwt.StandardClaims
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")
	res, err := api.usersRepo.Login(username, password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	// Task: 1. Deklarasi expiry time untuk token jwt
	//       2. Buat claim menggunakan variable yang sudah didefinisikan diatas
	//       3. expiry time menggunakan time millisecond
	// Deklarasi expiry time
	expirationTime := time.Now().Add(time.Millisecond)

	// Buat claim
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// expiry time menggunakan time millisecond
			ExpiresAt: expirationTime.UnixMilli(),
		},
	}

	//json.NewEncoder(w).Encode(LoginSuccessResponse{Username: *res}) // TODO: replace this
	// TODO: answer here

	// Task: Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai
	// TODO: answer here
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Task: 1. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
	//       2. return internal error ketika ada kesalahan ketika pembuatan JWT string
	// TODO: answer here
	// Buat jwt string dari token
	tokenString, _ := token.SignedString(jwtKey)
	if err != nil {
		// return internal error ketika ada kesalahan ketika pembuatan JWT string
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Task: Set token string kedalam cookie response
	// TODO: answer here
	http.SetCookie(w, &http.Cookie{
		Name:    jwtCookieKey,
		Value:   tokenString,
		Expires: expirationTime,
	})

	// Task: Return response berupa username dan token JWT yang sudah login
	//json.NewEncoder(w).Encode(LoginSuccessResponse{Username: *res, Token: tokenString}) // TODO: replace this
	encoder.Encode(LoginSuccessResponse{Username: *res, Token: tokenString}) // TODO: replace this
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	username := req.URL.Query().Get("username")
	err := api.usersRepo.Logout(username)

	w.Header().Set("Content-Type", "application/json")
	//encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder := json.NewEncoder(w)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	//encoder.Encode(AuthErrorResponse{Error: ""}) // TODO: replace this
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(AuthErrorResponse{Error: username}) // TODO: replace this
}
