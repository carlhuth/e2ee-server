package db

import (
	"github.com/jinzhu/gorm"
)

type Accounts []Account

type UserCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type User struct {
        gorm.Model
        Username             string
        HashedPassword       string
        Uuid                 string
        Token                string `sql:"size:600"`
}

type Account struct {
	ContainerNameHmacKeyCiphertext        string       `json:"containerNameHmacKeyCiphertext" sql:"type:text"`
	HmacKeyCiphertext      string    `json:"hmacKeyCiphertext" sql:"type:text"`
	KeypairCiphertext string      `json:"keypairCiphertext" sql:"type:text"`
	KeypairMac string      `json:"keypairMac" sql:"type:text"`
	KeypairMacSalt string      `json:"keypairMacSalt" sql:"type:text"`
	KeypairSalt string      `json:"keypairSalt" sql:"type:text"`
	PubKey string      `json:"pubKey" sql:"type:text"`
	SignKeyPrivateCiphertext string      `json:"signKeyPrivateCiphertext" sql:"type:text"`
	SignKeyPrivateMac string      `json:"signKeyPrivateMac" sql:"type:text"`
	SignKeyPrivateMacSalt string      `json:"signKeyPrivateMacSalt" sql:"type:text"`
	SignKeyPub string      `json:"signKeyPub" sql:"type:text"`
	Username string      `json:"username"`
	AccountId uint      `json:"accountId"` // ID of user
}


