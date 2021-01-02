package connection

type CreateRequest struct {
	Name    string `json:"name"`
	Display string `json:"display"`
	Config  string `json:"config"`
}
