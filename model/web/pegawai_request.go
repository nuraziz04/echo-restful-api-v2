package web

type PegawaiCreateRequest struct {
	Name    string `json:"name"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

type PegawaiUpdateRequest struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}
