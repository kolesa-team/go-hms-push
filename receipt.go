package hms

import "encoding/json"

type Status struct {
	BiTag     string `json:"biTag"`
	RequestId string `json:"requestId"`
	Appid     string `json:"appId"`
	Token     string `json:"token"`
	Status    int    `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

type Receipt struct {
	Statuses []Status `json:"statuses"`
}

func ExtractReceiptFromBody(body []byte) (*Receipt, error) {
	receipt := Receipt{}
	err := json.Unmarshal(body, &receipt)

	if err != nil {
		return nil, err
	}

	return &receipt, nil
}
