package dto

type PageMetaData struct {
	Page      int64  `json:"page"`
	Size      int64  `json:"size"`
	TotalItem int64  `json:"total_item"`
	TotalPage int64  `json:"total_page"`
	Links     *Links `json:"links"`
}

type Links struct {
	Self  string `json:"self"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	First string `json:"first"`
	Last  string `json:"last"`
}
