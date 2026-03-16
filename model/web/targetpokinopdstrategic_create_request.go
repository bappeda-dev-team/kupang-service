package web

type TargetPokinOpdStrategicCreateRequest struct {
	IndikatorPokinOpdStrategicId int    `json:"indikator_pokin_opd_strategic_id" validate:"required"`
	NilaiTarget                  int    `json:"target" validate:"required"`
	Satuan                       string `json:"satuan" validate:"required"`
}
