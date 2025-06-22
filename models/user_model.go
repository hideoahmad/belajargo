package models

type User struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Name      string `json:"name"`
    Email     string `gorm:"unique" json:"email"`
    Password  string `json:"-"`
    Phone     string `json:"phone"`
    Address   string `json:"address"`
    Age       int    `json:"age"`
    Education string `json:"education"`
}
