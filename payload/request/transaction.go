package request

type TransactionRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	NumberAccount string `json:"sendToByNumberAccount"`
	Amount        int64  `json:"amount"`
}
