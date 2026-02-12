package usecase_chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"headliner-be/model/chat" 
)

type ChatUsecase interface {
	AskAI(req model_chat.ChatRequest) (model_chat.ChatResponse, error)
}

type chatUsecase struct{}

func NewChatUsecase() ChatUsecase {
	return &chatUsecase{}
}

func (u *chatUsecase) AskAI(req model_chat.ChatRequest) (model_chat.ChatResponse, error) {
	pythonURL := "http://host.docker.internal:8000/ask"

	payload := model_chat.PythonRequest{
		Question: req.Message,
		PlayerID: req.PlayerID,
	}
	jsonData, _ := json.Marshal(payload)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Post(pythonURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return model_chat.ChatResponse{}, fmt.Errorf("ติดต่อ Python ไม่ได้: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return model_chat.ChatResponse{}, fmt.Errorf("Python Error Code: %d", resp.StatusCode)
	}

	var pyResp model_chat.PythonResponse
	if err := json.NewDecoder(resp.Body).Decode(&pyResp); err != nil {
		return model_chat.ChatResponse{}, err
	}

	return model_chat.ChatResponse{
		Answer: pyResp.Answer,
		Status: "success",
	}, nil
}