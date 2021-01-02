package connection

import "lithum/model"

type ListResponse struct {
	Count uint64              `json:"count"`
	Items []*model.Connection `json:"items"`
}
