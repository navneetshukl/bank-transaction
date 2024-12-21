package user

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Account string `json:"account,omitempty"`
	Money   int64  `json:"money"`
	Bank    string `json:"bank"`
}
