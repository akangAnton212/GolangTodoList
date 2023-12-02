package helpers

type ResponseApi struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseData(status bool, message string, data any) ResponseApi {
	responseApi := ResponseApi{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return responseApi
}
