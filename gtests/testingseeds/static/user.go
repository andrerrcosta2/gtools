// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package static

type Credential struct {
	Username string
	Password string
}

type User struct {
	Name        string
	Age         int
	Credentials map[string]Credential
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

func GetUserAsData(name string, age int, password string) User {
	return User{
		Name: name,
		Age:  age,
		Credentials: map[string]Credential{
			"email": {
				Username: name,
				Password: password,
			},
		},
	}
}

func GetUserAsPointer(name string, age int, password string) *User {
	return &User{
		Name: name,
		Age:  age,
		Credentials: map[string]Credential{
			"email": {
				Username: name,
				Password: password,
			},
		},
	}
}
