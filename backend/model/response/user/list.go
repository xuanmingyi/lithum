package user

import (
	"sync"
)

type UserInfo struct {
	Id           uint64 `json:"id"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type ListResponse struct {
	Count uint64      `json:"count"`
	Items []*UserInfo `json:"items"`
}
