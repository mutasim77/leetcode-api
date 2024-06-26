package api

import (
	"fmt"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `<h1 style="text-align: center;">Welcome to LeetCode REST API Wrapper 🚀</h1>`)
	fmt.Fprintf(w, `<h3 style="text-align: center;">Explore LeetCode user information by visiting: <a href="/user/example">/user/example</a></h3>`)
}
