package web

type OpdResponse struct {
	Id                           int    `json:"id,omitempty"`
	KodeOpd                      string `json:"kode_opd"`
	NamaOpd                      string `json:"nama_opd"`
	NamaKepalaPerangkatDaerah    string `json:"nama_kepala_perangkat_daerah"`
	NipKepalaPerangkatDaerah     string `json:"nip_kepala_perangkat_daerah"`
	PangkatKepalaPerangkatDaerah string `json:"pangkat_kepala_perangkat_daerah"`
	KodeLembaga                  string `json:"kode_lembaga"`
}
