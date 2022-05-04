package model

type LoginDto struct {
	Email 	 string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,password"`
}

type SignupDto struct {
	LoginDto
	Name string `json:"name" validate:"required,name"`
}

type UserResponse struct {
	ID 		 uint   `json:"id"`
	Name 	 string `json:"name"`
	Email	 string `json:"email"`
	Password string `json:"password"`
}
type AccessResponse struct {
	Token string `json:"token"`
}
type AuthResponse struct {
	User *UserResponse   `json:"user"`
	Auth *AccessResponse `json:"auth"`
}
