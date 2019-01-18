package main
import (
        "fmt"
        "net/http"
        "goji.io"
        "goji.io/pat"
        "html/template"
        )

func hello(w http.ResponseWriter, r *http.Request) {

	   // name := pat.Param(r, "name")
	   // fmt.Fprintf(w, "Hello %s", name)

	t, err := template.ParseFiles("home.html")
	if err != nil{
		fmt.Println(err)
	}
	t.Execute(w, "")
}

func main(){

 mux := goji.NewMux()
 mux.HandleFunc(pat.Get("/"), hello)
 http.ListenAndServe(":8080", mux)

}
