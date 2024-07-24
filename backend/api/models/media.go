package models

type Media struct {
	ID           int64
	UploadedBy   int64
	ClientEmail  string
	PreviewPath  string
	OriginalPath string
	Price        float64
	Locked       bool
}
