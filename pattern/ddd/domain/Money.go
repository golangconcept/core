package domain

type Money struct {
	Amount   float64
	Currency string
}

func (m Money) IsEqual(other Money) bool {
	return m.Amount == other.Amount && m.Currency == other.Currency
}

func NewMoney(amount float64, currency string) Money {
	return Money{Amount: amount, Currency: currency}
}
