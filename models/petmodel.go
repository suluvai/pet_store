// This files contain the information about Pet Moddels.
package models

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

const (
	Available = "available"
	Pending   = "pending"
	Sold      = "sold"
)

type PhotoUrl string

type Pet struct {
	Id        int64      `json:"id"`
	Category  Category   `json:"category"`
	Name      string     `json:"name"`
	PhotoUrls []PhotoUrl `json:"photoUrls"`
	Tags      []Tag      `json:"tags"`
	Status    string     `json:"status"`
}
