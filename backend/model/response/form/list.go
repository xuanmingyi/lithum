package form

type FormInfo struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

type ListResponse struct {
	Count uint64      `json:"count"`
	Items []*FormInfo `json:"items"`
}
