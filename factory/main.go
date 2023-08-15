package main

import "fmt"

//SMS Email

type iNotificationFactory interface {
	sendNotification()
	getSender() iSender
}

type iSender interface {
	getSenderMethod() string
	getSenderChannel() string
}

type smsNotification struct {
}

func (sms smsNotification) sendNotification() {
	fmt.Printf("Sending SMS notification\n")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) getSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) getSenderChannel() string {
	return "Twilio" //sirve para mandar mensajes de confirmacion de sms
}

func (sms smsNotification) getSender() iSender {
	return SMSNotificationSender{}
}

type EmailNotification struct {
}

func (email EmailNotification) sendNotification() {
	fmt.Printf("Sending Email notification\n")
}

func (EmailNotification) getSender() iSender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) getSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) getSenderChannel() string {
	return "SES"

}

func getNotificationFactory(notificationType string) (iNotificationFactory, error) {
	if notificationType == "SMS" {
		return &smsNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No notification type")
}

func sendNotification(f iNotificationFactory) {
	f.sendNotification()
}

func getMethod(f iNotificationFactory) {
	fmt.Println(f.getSender().getSenderMethod())
}

func main() {

	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)

}
