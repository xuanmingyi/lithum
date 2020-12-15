package user

import (
	. "lithum/model/response/user"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type SwaggerListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	UserList   []UserInfo `json:"userList"`
}
