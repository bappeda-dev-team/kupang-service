package web

type TargetPokinOpdOperationalUpdateRequest struct {
	Id                             int    `json:"id" validate:"required"`
	IndikatorPokinOpdOperationalId int    `json:"indikator_pokin_opd_operational_id" validate:"required"`
	NilaiTarget                    int    `json:"target" validate:"required"`
	Satuan                         string `json:"satuan" validate:"required"`
}
