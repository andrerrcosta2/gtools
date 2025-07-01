// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import "github.com/andrerrcosta2/gtools/core/seeders/random"

type Credential struct {
	Username string
	Password string
}

func CredentialAsValue(username string, password string) Credential {
	return Credential{
		Username: username,
		Password: password,
	}
}

func CredentialAsPointer(username string, password string) *Credential {
	return &Credential{
		Username: username,
		Password: password,
	}
}

func CredentialAsRandValue() Credential {
	arr := random.Alphanumeric(2, 2, 30)
	return Credential{
		Username: arr.At(0),
		Password: arr.At(1),
	}
}

func CredentialAsRandRef() *Credential {
	arr := random.Alphanumeric(2, 2, 30)
	return &Credential{
		Username: arr.At(0),
		Password: arr.At(1),
	}
}

// User represents a user with a name, age, and credentials.
type User struct {
	Name        string
	Age         int
	Credentials map[string]Credential
}

func UserAsValue(name string, age int, credentials map[string]Credential) User {
	return User{
		Name:        name,
		Age:         age,
		Credentials: credentials,
	}
}

func UserAsPointer(name string, age int, credentials map[string]Credential) *User {
	return &User{
		Name:        name,
		Age:         age,
		Credentials: credentials,
	}
}

func UserAsRandValue() User {
	return User{
		Name: random.Alphabet(1, 3, 30).At(0),
		Age:  random.Int(1, 1, 100).At(0),
		Credentials: map[string]Credential{
			"email": {
				Username: random.Alphanumeric(1, 2, 30).At(0),
				Password: random.Alphanumeric(1, 2, 30).At(0),
			},
		},
	}
}

func UserAsRandRef() *User {
	return &User{
		Name: random.Alphabet(1, 3, 30).At(0),
		Age:  random.Int(1, 1, 100).At(0),
		Credentials: map[string]Credential{
			"email": {
				Username: random.Alphanumeric(1, 2, 30).At(0),
				Password: random.Alphanumeric(1, 2, 30).At(0),
			},
		},
	}
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetCredentials() map[string]Credential {
	return u.Credentials
}
