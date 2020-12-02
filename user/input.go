package user

type RegisterUserInput struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NamaLengkap string `json:"nama_lengkap" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type FormUpdateUserInput struct {
	ID          int
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NamaLengkap string `json:"nama_lengkap" binding:"required"`
	Error       error
}
