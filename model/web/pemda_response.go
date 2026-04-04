package web

type PemdaResponse struct {
	Id        int    `json:"id,omitempty"`
	KodePemda string `json:"kode_pemda"`
	NamaPemda string `json:"nama_pemda"`
}
