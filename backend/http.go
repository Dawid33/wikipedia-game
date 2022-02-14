package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const PORT = 3333

type commentRequest struct {
	Thread string
}

type Page struct {
	body []byte
}

func startHttpServer() {
	fmt.Println("Starting HTTP Server!")

	mux := http.NewServeMux()
	mux.HandleFunc("/", fileSendHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}

// Function that handles all regular requests
func fileSendHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)
	switch req.Method {
	case "POST":
		switch req.URL.Path {
		case "/api/post_comment":
			name, err := getFieldFromPost("name", w, req)
			comment, err := getFieldFromPost("comment", w, req)
			name = sanitize_text(namePolicy.Sanitize(name))
			comment = sanitize_text(commentPolicy.Sanitize(comment))
			if len(name) > 50 || len(comment) > 500 {
				redirectToUrl(w, req, "/")
				return
			}

			query := "INSERT INTO dawid.guestbook.posts (name, comment) VALUES ($1, $2)"
			_, err = db.Exec(query, name, comment)
			if err != nil {
				fmt.Println(err)
			}

			redirectToUrl(w, req, "/guest_book")
		case "/api/post_blog_comment":
			name, err := getFieldFromPost("name", w, req)
			comment, err := getFieldFromPost("comment", w, req)
			nestingLevel, err := getFieldFromPost("nesting_level", w, req)
			blogId, err := getFieldFromPost("blog_id", w, req)
			redirect, err := getFieldFromPost("redirect", w, req)
			if err != nil {
				break
			}

			rawParent, err := getFieldFromPost("parent", w, req)
			var parent sql.NullString
			if err != nil {
				parent = sql.NullString{Valid: false}
			} else {
				parent = sql.NullString{String: rawParent}
			}

			name = sanitize_text(namePolicy.Sanitize(name))
			comment = sanitize_text(commentPolicy.Sanitize(comment))
			if len(name) > 50 || len(comment) > 500 {
				redirectToUrl(w, req, "/")
				return
			}

			query := "INSERT INTO dawid.blog.comments (name, comment, nesting_level, parent, blog_post_id) VALUES ($1, $2, $3, $4, $5)"
			_, err = db.Exec(query, name, comment, nestingLevel, parent, blogId)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(redirect)
			redirectToUrl(w, req, redirect)

		default:
			fmt.Fprintf(w, "Cannot handle request %s", req.URL.Path)
		}
	case "GET":
		switch req.URL.Path {
		case "/api/guest_list":
			users := GetNewestGuests(db)

			output, err := json.Marshal(users)
			if err != nil {
				fmt.Println(err)
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(output))
		case "/api/blog_comments":
			rawBlogId := req.URL.Query().Get("id")

			blogId, err := strconv.Atoi(rawBlogId[1 : len(rawBlogId)-1]) // remove " before and after the string.
			if err != nil {
				fmt.Println(blogId)
				break
			}

			comments := GetCommentsForBlogPost(db, blogId)
			output, err := json.Marshal(comments)
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

func sanitize_text(input string) string {
	return input
}
