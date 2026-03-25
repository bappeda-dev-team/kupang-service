package web

type IndikatorPokinOpdOperationalNCreateRequest struct {
	PokinOpdOperationalNId int    `json:"pokin_opd_operationalN_id" validate:"required"`
	NamaIndikator          string `json:"indikator" validate:"required"`
}
