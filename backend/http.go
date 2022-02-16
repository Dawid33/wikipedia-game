package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const PORT = 3333

func startHttpServer() {
	fmt.Println("Starting HTTP Server!")

	mux := http.NewServeMux()
	mux.HandleFunc("/", requestHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}

// Function that handles all regular requests
func requestHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)
	switch req.Method {
	case "POST":
		switch req.URL.Path {
		case "/api/register_session":
		case "/api/register_user":
		case "/api/start_session":
		default:
			fmt.Fprintf(w, "Cannot handle request %s", req.URL.Path)
		}
	case "GET":
		switch req.URL.Path {
		case "/api/active_sessions":
			sessions := GetActiveSessions(db)
			output, err := json.Marshal(sessions)
			if err != nil {
				fmt.Println(err)
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(output))
		default:
			fmt.Fprintf(w, "Cannot handle request %s", req.URL.Path)
		}
	default:
		fmt.Fprintf(w, "Cannot handle method %s", req.Method)
	}
}

func getFieldFromPost(field string, w http.ResponseWriter, req *http.Request) (string, error) {
	text := req.FormValue(field)
	if text == "" {
		err := errors.New(fmt.Sprintf("%s field does not exist in post form", field))
		return "", err
	}
	return text, nil
}

func redirectToUrl(w http.ResponseWriter, req *http.Request, url string) {
	http.Redirect(w, req, url, http.StatusMovedPermanently)
}
