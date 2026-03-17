package web

type IndikatorPokinOpdTacticalResponse struct {
	Id                 int    `json:"id,omitempty"`
	PokinOpdTacticalId int    `json:"pokin_opd_tactical_id,omitempty"`
	NamaIndikator      string `json:"indikator"`
}
