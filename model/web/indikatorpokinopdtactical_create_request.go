package web

type IndikatorPokinOpdTacticalCreateRequest struct {
	PokinOpdTacticalId int    `json:"pokin_opd_tactical_id" validate:"required"`
	NamaIndikator      string `json:"indikator" validate:"required"`
}
