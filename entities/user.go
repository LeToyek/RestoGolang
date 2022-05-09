package entities

type User struct {
	First_name string `db:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	Last_name  string `db:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Email      string `db:"email" json:"email" validate:"email,required"`
	Password   string `db:"password" json:"password" validate:"required,min=6"`
	Phone      string `db:"phone" json:"phone" validate:"required"`
	Created_at string `db:"created_at" json:"created_at"`
	Updated_at string `db:"updated_at" json:"updated_at"`
	User_id    string `db:"user_id" json:"user_id" validate:"required"`
}
