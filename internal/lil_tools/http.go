package gokupkg

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
)


/*

HTTP DEMO, but generally use frames like beego, gin ..

*/



// define port
func Http_run(){
	http.HandleFunc("/", http_server)
    log.Fatal(http.ListenAndServe(":9090", nil))
    if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

// http server
func Http_server(w http.ResponseWriter, r *http.Request) {
    // html post on web
    data, err := template.ParseFiles("./web/test.html")
    if err != nil {
        log.Fatal("template.ParseFiles : ", err)
        return
    }
    data.Execute(w, nil)

//    // file data post on web
//    data,err:= ioutil.ReadFile("./web/result.json")
//        if err != nil {
//                log.Fatal(err)
//        }
//    fmt.Fprintln(w, string(data)) //output json

}
