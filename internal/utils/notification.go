package utils

import "log"

func SendEmailWarning(email string) {
	log.Printf("user ip adress changed: send message on %s", email)
}
