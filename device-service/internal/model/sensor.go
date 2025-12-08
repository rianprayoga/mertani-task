package model

type CreateSensorReq struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type CreateSensorRes struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
