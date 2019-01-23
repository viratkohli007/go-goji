package main
import (
       "fmt"
       "net/http"
       "github.com/zenazn/goji"
       "github.com/zenazn/goji/web"
       "html/template"
       "go-goji/funcs"
       //"regexp"
       )

func home(c web.C, w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("home.html")
    t.Execute(w, "")
}

func form(c web.C, w http.ResponseWriter, r *http.Request) {

    t, _ := template.ParseFiles("form.html")
    t.Execute(w, "")
}

// func editform(c web.C, w http.ResponseWriter, r *http.Request){

// 	  editquery := `update users set age= $2, first_name = $3, last_name = $4, email=$5 where id=$1`
// 	  t, _ := template.ParseFiles("Editform.html")
// 	  t.Execute(w, "")
// }

func main() {

	    db2 := funcs.Dbconn()
      fmt.Println(db2)

     goji.Get("/home", home)
     ///fmt.Println(a)
     goji.Get("/form", form)
     goji.Get("/display", funcs.Display)
     goji.Get("/list", funcs.List)
     goji.Get("/edit", funcs.Update)
     goji.Get("/display2", funcs.Display2)
     //goji.Get("/display2", funcs.Delete)
     goji.Get("/del", funcs.Delete)
     goji.Serve()
}


