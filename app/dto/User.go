package dto


type Login struct {
	Username    string `json:"username" binding:"required"`
}
