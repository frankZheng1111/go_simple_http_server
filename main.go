package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	respJson := map[string]interface{}{
		"code": 1,
		"msg":  "Hello world",
	}
	for key, params := range queryParams {
		respJson[key] = params
	}
	resp, err := json.Marshal(respJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(resp))

}

func main() {
	http.HandleFunc("/hello", getHandler)
	http.ListenAndServe(":8080", nil)
}
