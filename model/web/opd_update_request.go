package web

type OpdUpdateRequest struct {
	Id                           int    `json:"id" validate:"required"`
	KodeOpd                      string `json:"kode_opd" validate:"required"`
	NamaOpd                      string `json:"nama_opd" validate:"required"`
	KodeLembaga                  string `json:"kode_lembaga" validate:"required"`
	NamaKepalaPerangkatDaerah    string `json:"nama_kepala_perangkat_daerah" validate:"required"`
	NipKepalaPerangkatDaerah     string `json:"nip_kepala_perangkat_daerah" validate:"required"`
	PangkatKepalaPerangkatDaerah string `json:"pangkat_kepala_perangkat_daerah" validate:"required"`
}
