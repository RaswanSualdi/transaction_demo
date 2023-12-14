package request

type Bank struct {
	BankName string `json:"bankName" binding:"required"`
	Saldo    int64  `json:"saldo" binding:"required"`
}
