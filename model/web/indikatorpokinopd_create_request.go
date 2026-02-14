package web

type IndikatorPokinOpdCreateRequest struct {
	TujuanPokinOpdId int    `json:"tujuan_pokin_opd_id"`
	NamaIndikator   string `json:"indikator"`
}
