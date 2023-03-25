package imageModel

type CreateImagePayload struct {
	Alt   string `json:"alt"`
	Order int    `json:"order"`
	Src   string `json:"src"`
}

func NewCreateImagePayload() CreateImagePayload {
	return CreateImagePayload{}
}

func (payload CreateImagePayload) WithAlt(alt string) CreateImagePayload {
	payload.Alt = alt
	return payload
}

func (payload CreateImagePayload) WithOrder(order int) CreateImagePayload {
	payload.Order = order
	return payload
}

func (payload CreateImagePayload) WithSrc(src string) CreateImagePayload {
	payload.Src = src
	return payload
}