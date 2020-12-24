package form

type CreateRequest struct {
	Title         string `form:"title"`
	Size          string `form:"size"`
	LabelPosition string `form"label_position"`
	LabelWidth    uint8  `form:"label_width"`
	Span          uint8  `form:"span"`
	FormBtns      bool   `form:"form_btns"`
	Disabled      bool   `form:"disabled"`
}
