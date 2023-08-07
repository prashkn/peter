package entities

type User struct {
	Name               string
	EncryptedPasswords map[string]string
	Email              string
	Validated          bool
	Password           string
}

type Password struct {
	Website  string
	Username string
	Email    string
	Hash     string
}

type Session struct {
	User User
}
