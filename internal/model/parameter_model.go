package model

type ParameterResponse struct {
	ID            int    `json:"id,omitempty"`
	ParameterName string `json:"parameter_name,omitempty"`
	TypeParam     string `json:"type_param,omitempty"`
	Good          string `json:"good,omitempty"`
	Bad           string `json:"bad,omitempty"`
	Notes         string `json:"notes,omitempty"`
}

type CreateParameterRequest struct {
	DeviceId      int    `json:"id,omitempty" validate:"required"`
	ParameterName string `json:"parameter_name,omitempty" validate:"required,max=100"`
	TypeParam     string `json:"type_param,omitempty" validate:"required,max=50"`
	Good          string `json:"good,omitempty" validate:"required,max=100"`
	Bad           string `json:"bad,omitempty" validate:"required,max=100"`
	Notes         string `json:"notes,omitempty" validate:"required"`
}

type UpdateParameterRequest struct {
	Id            int    `json:"id,omitempty" validate:"required"`
	ParameterName string `json:"parameter_name,omitempty" validate:"required,max=100"`
	TypeParam     string `json:"type_param,omitempty" validate:"required,max=50"`
	Good          string `json:"good,omitempty" validate:"required,max=100"`
	Bad           string `json:"bad,omitempty" validate:"required,max=100"`
	Notes         string `json:"notes,omitempty" validate:"required"`
}
