// Package event provides types and functionality for handling various events in the system.
package event

import "github.com/google/uuid"

// PaymentMethodCreated represents an event that is triggered when a payment method is created.
// This event is an example of demonstrating how to handle events in a system.
type PaymentMethodCreated struct {
	EventID           uuid.UUID         `json:"eventID"`
	UserID            uuid.UUID         `json:"userID"`
	PaymentMethodID   uuid.UUID         `json:"paymentMethodID"`
	PaymentMethodType PaymentMethodType `json:"paymentMethodType"`
}

// PaymentMethodType represents the type of payment method used by a user.
// It is used to categorize different payment methods in the system.
type PaymentMethodType string

const (
	// Card represents a credit or debit card payment method.
	Card PaymentMethodType = "card"

	// BankAccount represents a direct bank account payment method.
	BankAccount PaymentMethodType = "bank_account"
)
