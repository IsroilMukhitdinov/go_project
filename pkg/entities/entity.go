package entities

type Currency struct {
	CurrencyName  string
	CurrencyValue string
}

type PageData struct {
	BankName string
	List     []Currency
}
