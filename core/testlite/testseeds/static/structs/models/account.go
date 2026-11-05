// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/charsets"
	"github.com/google/uuid"
	"time"
)

var AccountZeroInst = new(Account)

// AccountAsValue Constructor function to return Account as a value
func AccountAsValue(id, username, passwordHash string, profile Profile) Account {
	return Account{
		ID:               id,
		Username:         username,
		PasswordHash:     passwordHash,
		Profile:          profile,
		TwoFactorEnabled: true,
		Settings:         map[string]interface{}{"theme": "dark", "notifications": true},
		Active:           true,
		JoinedAt:         time.Now().AddDate(-1, 0, 0), // 1 year ago
		LastModified:     time.Now(),
	}
}

// AccountAsPointer Constructor function to return Account as a pointer
func AccountAsPointer(id, username, passwordHash string, profile Profile) *Account {
	return &Account{
		ID:               id,
		Username:         username,
		PasswordHash:     passwordHash,
		Profile:          profile,
		TwoFactorEnabled: true,
		Settings:         map[string]interface{}{"theme": "dark", "notifications": true},
		Active:           true,
		JoinedAt:         time.Now().AddDate(-1, 0, 0), // 1 year ago
		LastModified:     time.Now(),
	}
}

// AccountAsRandValue Function to generate a random Profile as a value
func AccountAsRandValue() Account {
	hash := sha256.Sum256([]byte(random.StringOf(10, charsets.AlphaNumeric+charsets.SafeSymbols).
		At(0)))
	joinedAt := random.Timestamp(1, time.Now().AddDate(-10, 0, 0), time.Now()).At(0)
	return Account{
		ID:               uuid.NewString(),
		Username:         random.Alphanumeric(1, 3, 15).At(0),
		PasswordHash:     hex.EncodeToString(hash[:]),
		Profile:          ProfileAsRandValue(),
		TwoFactorEnabled: random.SingleOf[bool](),
		Settings:         random.SingleOf[map[string]any](),
		Active:           random.SingleOf[bool](),
		JoinedAt:         joinedAt,
		LastModified:     random.Timestamp(1, joinedAt, time.Now()).At(0),
	}
}

func AccountAsRandRef() *Account {
	randomUUID := uuid.New()
	hash := sha256.Sum256([]byte(randomUUID.String()))
	joinedAt := random.Timestamp(1, time.Now().AddDate(-10, 0, 0), time.Now()).At(0)
	return &Account{
		ID:               uuid.NewString(),
		Username:         random.Alphanumeric(1, 3, 15).At(0),
		PasswordHash:     hex.EncodeToString(hash[:]),
		Profile:          ProfileAsRandValue(),
		TwoFactorEnabled: random.SingleOf[bool](),
		Settings:         random.SingleOf[map[string]any](),
		Active:           random.SingleOf[bool](),
		JoinedAt:         joinedAt,
		LastModified:     random.Timestamp(1, joinedAt, time.Now()).At(0),
	}
}

// Account represents the main struct to hold user account information
type Account struct {
	ID               string
	Username         string
	PasswordHash     string
	Profile          Profile
	TwoFactorEnabled bool
	Settings         map[string]interface{}
	Active           bool
	JoinedAt         time.Time
	LastModified     time.Time
}
