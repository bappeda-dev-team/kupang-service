package web

type JabatanOpdResponse struct {
	Id           int    `json:"id,omitempty"`
	KodeJabatan  string `json:"kode_jabatan"`
	NamaJabatan  string `json:"nama_jabatan"`
	KodeOpd      string `json:"kode_opd"`
	Tahun        string `json:"tahun"`
}
