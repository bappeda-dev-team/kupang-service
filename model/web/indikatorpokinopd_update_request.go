package web

type IndikatorPokinOpdUpdateRequest struct {
	Id            int    `json:"id" validate:"required"`
	NamaIndikator string `json:"indikator"`
}
