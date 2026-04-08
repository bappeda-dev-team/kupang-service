package web

type PeriodeCreateRequest struct {
	TahunAwal    string `json:"tahun_awal" validate:"required"`
	TahunAkhir   string `json:"tahun_akhir" validate:"required"`
	JenisPeriode string `json:"jenis_periode" validate:"required"`
}
