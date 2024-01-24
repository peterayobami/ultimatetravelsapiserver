package datamodels

type AmadeusCredential struct {
	BaseDataModel `gorm:"embedded"`
	Type          string
	Username      string
	ClientId      string
	TokenType     string
	AccessToken   string
	ExpiresIn     int
	State         string
	Scope         string
}
