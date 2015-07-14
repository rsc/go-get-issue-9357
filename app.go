package app

import "net/http"

func init() {
	http.HandleFunc("/", meta)
}

const metaPage = `<html>
<meta name="go-import" content="go-get-issue-9357.appspot.com git https://github.com/rsc/go-get-issue-9357">
`

func meta(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(metaPage))
}
