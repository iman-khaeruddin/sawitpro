package entity

type User struct {
	ID           int    `gorm:"primaryKey;type:INT UNSIGNED" json:"id"`
	FullName     string `json:"fullName"`
	PhoneNumber  string `gorm:"default:null" json:"phoneNumber"`
	Password     string `json:"password"`
	LoginAttempt int    `json:"loginAttempt"`
}

func (User) TableName() string {
	return "users"
}
