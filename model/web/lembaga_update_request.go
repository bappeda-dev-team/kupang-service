package web

type LembagaUpdateRequest struct {
	Id          int    `json:"id" validate:"required"`
	KodeLembaga string `json:"kode_lembaga" validate:"required"`
	NamaLembaga string `json:"nama_lembaga" validate:"required"`
}
