package model

type EmailDto struct {
	Subject string `json:"subject"`
	Content string `json:"content"`
	Type 	string `json:"type"`
}

type EmailSingleDto struct {
	To 		string `json:"to"`
	EmailDto
}

type EmailGroupDto struct {
	Tos []string `json:"emails"`
	EmailDto
}