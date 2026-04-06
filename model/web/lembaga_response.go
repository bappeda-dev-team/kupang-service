package web

type LembagaResponse struct {
	Id          int    `json:"id,omitempty"`
	KodeLembaga string `json:"kode_lembaga"`
	NamaLembaga string `json:"nama_lembaga"`
}
