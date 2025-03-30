package common

type ReferenceDto struct {
	Id   string `json:"Id" validate:"required"`
	Type string `json:"Type" validate:"required"`
}
