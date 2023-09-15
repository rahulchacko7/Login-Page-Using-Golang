package main

import (
	"html/template"
	"net/http"
	
)


var tmpl *template.Template

type UserSt struct{
	Username string
	Password string
}

var users=make(map[string]UserSt)
var session=make(map[string]string)

type errors struct {
	UsernameError string
	PasswordError string
}


func init() {
	 
	tmpl = template.Must(template.ParseGlob("templates/*"))

	users["rahulchacko7@gmail.com"] = UserSt{"Rahul","4732"}
	users["dileep@gmail.com"] = UserSt{"Dileep","4732"}
	users["achu@gmail.com"]= UserSt{"Achu","4732"}
}


func main() {

	http.HandleFunc("/", loginPage)
	http.HandleFunc("/home", home)
	http.HandleFunc("/logout", logout)
	//http.HandleFunc("/signup",signup)
	http.ListenAndServe(":8000", nil)
}
