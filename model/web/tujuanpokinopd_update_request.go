package web

type TujuanPokinOpdUpdateRequest struct {
	Id                int    `json:"id" validate:"required"`
	PokinOpdId        int    `json:"pokin_opd_id" validate:"required,gt=0"`
	KodeOpd           string `json:"kode_opd"`
	NamaTujuan        string `json:"tujuan"`
	BidangUrusan      string `json:"bidang_urusan"`
	TahunAwalPeriode  int    `json:"tahun_awal_periode"`
	TahunAkhirPeriode int    `json:"tahun_akhir_periode"`
}
