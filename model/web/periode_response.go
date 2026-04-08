package web

type PeriodeResponse struct {
	Id           int    `json:"id,omitempty"`
	TahunAwal    string `json:"tahun_awal"`
	TahunAkhir   string `json:"tahun_akhir"`
	JenisPeriode string `json:"jenis_periode"`
}
