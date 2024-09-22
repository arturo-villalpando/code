package payments

type PaymenthMethod interface {
	Pay(amount float64) string
}
