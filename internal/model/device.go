package model

type CreateDeviceReq struct {
	Name string  `json:"name" validate:"required,min=3,max=100"`
	Lat  float32 `json:"lat" validate:"required"`
	Long float32 `json:"long" validate:"required"`
}

type CreateDeviceRes struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Lat       float32 `json:"lat"`
	Long      float32 `json:"long"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}
