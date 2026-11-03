package main

import "fmt"


type Payment interface {
    Pay()
}

type CreditCard struct {
    CardNumber string
}

type EWallet struct {
    Provider string
}

func (c CreditCard) Pay()  {
    fmt.Println("Payment using Credit Card")
}

func (e EWallet) Pay()  {
    fmt.Println("Payment using E-Wallet")
}

func processPayment(payment Payment) {
    payment.Pay()
}

func main() {
    
    creditCard := CreditCard{
        CardNumber: "09121213218",
    }

    ewallet := EWallet{
        Provider: "ovo",
    }

    processPayment(creditCard)
    processPayment(ewallet)
    
}