package web

type PokinOpdUpdateRequest struct {
	Id      int    `json:"id" validate:"required"`
	KodeOpd string `json:"kode_opd"`
	NamaOpd string `json:"nama_opd"`
	Tahun   int    `json:"tahun"`
}
