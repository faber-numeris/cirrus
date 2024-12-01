package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	// Final handler
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Request reached the final handler!")
	})

	// Wrap the handler with the middleware
	interceptedHandler := interceptMiddleware(finalHandler)

	// Register the middleware-wrapped handler
	http.Handle("GET /", interceptedHandler)
	http.Handle("POST /", interceptedHandler)
	http.Handle("PUT /", interceptedHandler)

	log.Println("Server is running on http://localhost:4566")
	log.Fatal(http.ListenAndServe(":4566", nil))

}

func interceptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Body
		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("error reading request body: %e", err)
		}
		log.Printf("Request body: %s", body)

		// Get the command sent :
		userAgentMetadata := r.Header.Get("User-Agent")
		values := strings.Split(userAgentMetadata, " ")

		// find the value containing md/command#
		var command string
		for _, value := range values {
			if strings.Contains(value, "md/command#") {
				command = strings.Split(value, "#")[1]
				log.Printf("Command: %s", command)
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)

		// Post-processing: Intercept the response
		log.Println("Response sent!")
	})
}
