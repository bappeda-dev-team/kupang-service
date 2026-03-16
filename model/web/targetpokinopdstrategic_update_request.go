package web

type TargetPokinOpdStrategicUpdateRequest struct {
	Id                           int    `json:"id" validate:"required"`
	IndikatorPokinOpdStrategicId int    `json:"indikator_pokin_opd_strategic_id" validate:"required"`
	NamaTarget                   int    `json:"target" validate:"required"`
	Satuan                       string `json:"satuan" validate:"required"`
}
