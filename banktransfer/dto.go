package banktransfer

type BankTransferRequest struct {
	SourceBankSwiftCode          string
	SourceBankAccountNumber      string
	SourceBankAccountName        string
	DestinationBankSwiftCode     string
	DestinationBankAccountNumber string
	Amount                       float64
}
