package web

type TargetPokinOpdUpdateRequest struct {
	Id         int    `json:"id" validate:"required"`
	NamaTarget int    `json:"target"`
	Satuan     string `json:"satuan"`
}
