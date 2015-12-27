package main


import "net/http"

func main() {

	http.HandleFunc("/", mainFunc)
	http.ListenAndServe(":9090", nil)

}

func mainFunc(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Hello World"))
}