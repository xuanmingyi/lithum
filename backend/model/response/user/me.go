package user

type MeResponse struct {
	Username     string `json:"username"`
	Nickname     string `json"nickname"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
}
