package web

type PokokPikiranUpdateRequest struct {
	Id      int     `json:"id" validate:"required"`
	Usulan  string  `json:"usulan" validate:"required"`
	Alamat  string  `json:"alamat" validate:"required"`
	Uraian  string  `json:"uraian" validate:"required"`
	Tahun   *string `json:"tahun,omitempty"`
	KodeOpd string  `json:"kode_opd" validate:"required"`
	NamaOpd string  `json:"nama_opd" validate:"required"`
	Status  *string `json:"status,omitempty"`
}
