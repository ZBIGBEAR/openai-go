package model

type Moderations struct {
	Id      string              `json:"id"`
	Model   string              `json:"model"`
	Results []ModerationsResult `json:"results"`
}

type ModerationsResult struct {
	Flagged        bool           `json:"flagged"`
	Categories     Categories     `json:"categories"`
	CategoryScores CategoryScores `json:"category_scores"`
}

type Categories struct {
	Sexual          bool `json:"sexual"`
	Hate            bool `json:"hate"`
	Violence        bool `json:"violence"`
	SelfHarm        bool `json:"self-harm"`
	SexualMinors    bool `json:"sexual/minors"`
	HateThreatening bool `json:"hate/threatening"`
	ViolenceGraphic bool `json:"violence/graphic"`
}

type CategoryScores struct {
	Sexual          float64 `json:"sexual"`
	Hate            float64 `json:"hate"`
	Violence        float64 `json:"violence"`
	SelfHarm        float64 `json:"self-harm"`
	SexualMinors    float64 `json:"sexual/minors"`
	HateThreatening float64 `json:"hate/threatening"`
	ViolenceGraphic float64 `json:"violence/graphic"`
}
