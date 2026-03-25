package web

type IndikatorPokinOpdOperationalNResponse struct {
	Id                     int    `json:"id,omitempty"`
	PokinOpdOperationalNId int    `json:"pokin_opd_operationalN_id"`
	NamaIndikator          string `json:"indikator"`
}
