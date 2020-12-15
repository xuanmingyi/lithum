package role

type CreateRequest struct {
	Rolename string `json:"rolename"`
}

type CreateResponse struct {
	Rolename string `json:"rolename"`
}
