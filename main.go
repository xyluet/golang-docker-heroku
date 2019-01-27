package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	foo := os.Getenv("FOO")

	http.ListenAndServe(fmt.Sprintf(":%s", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"port": port,
			"foo":  foo,
		})
	}))
}
