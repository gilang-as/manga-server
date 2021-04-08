package dto

type ResponseGetManga struct {
	Total uint `json:"total"`
	Manga interface{} `json:"manga"`
}
