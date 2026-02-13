package web

type PokinOpdCreateRequest struct {
	KodeOpd string `json:"kode_opd"`
	NamaOpd string `json:"nama_opd"`
	Tahun   int    `json:"tahun"`
}
