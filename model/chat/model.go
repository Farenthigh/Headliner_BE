package model_chat

type ChatRequest struct {
	PlayerID string `json:"player_id"`
	Message  string `json:"message"`
}

type ChatResponse struct {
	Answer string `json:"answer"`
	Status string `json:"status"`
}

type PythonRequest struct {
	Question string `json:"question"`
	PlayerID string `json:"player_id"`
}

type PythonResponse struct {
	Answer string `json:"answer"`
}