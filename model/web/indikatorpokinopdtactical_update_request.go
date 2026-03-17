package web

type IndikatorPokinOpdTacticalUpdateRequest struct {
	Id                 int    `json:"id" validate:"required"`
	PokinOpdTacticalId int    `json:"pokin_opd_tactical_id" validate:"required"`
	NamaIndikator      string `json:"indikator"`
}
