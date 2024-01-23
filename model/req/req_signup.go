package req

type ReqSignUp struct {
	Email    string `json:"email, omitempty" validate:"required"`
	PassWord string `json:"password, omitempty" validate:"required"`
	FullName string `json:"fullName, omitempty" validate:"required"`
}
