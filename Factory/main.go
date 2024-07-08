package main

import "fmt"

// SMS Email

type INotificacionFactory interface {
	SendNotificacion()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotificacion struct {
}

func (SMSNotificacion) SendNotificacion() {
	fmt.Println("Se envia la notificacion via SMS")
}

func (SMSNotificacion) GetSender() ISender {
	return SMSNotificacionSender{}
}

type SMSNotificacionSender struct {
}

func (SMSNotificacionSender) GetSenderMethod() string {

	return "SMS"
}

func (SMSNotificacionSender) GetSenderChannel() string {
	return "Twiilio"
}

type EmailNotification struct {
}

func (EmailNotification) SendNotificacion() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func GetNotificationFactory(notificationType string) (INotificacionFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotificacion{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No Notification Type")
}

func SendNotificacion(f INotificacionFactory) {
	f.SendNotificacion()
}

func getMethod(f INotificacionFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := GetNotificationFactory("SMS")
	emailFatory, _ := GetNotificationFactory("Email")

	SendNotificacion(smsFactory)
	SendNotificacion(emailFatory)

	getMethod(smsFactory)
	getMethod(emailFatory)
}
