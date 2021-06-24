package muxtesting

type ResponseProduct struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}
