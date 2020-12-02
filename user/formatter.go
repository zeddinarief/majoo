package user

type UserFormatter struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	NamaLengkap string `json:"nama_lengkap"`
	Token       string `json:"token"`
	Foto        string `json:"foto"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Username:    user.Username,
		NamaLengkap: user.NamaLengkap,
		Token:       token,
		Foto:        user.Foto,
	}

	return formatter
}
