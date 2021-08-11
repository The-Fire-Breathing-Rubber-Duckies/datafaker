package entities

import (
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func GetEmails(number int, domain string, address string) (emails []string) {
	gofakeit.Seed(0)
	if number < 1 {
		panic("Number of emails should be greater than 0")
	}

	if domain == "" {
		domain = gofakeit.DomainName()
	}

	if number > 1 && address != "" {
		panic("Address is only useful for generating one email")
	}

	i := time.Now().UnixNano()
	for len(emails) < number {
		gofakeit.Seed(0)
		i = i + 1
		email := ""
		if address == "" {
			email = gofakeit.Name()
		}
		email = strings.ReplaceAll(email, " ", ".")
		email = strings.ToLower(email)

		email += "@" + domain

		emails = append(emails, email)
	}

	return emails
}
