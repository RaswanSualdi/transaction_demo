package request

type TransactionRequest struct {
	NumberAccount string `json:"sendToByNumberAccount"`
	Amount        int64  `json:"amount"`
}
