package web

type KegiatanCreateRequest struct {
	KodeKegiatan string `json:"kode_kegiatan" validate:"required"`
	NamaKegiatan string `json:"nama_kegiatan" validate:"required"`
	Tahun        string `json:"tahun" validate:"required"`
}
