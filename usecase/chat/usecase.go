package usecase_chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"os"
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
	pythonURL := os.Getenv("PYTHON_URL")

	payload := model_chat.PythonRequest{
		Question: req.Message,
		PlayerID: req.PlayerID,
	}

	jsonData, err := json.Marshal(payload)
    if err != nil {
        return model_chat.ChatResponse{}, fmt.Errorf("failed to marshal request payload: %v", err)
    }

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Post(pythonURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return model_chat.ChatResponse{}, fmt.Errorf("fail connect to python (%s): %v", pythonURL, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return model_chat.ChatResponse{}, fmt.Errorf("Python Error Code: %d", resp.StatusCode)
    }

    var pyResp model_chat.PythonResponse
    if err := json.NewDecoder(resp.Body).Decode(&pyResp); err != nil {
        return model_chat.ChatResponse{}, fmt.Errorf("failed to decode python response: %v", err)
    }

	return model_chat.ChatResponse{
		Answer: pyResp.Answer,
		Status: "success",
	}, nil
}