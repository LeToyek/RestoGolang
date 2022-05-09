package entities

type LoginReq struct {
	Email    string `db:"email" json:"email" validate:"email,required"`
	Password string `db:"password" json:"password" validate:"required,min=6"`
}
