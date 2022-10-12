package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v4"
)

// define vars
var Port string = "3000"
var HttpPort = fmt.Sprintf(":%s", Port)
var mySigningKey = []byte("your-256-bit-secret")

// http handler
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Content protected by JWT Auth!"))
}

// middleware for parsing HTTP Token Header from incoming requests
func jwtTokenParser(endpoint http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middleware")
		// verify if Token header exists
		headers := r.Header
		_, exists := headers["Token"]
		if exists {
			tokenString := r.Header.Get("Token")
			// validate token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Println(err)
			}
			// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 	fmt.Println(claims)
			if token.Valid {
				log.Print("JWT Auth is successful!")
				endpoint.ServeHTTP(w, r)
			} else {
				log.Print("JWT Auth Token is NOT valid!")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Not Authorised!\nJWT Auth Token is NOT valid!"))
			}
		} else {
			log.Print("JWT Auth Token is NOT Present!")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorised!\nJWT Auth Token is NOT Present!"))
		}
	})
}

// http server config and mux router
func handleRequests() {
	fmt.Println("Starting server on port", Port)
	mux := http.NewServeMux()
	mux.Handle("/", jwtTokenParser(homePage))
	err := http.ListenAndServe(HttpPort, mux)
	log.Fatal(err)
}

func main() {
	handleRequests()
}
