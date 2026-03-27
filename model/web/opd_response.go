package web

type OpdResponse struct {
	Id      int    `json:"id,omitempty"`
	KodeOpd string `json:"kode_opd"`
	NamaOpd string `json:"nama_opd"`
}
