package models

type Data struct {
	ID   string            `json:"id" validate:"required"`
	Data map[string]string `json:"data" validate:"required"`
}

type DataField struct {
	Found bool   `json:"found"`
	Value string `json:"value"`
}

type ResponseData struct {
	ID   string               `json:"id" validate:"required"`
	Data map[string]DataField `json:"data" validate:"required"`
}

type Token struct {
	Username string `json:"Username" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

type Config struct {
	DataEncryptionKey  string `json:"DataEncryptionKey" validate:"required"`
	DataEncryptionIV   string `json:"DataEncryptionIV" validate:"required"`
	TokenEncryptionKey string `json:"TokenEncryptionKey" validate:"required"`
}
