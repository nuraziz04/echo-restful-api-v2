package web

type PegawaiCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Alamat  string `json:"alamat" validate:"required"`
	Telepon string `json:"telepon" validate:"required"`
}

type PegawaiCreateRequestLoop struct {
	Created string                 `json:"created" validate:"required"`
	Data    []PegawaiCreateRequest `json:"data" validate:"required"`
}

type PegawaiUpdateRequest struct {
	Id      int    `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Alamat  string `json:"alamat" validate:"required"`
	Telepon string `json:"telepon" validate:"required"`
}
