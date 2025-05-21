// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package static

import "time"

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

// GetAccountAsValue Constructor function to return Account as a value
func GetAccountAsValue(id, username, passwordHash string, profile Profile) Account {
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

// GetAccountAsPointer Constructor function to return Account as a pointer
func GetAccountAsPointer(id, username, passwordHash string, profile Profile) *Account {
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
