package TaxGameScoreModels

type CreateTaxGameScoreInput struct {
	UserID uint `json:"user_id"`
	Score  int  `json:"score"`
}

type UpdateTaxGameScoreInput struct {
	Id uint `json:"id"`
	UserID uint `json:"user_id"`
	Score  int  `json:"score"`
}