package user

type ListRequest struct {
	Username string `form:"username"`
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
}
