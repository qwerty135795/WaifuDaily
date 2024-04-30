package dto

type RegisterDto struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	Waifu    string `json:"waifu"`
	Password string `json:"password"`
}
