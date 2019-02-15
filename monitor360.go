package monitor360

type MonitorRequest struct {
	Result struct {
		Intent struct {
			DisplayName string `json:"displayName"`
		} `json:"intent"`
	} `json:"queryResult"`
}

type MonitorResponse struct {
	Payload payload `json:"payload"`
}

type payload struct {
	Google googlePayload `json:"google"`
}

type googlePayload struct {
	ExpectUserResponse bool         `json:"expectUserResponse"`
	RichResponse       richResponse `json:"richResponse"`
}

type richResponse struct {
	Items []richResponseItem `json:"items"`
}

type richResponseItem struct {
	SimpleResponse simpleResponse `json:"simpleResponse"`
}

type simpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}

func NewResponse(expectUserResponse bool, response string) MonitorResponse {
	return MonitorResponse{
		Payload: payload{
			Google: googlePayload{
				ExpectUserResponse: expectUserResponse,
				RichResponse: richResponse{
					Items: []richResponseItem{
						richResponseItem{
							SimpleResponse: simpleResponse{
								TextToSpeech: response,
							},
						},
					},
				},
			},
		},
	}
}
