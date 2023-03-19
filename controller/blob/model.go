package blobController

import "time"

type BlobResponse struct {
	Title     string    `json:"title"`
	Alt       string    `json:"alt"`
	Id        string    `json:"id"`
	CreateTime time.Time `json:"createTime"`
}