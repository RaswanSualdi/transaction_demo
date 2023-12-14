package request

type RegisterMerchant struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Bank    Bank   `json:"bank" binding:"required"`
}
