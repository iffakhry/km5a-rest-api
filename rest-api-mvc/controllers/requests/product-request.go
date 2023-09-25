package requests

type ProductRequest struct {
	Name  string `json:"name" form:"name"`
	Stock int    `json:"stock" form:"stock"`
}
