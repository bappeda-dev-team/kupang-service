package web

type KegiatanUpdateRequest struct {
	Id           int    `json:"id"`
	KodeKegiatan string `json:"kode_kegiatan" validate:"required"`
	NamaKegiatan string `json:"nama_kegiatan" validate:"required"`
	Tahun        string `json:"tahun" validate:"required"`
}
