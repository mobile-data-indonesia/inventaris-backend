package validators

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username    string  `json:"username" binding:"required,min=3,max=50"`
	Password    string  `json:"password" binding:"required,min=6,max=100"`
	Email       *string `json:"email" binding:"omitempty,email"`
	PhoneNumber *string `json:"phone_number" binding:"omitempty,max=20"`
	Role        string  `json:"role" binding:"required,oneof=staff admin"`
	Title 		string 	`json:"title" binding:"required,oneof=advisor member"`
	Department  string  `json:"department" binding:"required,oneof=it finance management"`
}

type UpdateUserRequest struct {
	Email       string `json:"email" binding:"omitempty,email"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,max=20"`
	Role        string  `json:"role" binding:"required,oneof=staff admin"`
	Title 		string `json:"title" binding:"required,oneof=advisor member"`
	Department  string  `json:"department" binding:"required,oneof=it finance management"`
}

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=100"`
}
