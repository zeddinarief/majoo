package user

// RegisterUserInput format
type RegisterUserInput struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NamaLengkap string `json:"nama_lengkap" binding:"required"`
}

// LoginInput format
type LoginInput struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// FormUpdateUserInput format
type FormUpdateUserInput struct {
	ID          int
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NamaLengkap string `json:"nama_lengkap" binding:"required"`
}

// DeleteInput format
type DeleteInput struct {
	ID int `json:"id" binding:"required"`
}
