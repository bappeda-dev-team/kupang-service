package web

type OpdCreateRequest struct {
	KodeOpd     string `json:"kode_opd" validate:"required"`
	NamaOpd     string `json:"nama_opd" validate:"required"`
	KodeLembaga string `json:"kode_lembaga" validate:"required"`
}
