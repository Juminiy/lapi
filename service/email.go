package service

import (
	envConfig "lapi/config"
	"log"
	"net/smtp"
)

var (
	user = envConfig.Config("MAIL_USER")
	pass = envConfig.Config("MAIL_PASS")
	host = envConfig.Config("MAIL_HOST")
	port = envConfig.Config("MAIL_PORT")
	nick = envConfig.Config("MAIL_NICK")
)
func SendEmail(mTo,mSubject,mBody,mType string) error {
	log.Println("mData",mTo,mSubject,mBody,mType)
	mAuth := smtp.PlainAuth("",user,pass,host)
	var contentType string
	if mType == "html" {
		contentType = "Content-Type: text/" + mType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	mMsg := []byte("To: "+mTo+ "\r\nFrom: "+nick+"<"+user+">" +"\r\nSubject: "+mSubject+"\r\n"+contentType+"\r\n\r\n"+mBody)
	err := smtp.SendMail(host+":"+port,mAuth,user,[]string{mTo},mMsg)
	return err
}