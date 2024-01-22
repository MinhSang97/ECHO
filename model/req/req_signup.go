package req

type ReqSignUp struct {
	FullName string `validate:"required"`
	Email    string `validate:"required"`
	PassWord string `validate:"required"`
}
