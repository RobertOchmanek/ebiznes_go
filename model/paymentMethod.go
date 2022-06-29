package model

type PaymentMethod int64

const (
	Blik PaymentMethod         = 0
	CreditCard PaymentMethod   = 1
	BankTransfer PaymentMethod = 2
)

func (pm PaymentMethod) String() string {
	switch pm {
	case Blik:
		return "BLIK"
	case CreditCard:
		return "Credit Card"
	case BankTransfer:
		return "Bank Transfer"
	}
	return "Unknown"
}