package funcs
import (
        "fmt"
        "net/http"
        "github.com/zenazn/goji/web"
        "database/sql"
      _ "github.com/lib/pq"
        "html/template"
      _ "github.com/go-sql-driver/mysql"
        //"sort"
      "regexp"
       )

type Formst struct{
   Id string
   FirstName string
   LastName string
   Age string
   Email string
}

func Dbconn() *sql.DB{

    const (
       port = 5432
       dbname = "postgres"
       username = "postgres"
       password = "test123"
       host = "localhost"
          )

     psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s"+
     	" dbname=%s sslmode=disable ",host, port, username, password, dbname)

     db, err := sql.Open("postgres", psqlinfo)
     if err != nil {
     	panic(err)
     }
     //fmt.Println("connected")
     return db
}

  func Display(c web.C, w http.ResponseWriter, r *http.Request) {
   flag := false
      if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
        fmt.Println("no")
    }else{
        fmt.Println("yes")
        flag = true
    }
    if flag == true{
   db2 := Dbconn()

   dis := new(Formst)
   dis.FirstName = r.FormValue("first_name")
   dis.LastName = r.FormValue("last_name")
   dis.Age = r.FormValue("age")
   dis.Email = r.FormValue("email")

  insert := `insert into users(age, first_name, last_name, email) values($1, $2, $3, $4)`
  _, err := db2.Exec(insert, dis.Age, dis.FirstName, dis.LastName, dis.Email)
  if err != nil{
    panic(err)
  }

   t, _:= template.ParseFiles("display.html")
   t.Execute(w, dis)
 }else{
  fmt.Fprintf(w, `<html>
            <head>
            </head>
            <body>
            <script type="text/javascript">
            alert('hahaha')
            </script>
            </body>
            </html>`)
  t, _:= template.ParseFiles("display.html")
   t.Execute(w, "")

 }
}

func List(c web.C, w http.ResponseWriter, r *http.Request){

      db2 := Dbconn()
      var d Formst
      var arr []Formst
      sql := `select * from users`
      rows, err := db2.Query(sql)
      if err!= nil {
        panic(err)
      }
      for rows.Next(){
        err = rows.Scan(&d.Id, &d.Age, &d.FirstName, &d.LastName, &d.Email)
        if err != nil {
          panic(err)
        }
      arr = append(arr, d)
    }
      //arr = sort.Sort(Formst(d))
      t, _ := template.ParseFiles("list.html")
      t.Execute(w, arr)
}

func Update(c web.C, w http.ResponseWriter, r *http.Request) {

     db2 := Dbconn()
      //var d Formst

       //d.Id = r.FormValue("btn1")
    q := new(Formst)
    q.Id = r.FormValue("btn1")
    //  q.Age = r.FormValue("age")
    // q.FirstName = r.FormValue("first_name")
    // q.LastName = r.FormValue("last_name")
    // q.Email = r.FormValue("email")
  fmt.Println(">>>>>>>>>>>>",q.Id)

   //return d.Id

   sql1 := `select * from users where id=$1`
   res, err := db2.Query(sql1, q.Id)
   if err != nil{
    fmt.Println(err)
   }
   fmt.Println(res)
   for res.Next(){
    err = res.Scan(&q.Id, &q.Age, &q.FirstName, &q.LastName, &q.Email)
    if err != nil {
    panic(err)
     }
   }

  t, _ := template.ParseFiles("Editform.html")
   t.Execute(w, q)
  //  editquery := `update users set age= $2, first_name = $3, last_name = $4, email=$5 where id=$1`
  // _, err = db2.Query(editquery,q.Id, q.Age, q.FirstName, q.LastName, q.Email)
  // if err != nil{
  //   panic (err)
  // }
}


func Display2 (c web.C, w http.ResponseWriter, r *http.Request) {
     db2 := Dbconn()

    dis := new(Formst)
   dis.Id = r.FormValue("id2")
    dis.Age = r.FormValue("age")
   dis.FirstName = r.FormValue("first_name")
   dis.LastName = r.FormValue("last_name")
    dis.Email = r.FormValue("email")

   Id := r.FormValue("id2")
    Age := r.FormValue("age")
   FirstName := r.FormValue("first_name")
   LastName := r.FormValue("last_name")
    Email := r.FormValue("email")


  // fmt.Println(">>>>>>>>>>>>")
    //Id2 := r.FormValue("id2")
 fmt.Println("######", Id, Age, FirstName, LastName, Email)

  //  err := db2.QueryRow("UPDATE users SET age=$2 WHERE id=$1", Id, Age)
  // if err != nil{
  //   fmt.Println(err)
  // }
  //   tx, err := db2.Begin()
  //   if err!= nil {
  //     fmt.Println(err)
  //   }
  // _, _ = tx.Stmt(updateMoney).Exec(7, "22")

insForm, err := db2.Prepare("UPDATE users SET age=($1), first_name=($2), last_name=($3), email=($4) WHERE id=($5)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(Age, FirstName, LastName, Email, dis.Id)

 //  editquery := `update users set age= ?, first_name = ?, last_name = ?, email=? where id=?;`
 //   _, _ = db2.Prepare(editquery)
 //  fmt.Println(editquery)
 // //  if err != nil{
 // //    panic (err)
 // // }
 t, _ := template.ParseFiles("display2.html")
   t.Execute(w, dis)
}

func Delete(w http.ResponseWriter, r *http.Request) {

   db2 := Dbconn()
   dis := new(Formst)
   dis.Id = r.FormValue("btn2")
   Id := r.FormValue("btn2")
   fmt.Println(Id)
   _, err := db2.Query("delete from users where id=($1) ", Id)
   if err != nil{
    fmt.Println(err)
   }
   //del.Query(dis.Id)
      var d Formst
      var arr []Formst
      sql := `select * from users`
      rows, err := db2.Query(sql)
      if err!= nil {
        panic(err)
      }
      for rows.Next(){
        err = rows.Scan(&d.Id, &d.Age, &d.FirstName, &d.LastName, &d.Email)
        if err != nil {
          panic(err)
        }
        arr = append(arr, d)
      }
      //arr = sort.Sort(Formst(d))
      t, _ := template.ParseFiles("list.html")
      t.Execute(w, arr)

}
