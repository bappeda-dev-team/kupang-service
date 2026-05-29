package web

type SubkegiatanResponse struct {
	Id              int    `json:"id,omitempty"`
	KodeSubkegiatan string `json:"kode_subkegiatan"`
	NamaSubkegiatan string `json:"nama_subkegiatan"`
	Tahun           string `json:"tahun"`
}
