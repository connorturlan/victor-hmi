package auth

type UserData struct {
	UserID              string              `json:"UserID,omitempty"`
	Email               string              `json:"Email,omitempty"`
	GoogleID            string              `json:"GoogleID,omitempty"`
	CurrentSubscription bool                `json:"CurrentSubscription,omitempty"`
	PointsOfInterest    map[string]Landmark `json:"PointsOfInterest,omitempty"`
}

type UserToken struct {
	UserID    string `json:"UserID,omitempty"`
	IssueTime string `json:"IssueTime,omitempty"`
	Token     string `json:"Token,omitempty"`
}

type Landmark struct {
	Name string  `json:"Name"`
	Lat  float64 `json:"Lat"`
	Lng  float64 `json:"Lng"`
}
