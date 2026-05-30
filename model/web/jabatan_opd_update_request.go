package web

type JabatanOpdUpdateRequest struct {
	Id          int    `json:"id"`
	KodeJabatan string `json:"kode_jabatan" validate:"required"`
	NamaJabatan string `json:"nama_jabatan" validate:"required"`
	KodeOpd     string `json:"kode_opd" validate:"required"`
	Tahun       string `json:"tahun" validate:"required"`
}
