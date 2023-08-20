package onepassword

import "time"

type Item struct {
	ID                    string        `json:"id"`
	Title                 string        `json:"title"`
	Tags                  []string      `json:"tags"`
	Version               int           `json:"version"`
	Vault                 ItemVault     `json:"vault"`
	Category              string        `json:"category"`
	LastEditedBy          string        `json:"last_edited_by"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
	AdditionalInformation string        `json:"additional_information"`
	Urls                  []ItemUrl     `json:"urls"`
	Sections              []ItemSection `json:"sections"`
	Fields                []ItemField   `json:"fields"`
}

type ItemVault struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ItemUrl struct {
	Primary bool   `json:"primary"`
	Href    string `json:"href"`
}

type ItemSection struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type ItemField struct {
	ID    string  `json:"id"`
	Type  string  `json:"type"`
	Label *string `json:"label"`
	Value *string `json:"value"`
	TOTP  string  `json:"totp"`
	// Omit other fields
}
