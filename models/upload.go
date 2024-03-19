package models

type MovieUpload struct {
	Name     string `json:"name"`
	FileID   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
}
