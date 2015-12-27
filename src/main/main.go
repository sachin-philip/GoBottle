package main


import (
	"net/http"
	"io/ioutil"
	"strings"
	"log"
	"text/template"
	)

// these are runservers

type urls struct{

}

func (this *urls) ServeHTTP(w http.ResponseWriter, r *http.Request){

	path := r.URL.Path[1:]
	log.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css"){
			contentType = "text/css"
		} else if strings.HasSuffix(path, "js"){
			contentType = "text/js"
		} else if strings.HasSuffix(path, "png"){
			contentType = "image/png"
		} else{
			contentType = "text/plain"
		}

		w.Header().Add("Content Type", contentType)
		w.Write(data)

	}  else {
		w.WriteHeader(404)
		w.Write([]byte("404 - Sorry Page Not Found" + http.StatusText(404)))
	}
}


type Context struct {
	Title string
	Desc string
}

func main() {

	// http.Handle("/", new(urls))
	http.HandleFunc("/", mainFunc)
	http.ListenAndServe(":9090", nil)

}

func mainFunc(w http.ResponseWriter, req *http.Request){
	
	w.Header().Add("Content Type", "text/html")
	temp, err := template.New("home").Parse(doc)

	if err == nil{
		context := Context{"todo", "my todo list"}
		temp.Execute(w, context)
	}

}

const doc = `

	<!DOCTYPE html>
		<html>
		<head>
			<title>Home | Go</title>
			<link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.0.0-alpha/css/bootstrap.css">
		</head>
		<body>
			<h1>{{.Title}}</h1>
			<p> {{.Desc}} </p>
		</body>
		</html>

`



















