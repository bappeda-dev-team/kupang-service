package web

type IndikatorPokinOpdStrategicUpdateRequest struct {
	Id                  int    `json:"id" validate:"required"`
	PokinOpdStrategicId int    `json:"pokin_opd_strategic_id" validate:"required"`
	NamaIndikator       string `json:"indikator"`
}
