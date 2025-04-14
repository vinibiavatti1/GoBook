package main

import "fmt"

type Sender interface {
	Send(message string)
}

type (
	EmailNotification struct{}
	SMSNotification   struct{}
	PushNotification  struct{}
)

func (e *EmailNotification) Send(message string) {
	fmt.Println("Sending Email:", message)
}
func (s *SMSNotification) Send(message string) {
	fmt.Println("Sending SMS:", message)
}
func (p *PushNotification) Send(message string) {
	fmt.Println("Sending Push Notification:", message)
}

type Notifier struct{}

func (n *Notifier) Notify(sender Sender, message string) {
	sender.Send(message)
}
func (n *Notifier) NotifyAll(senders []Sender, message string) {
	for _, s := range senders {
		s.Send(message)
	}
}

func runNotification() {
	en := &EmailNotification{}
	sn := &SMSNotification{}
	pn := &PushNotification{}

	n := &Notifier{}

	n.Notify(en, "Hello")
	n.Notify(sn, "Hello")
	n.Notify(pn, "Hello")

	n.NotifyAll([]Sender{en, sn, pn}, "Hello")
}
