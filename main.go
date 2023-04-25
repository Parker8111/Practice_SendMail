package main

import (
	"log"
	"net/smtp"
	"strings"
)

var (
	host     = "smtp.gmail.com:587"
	username = "parkerlart@gmail.com"
	password = ""
	subject  = "測試用Golang發送郵件"
	target   = []string{"parker.fang@tmnewa.com.tw"}
)

func main() {
	auth := smtp.PlainAuth(host, username, password, "smtp.gmail.com")

	msg := []string{"Subject: This is a test mail!", "測試郵件由golang發送測試"}
	if err := smtp.SendMail(host, auth, username, target, []byte(strings.Join(msg, "\r\n"))); err != nil {
		log.Fatal(err)
	}

}
