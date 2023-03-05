package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"gopkg.in/gomail.v2"
)

func sendMailSimple(subject string, body string, to []string){
	auth := smtp.PlainAuth(
		"",
		"myemail@gmail.com",
		"apppassword",
		"smtp.gmail.com",

	)
	msg := "Subject :" + subject + "\n" + body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"okethis99@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}
func sendMailSimpleHTML(subject string, templatePath string, to []string){

	// Get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{Name string}{Name: "Aly"})
	if err != nil {
		fmt.Println(err)
		return
		
	}

	auth := smtp.PlainAuth(
		"",
		"myemail@gmail.com",
		"appassword",
		"smtp.gmail.com",

	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject :" + subject + "\n" + headers + "\n\n"+ body.String()
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"myemail@gmail.com",
		to,
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}
func sendGomail(templatePath string){
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{Name string}{Name: "Aly"})

	if err != nil {
		fmt.Println(err)
		return

	}

//send with gomail
m := gomail.NewMessage()
m.SetHeader("From", "myemail@gmail.com")
m.SetHeader("To", "sending@gmail.com", "sending@gmail.com","sending@gmail.com")
m.SetHeader("Subject", "MAIL TITLE!")
m.SetBody("text/html", body.String())
m.Attach("./me.jpg")

d := gomail.NewDialer("smtp.gmail.com", 587, "myemail@gmail.com", "apppassword")

// Send the email to Bob, Cora and Dan.
if err := d.DialAndSend(m); err != nil {
	panic(err)
}
}
func main() {
	fmt.Println("Sending Mail...")
	/*sendMailSimple(
		"Test 2", 
		"BOdy MEsage", 
		[]string{"okethis@gmail.com","aliyilmaz9977@gmail.com"},
	)*/
	/*
	sendMailSimpleHTML(
		"Baslik Deneme 1", 
		"./test.html", 
		[]string{"okethis@gmail.com","aliyilmaz9977@gmail.com"},
	)*/
	sendGomail("./test.html")
	
}