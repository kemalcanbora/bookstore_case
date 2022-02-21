package models

type Book struct {
	ID       string  `json:"id"     validate:"required" bson:"_id" `
	Quantity int32   `json:"quantity"  validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Title    string  `json:"title"  validate:"required"`
	Author   string  `json:"author" validate:"required"`
	Type     string  `json:"type"   validate:"required"`
}
