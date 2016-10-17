package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "net/smtp"
)

    func sendEmail(w http.ResponseWriter, r *http.Request) {
        err := r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
        if(err != nil){
          //throw error
        } else{

          varFrom := r.PostFormValue("from")
          varSubject := r.PostFormValue("subject")
          varMessage := r.PostFormValue("messgae")
          varDonEmail := "donstringham@weber.edu"
          varMeEmail := "markrichardson@mail.weber.edu"

          varEmailRobot := "emailservicemarkrichardson@gmail.com"
          varPassword := "password2@"

          err := smtp.SendMail("smtp.gmail.com:587",smtp.PlainAuth("", varFrom,))
        }



    }

    func main() {
        http.HandleFunc("/", sendEmail) // setting router rule

        err := http.ListenAndServe(":9090", nil) // setting listening port
        if err != nil {
            log.Fatal("ListenAndServe: ", err)
        }
}
