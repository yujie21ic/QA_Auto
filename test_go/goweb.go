package main
import (
"html/template"
"log"
"net/http"

)

const (
username = "root"
userpwd = ""
dbname = "gooo"
)

// User is user
type User struct {
ID string
Name string
}

// 单独提出来，渲染模板
func render(w http.ResponseWriter, tmplName string, context map[string]interface{}) {

// tmpl := template.New("index.html")
tmpl,err:= template.ParseFiles(tmplName)
// err:=nil
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
tmpl.Execute(w, nil)
return
}


//home界面
func indexHandler(w http.ResponseWriter, r *http.Request) {
_, err := getDB(username, userpwd, dbname)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
//locals可以作为模板变量去渲染，这里我没有使用
locals := make(map[string]interface{})
users := []User{}
locals["users"] = users
render(w, "./index.html",nil)
return
}


//用来写获取数据库的连接配置,并未真正实现
func getDB(username, userpwd, dbname string)( int,error) {	
return 123,nil
}

func main() {
//绑定
http.HandleFunc("/", indexHandler)
//绑定端口
err := http.ListenAndServe(":8880", nil)
if err != nil {
log.Fatal("ListenAndServe: ", err.Error())
}
}