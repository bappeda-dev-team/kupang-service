package web

type MusrenbangCreateRequest struct {
	Usulan  string `json:"usulan" validate:"required"`
	Alamat  string `json:"alamat" validate:"required"`
	Uraian  string `json:"uraian" validate:"required"`
	Tahun   string `json:"tahun" validate:"required"`
	KodeOpd string `json:"kode_opd" validate:"required"`
	NamaOpd string `json:"nama_opd" validate:"required"`
	Status  string `json:"status" validate:"required"`
}
