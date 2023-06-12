package helpers

import (
	"fmt"
	"net/mail"
)

func ValidEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    fmt.Println(err)
	return err == nil
}