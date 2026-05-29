package web

type KegiatanResponse struct {
	Id           int    `json:"id,omitempty"`
	KodeKegiatan string `json:"kode_kegiatan"`
	NamaKegiatan string `json:"nama_kegiatan"`
	Tahun        string `json:"tahun"`
}
