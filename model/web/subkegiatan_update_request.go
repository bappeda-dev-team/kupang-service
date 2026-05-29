package web

type SubkegiatanUpdateRequest struct {
	Id              int    `json:"id"`
	KodeSubkegiatan string `json:"kode_subkegiatan" validate:"required"`
	NamaSubkegiatan string `json:"nama_subkegiatan" validate:"required"`
	Tahun           string `json:"tahun" validate:"required"`
}
