package web

type PemdaCreateRequest struct {
	KodePemda string `json:"kode_pemda" validate:"required"`
	NamaPemda string `json:"nama_pemda" validate:"required"`
}
