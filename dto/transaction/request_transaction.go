package transactiondto

type CreateTransaction struct {
	UserID int `json:"user_id" form:"user_id" validate:"required"`
}
