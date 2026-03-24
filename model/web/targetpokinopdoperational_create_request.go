package web

type TargetPokinOpdOperationalCreateRequest struct {
	IndikatorPokinOpdOperationalId int    `json:"indikator_pokin_opd_operational_id" validate:"required"`
	NilaiTarget                    int    `json:"target" validate:"required"`
	Satuan                         string `json:"satuan" validate:"required"`
}
