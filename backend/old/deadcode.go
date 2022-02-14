package old

/*
func serveIndexFile(w http.ResponseWriter, req *http.Request) {
	db := connectToDB()
	rows, err := db.Query("SELECT * FROM forum.posts")

	var prefetchLinks string
	for rows.Next() {
		var threadId uint64
		var userId string
		var category string
		var title string
		var content string
		err = rows.Scan(&threadId, &userId, &category, &title, &content)
		CheckError(err)
		//newThreads += fmt.Sprintf(, title, threadId)
		prefetchLinks += fmt.Sprintf("	<link rel=\"prefetch\" href=\"thread?%d\">\n", threadId)
	}
	doc := getTemplateFile("index.html")
	/*
	addContentToTagByIdInDoc(doc, "thread-form", getSubTemplateFile("thread_form.html"))
	addContentToTagByIdInDoc(doc, "posts", newThreads)
	addContentToTagInDoc(doc, "head", prefetchLinks)

	_, err = fmt.Fprintf(w, htmlToString(doc))
	CheckError(err)
}

func serveThreadFile (w http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Has("id") {

			id := req.URL.Query().Get("id")
			post, err := getThread(id)
			if err != nil {
				redirectToUrl(w, req, "404.html")
			}
			comments, err := getCommentsInThread(connectToDB(), id)

		doc := getTemplateFile("thread.html")


		addContentToTagByIdInDoc(doc, "comment-form", fmt.Sprintf(getSubTemplateFile("comment_form.html"), post.threadId, 0, post.threadId))
		addContentToTagByIdInDoc(doc, "post-title", post.title)
		addContentToTagByIdInDoc(doc, "post-content", post.content)
		addContentToTagByIdInDoc(doc, "comments", comments)

		fmt.Fprintf(w, htmlToString(doc))
	} else {
		redirectToUrl(w, req, "404.html")
	}
}

func getThreads(db *sql.DB, whereClause string) ([]Thread, error) {
	query := "SELECT * FROM forum.posts " + whereClause + ";"

	rows, err := db.Query(query)
	if err != nil {
		PrintError(err)
		return []Thread{}, nil
	}
	var output []Thread
	for rows.Next() {
		var postid uint64
		var userid string
		var category string
		var title string
		var content string
		err = rows.Scan(&postid, &userid, &category, &title, &content)
		if err != nil {
			return []Thread{}, nil
		}
		output = append(output, Thread{
			threadId: postid,
			userId:   userid,
			category: category,
			title:    title,
			content:  content,
		})
	}
	return output, nil
}

func getComments(db *sql.DB, whereClause string) ([]Comment, error) {
	query := "SELECT * FROM forum.comments " + whereClause + ";"
	rows, err := db.Query(query)
	if err != nil {
		return []Comment{}, err
	}
	var output []Comment
	for rows.Next() {
		var commentId string
		var threadId int
		var parentId int
		var kidsId []int
		var userId string
		var content string
		err = rows.Scan(&commentId, &threadId, &parentId, pq.Array(&kidsId), &userId, &content)
		if err != nil {
			return []Comment{}, err
		}
		output = append(output, Comment{
			commentId: commentId,
			threadId:  threadId,
			parentId:  parentId,
			kidsId:    kidsId,
			userId:    userId,
			content:   content,
		})
	}
	return output, nil
}
*/

/*
func redirectTo503OnError(w http.ResponseWriter, req *http.Request, err error) {
	if err != nil {
		PrintError(err)
		redirectToUrl(w, req, "404.html")
	}
}

// TODO: This code needs to be rewritten and done properly.
func isStaticFile(w http.ResponseWriter, req *http.Request) (bool, string) {
	if strings.HasSuffix(req.URL.Path, ".css") {
		w.Header().Set("Content-Type", "text/css")
		return true, "./public" + req.URL.Path
	}
	if strings.HasSuffix(req.URL.Path, ".js") {
		w.Header().Set("Content-Type", "application/javascript")
		return true, "./public" + req.URL.Path
	}
	if strings.HasSuffix(req.URL.Path, ".ico") {
		return true, "./public" + req.URL.Path
	}
	if strings.HasSuffix(req.URL.Path, ".html") {
		return true, "./public" + req.URL.Path
	}
	if strings.HasSuffix(req.URL.Path, ".js.map") {
		return true, "./public" + req.URL.Path
	}
	return false, ""
}

func sendGeneratedHtmlFile(w http.ResponseWriter, req *http.Request) {

}

func acceptNewThread(db *sql.DB, w http.ResponseWriter, req *http.Request) {
	text, err := getFieldFromPost("text", w, req)
	if err != nil {
		redirectToUrl(w, req, req.URL.RawPath+req.URL.RawQuery)
		return
	}
	title, err := getFieldFromPost("title", w, req)
	if err != nil {
		redirectToUrl(w, req, req.URL.RawPath+req.URL.RawQuery)
		return
	}

	_, err = db.Exec("INSERT INTO forum.posts (userid, category, title, content) VALUES ($1, $2, $3, $4);", "Anonymous", "default", title, text)
	if err != nil {
		PrintError(err)
		redirectToUrl(w, req, req.URL.RawPath+req.URL.RawQuery)
		return
	}

	if gotoUrl := req.FormValue("goto"); gotoUrl == "" {
		redirectToUrl(w, req, req.URL.RawPath+req.URL.RawQuery)
		return
	}
	redirectToUrl(w, req, req.FormValue("goto"))
}

func acceptNewComment(db *sql.DB, w http.ResponseWriter, req *http.Request) {
	thread, err := getFieldFromPost("thread", w, req)
	if err != nil {
		redirectToUrl(w, req, "/")
		return
	}

	content, err := getFieldFromPost("text", w, req)
	if err != nil {
		redirectToUrl(w, req, "/thread?threadid="+thread)
		return
	}

	parentId, err := getFieldFromPost("parentId", w, req)
	if err != nil {
		redirectToUrl(w, req, req.URL.RawPath+req.URL.RawQuery)
		return
	}

	threadId, err := strconv.ParseInt(thread, 10, 32)
	if err != nil {
		redirectToUrl(w, req, req.URL.RawPath+req.URL.RawQuery)
		return
	}

	_, err = db.Exec("INSERT INTO forum.comments (threadid, parentid, kidsid, userid, content) VALUES ($1, $2, $3, $4, $5);", threadId, parentId, nil, "Anonymous", content)
	if err != nil {
		redirectToUrl(w, req, req.URL.RawPath+req.URL.RawQuery)
		return
	}

	if gotoUrl := req.FormValue("goto"); gotoUrl == "" {
		redirectToUrl(w, req, req.URL.RawPath+req.URL.RawQuery)
		return
	}
	redirectToUrl(w, req, req.FormValue("goto"))
}
*/

//func getTemplateFile(file string) (*html.Node, error) {
//	content, err := ioutil.ReadFile("./public/" + file)
//	if err != nil {
//		return nil, err
//	}
//	doc, err := html.Parse(bytes.NewReader(content))
//	if err != nil {
//		return nil, err
//	}
//	return doc, nil
//}
//
//func getSubTemplateFile(file string) string {
//	doc, err := getTemplateFile(file)
//	Panic(err)
//	// Parsing html adds stuff that doesn't exist in the file.
//	doc, err = getNodeByTag(doc, "body")
//	Panic(err)
//	return getContentFromNode(doc)
//}
