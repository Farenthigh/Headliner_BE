package model

type ContactRequest struct {
    Subject     string `json:"subject" binding:"required"`
    Description string `json:"description" binding:"required"`
}