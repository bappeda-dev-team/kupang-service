package web

type IndikatorPokinOpdStrategicCreateRequest struct {
	PokinOpdStrategicId int    `json:"pokin_opd_strategic_id" validate:"required"`
	NamaIndikator       string `json:"indikator"`
}
