package model

type Image struct {
	Created int         `json:"created"`
	Data    []ImageData `json:"data"`
}

type ImageData struct {
	Url string `json:"url"`
}
