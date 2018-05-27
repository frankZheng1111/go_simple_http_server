package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	respJson := map[string]interface{}{
		"code": 1,
		"msg":  "Hello world",
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
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
