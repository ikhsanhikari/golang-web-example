package main

import "fmt"
import "net/http"
import "html/template"
import "path"

type M map[string]interface{}


func handlerIndex(w http.ResponseWriter, r *http.Request){
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request){
	
}


func main(){

	http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("assets"))))
			
			
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		 var data = M{"name": "Batman"}
		 var tmpl = template.Must(template.ParseFiles(
			 "views/index.html",
			 "views/_header.html",
			 "views/_message.html",
		 ))

		 var err = tmpl.ExecuteTemplate(w, "index", data)
		 if err != nil {
			 http.Error(w, err.Error(), http.StatusInternalServerError)
		 }
	})
	
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		 var tmpl = template.Must(template.ParseFiles(
			 "views/about.html",
			 "views/_header.html",
			 "views/_message.html",
		 ))

		 var err = tmpl.ExecuteTemplate(w, "about", data)
		 if err != nil {
			 http.Error(w, err.Error(), http.StatusInternalServerError)
		 }
	 })
	
	
	
	
	
	//var address = "localhost:9090"
	//fmt.Printf("running at %s",address)
	//err:=http.ListenAndServe(address,nil)
	//if err!=nil{
		//fmt.Print(err.Error())
	//}
	
	var address = ":9000"
	fmt.Printf("server started at %s\n", address)

	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}

}