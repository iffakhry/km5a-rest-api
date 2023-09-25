package requests

type UserRequest struct {
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Alamat   AlamatRequest `json:"alamat"`
}

type AlamatRequest struct {
	Jalan string `json:"jalan"`
	Kota  string `json:"kota"`
}
