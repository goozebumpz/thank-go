package lesson12

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Email struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
}

func FilterEmails(dst io.Writer, src io.Reader, predicate func(e Email) bool) (int, error) {
	counter := 0
	d := json.NewDecoder(src)
	enc := json.NewEncoder(dst)

	for {
		var email Email
		err := d.Decode(&email)

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}

		isTrue := predicate(email)

		if !isTrue {
			continue
		}

		err = enc.Encode(email)

		if err != nil {
			return 0, err
		}
		counter++
	}

	return counter, nil
}

func Test() {
	src := strings.NewReader(`
		{ "from": "alice@go.dev",      "to": "zet@php.net",              "subject": "How are you?" }
		{ "from": "bob@temp-mail.org", "to": "yolanda@java.com",         "subject": "Re: Indonesia" }
		{ "from": "cindy@go.dev",      "to": "xavier@rust-lang.org",     "subject": "Go vs Rust" }
		{ "from": "diane@dart.dev",    "to": "wanda@typescriptlang.org", "subject": "Our crypto startup" }
	`)
	dst := os.Stdout

	predicate := func(email Email) bool {
		return !strings.Contains(email.Subject, "crypto")
	}

	n, err := FilterEmails(dst, src, predicate)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, "good emails")

	// {"from":"alice@go.dev","to":"zet@php.net","subject":"How are you?"}
	// {"from":"bob@temp-mail.org","to":"yolanda@java.com","subject":"Re: Indonesia"}
	// {"from":"cindy@go.dev","to":"xavier@rust-lang.org","subject":"Go vs Rust"}
	// 3 good emails
}
