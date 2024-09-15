package utils

import "log"

func SendEmailWarning(guid string) {
	log.Printf("user ip adress changed: %s", guid)
}
