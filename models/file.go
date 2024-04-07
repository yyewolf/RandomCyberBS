package models

type File struct {
	Name        string `json:"name" bson:"name"`
	Slug        string `json:"slug" bson:"slug"`
	Location    string `json:"-" bson:"location"`
	Checksum    string `json:"checksum" bson:"checksum"`
	Link        string `json:"link" bson:"link"`
	FileSize    string `json:"file_size" bson:"file_size"`
	ContentType string `json:"content_type" bson:"content_type"`
}
