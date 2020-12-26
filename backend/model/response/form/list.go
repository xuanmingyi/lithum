package form

type FormInfo struct {
	ID            uint64 `json:"id"`
	Title         string `json:"title"`
	Size          string `json:"size"`
	LabelPosition string `json:"labelPosition"`
	LabelWidth    uint8  `json:"labelWidth"`
	Span          uint8  `json:"span"`
	Gutter        uint8  `json:"gutter"`
	FormBtns      bool   `json:"formBtns"`
	Disabled      bool   `json:"disabled"`
}

type ListResponse struct {
	Count uint64      `json:"count"`
	Items []*FormInfo `json:"items"`
}
