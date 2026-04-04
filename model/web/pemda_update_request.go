package web

type PemdaUpdateRequest struct {
	Id        int    `json:"id" validate:"required"`
	KodePemda string `json:"kode_pemda" validate:"required"`
	NamaPemda string `json:"nama_pemda" validate:"required"`
}
