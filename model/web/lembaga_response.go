package web

type LembagaResponse struct {
	Id                   int    `json:"id,omitempty"`
	KodeLembaga          string `json:"kode_lembaga"`
	NamaLembaga          string `json:"nama_lembaga"`
	JabatanKepalaLembaga string `json:"jabatan_kepala_lembaga"`
	NamaKepalaLembaga    string `json:"nama_kepala_lembaga"`
	NipKepalaLembaga     string `json:"nip_kepala_lembaga"`
}
