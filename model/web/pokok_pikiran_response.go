package web

type PokokPikiranResponse struct {
	Id      int     `json:"id,omitempty"`
	Usulan  string  `json:"usulan"`
	Alamat  string  `json:"alamat"`
	Uraian  string  `json:"uraian"`
	Tahun   *string `json:"tahun,omitempty"`
	KodeOpd string  `json:"kode_opd"`
	NamaOpd string  `json:"nama_opd"`
	Status  *string `json:"status,omitempty"`
}
