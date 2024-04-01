package model

type DeviceResponse struct {
	ID          int    `json:"id,omitempty"`
	DeviceName  string `json:"device_name,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateDeviceRequest struct {
	DeviceName  string `json:"device_name,omitempty" validate:"required,max=50"`
	Description string `json:"description,omitempty" validate:"required"`
}

type UpdateDeviceResponse struct {
	ID          int    `json:"id,omitempty" validate:"required"`
	DeviceName  string `json:"device_name,omitempty" validate:"required,max=50"`
	Description string `json:"description,omitempty" validate:"required"`
}
