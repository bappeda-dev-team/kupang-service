package web

type LembagaCreateRequest struct {
	KodeLembaga          string `json:"kode_lembaga" validate:"required"`
	NamaLembaga          string `json:"nama_lembaga" validate:"required"`
	JabatanKepalaLembaga string `json:"jabatan_kepala_lembaga" validate:"required"`
	NamaKepalaLembaga    string `json:"nama_kepala_lembaga" validate:"required"`
	NipKepalaLembaga     string `json:"nip_kepala_lembaga" validate:"required"`
}
