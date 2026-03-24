package web

type IndikatorPokinOpdOperationalResponse struct {
	Id                     int    `json:"id,omitempty"`
	PokinOpdOperationalId  int    `json:"pokin_opd_operational_id"`
	NamaIndikator          string `json:"indikator"`
}
