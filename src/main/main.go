package main


import (
	"net/http"
	"io/ioutil"
	"log"
	)

type urls struct{

}

func (this *urls) ServeHTTP(w http.ResponseWriter, r *http.Request){

	path := r.URL.Path[1:]
	log.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if string.HasSuffix(path, ".css"){
			contentType = "text/css"
		} else if string.HasSuffix(path, "js"){
			contentType = "text/js"
		} else if string.HasSuffix(path, "png"){
			contentType = "image/png"
		} else{
			contentType = "text/plain"
		}

		w.Header().Add("Content Type", contentType)
		w.Write(data)

	}  else {
		w.WriteHeader(404)
		w.Write([]byte("404" + http.StatusText(404)))
	}
}



func main() {

	http.Handle("/", new(urls))
	// http.HandleFunc("/", mainFunc)
	http.ListenAndServe(":9090", nil)

}

// func mainFunc(w http.ResponseWriter, req *http.Request){
// 	w.Write([]byte("Hello World"))
// }