//"If a type has these methods, it can be used wherever this interface is expected."

//This is called implicit implementation, and it's one of Go's most powerful features.
package payment

import (
	"errors"
	"fmt"
	"testing"
)

type Payment interface {
	Pay(amount float64) error
}
// Business Logic
type PaymentService struct {
	payment Payment
}

func NewPaymentService(p Payment) *PaymentService {
	return &PaymentService{
		payment: p,
	}
}

func (p *PaymentService) Checkout(amount float64) error {
	return p.payment.Pay(amount)
}
// Real Implementation


type Razorpay struct{}

func (r Razorpay) Pay(amount float64) error {
	fmt.Printf("Payment of ₹%.2f successful using Razorpay\n", amount)
	return nil
}
// Mock for Testing

type MockPayment struct {
	Called bool
	Amount float64
}

func (m *MockPayment) Pay(amount float64) error {
	m.Called = true
	m.Amount = amount
	return nil
}
// Mock Failure
type FailedPayment struct{}

func (f FailedPayment) Pay(amount float64) error {
	return errors.New("payment failed")
}
// Tests

func TestCheckoutSuccess(t *testing.T) {
	mock := &MockPayment{}

	service := NewPaymentService(mock)

	err := service.Checkout(999)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if !mock.Called {
		t.Error("Pay() was not called")
	}

	if mock.Amount != 999 {
		t.Errorf("expected amount 999, got %.2f", mock.Amount)
	}
}

func TestCheckoutFailure(t *testing.T) {
	service := NewPaymentService(FailedPayment{})

	err := service.Checkout(500)

	if err == nil {
		t.Fatal("expected an error but got nil")
	}

	if err.Error() != "payment failed" {
		t.Errorf("unexpected error: %v", err)
	}
}