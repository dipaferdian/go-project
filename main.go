package main

import "fmt"

type Notifier interface {
    Send() string
}

type EmailNotifier struct {
    Email string
}

// Implementasikan method Send() menggunakan pointer receiver

func (e *EmailNotifier) Send() string  {
    return e.Email
}

func main() {
    email := EmailNotifier{
        Email: "user@example.com",
    }

    var notifier Notifier

    // Assign pointer EmailNotifier ke notifier

    notifier = &email

    fmt.Println(notifier.Send())
}