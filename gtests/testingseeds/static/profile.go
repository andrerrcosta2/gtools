// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package static

import "time"

// Profile represents a user's profile with different types of fields
type Profile struct {
	FullName     string
	Age          int
	Email        string
	PhoneNumbers []string
	DateOfBirth  time.Time
	Address      Address
	Preferences  map[string]bool
	Friends      []*Profile
	ExtraInfo    map[string]interface{}
	LastLogin    time.Time
}

func (p *Profile) GetFullName() string {
	return p.FullName
}

func (p *Profile) GetAge() int {
	return p.Age
}

func (p *Profile) GetEmail() string {
	return p.Email
}

func (p *Profile) GetPhoneNumbers() []string {
	return p.PhoneNumbers
}

func (p *Profile) GetDateOfBirth() time.Time {
	return p.DateOfBirth
}

func (p *Profile) GetAddress() Address {
	return p.Address
}

func (p *Profile) GetPreferences() map[string]bool {
	return p.Preferences
}

func (p *Profile) GetFriends() []*Profile {
	return p.Friends
}

func (p *Profile) GetExtraInfo() map[string]interface{} {
	return p.ExtraInfo
}

func (p *Profile) GetLastLogin() time.Time {
	return p.LastLogin
}

// ProfileAsValue Constructor for creating a sample Profile
func ProfileAsValue(fullName, email string, age int, phoneNumbers []string, address Address) Profile {
	return Profile{
		FullName:     fullName,
		Age:          age,
		Email:        email,
		PhoneNumbers: phoneNumbers,
		DateOfBirth:  time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Address:      address,
		Preferences:  map[string]bool{"newsletter": true, "sms_alerts": false},
		Friends:      []*Profile{}, // Add references to other profiles if needed
		ExtraInfo:    map[string]interface{}{"hobbies": []string{"reading", "gaming"}, "profession": "Engineer"},
		LastLogin:    time.Now(),
	}
}

// ProfileAsPointer Constructor for creating a sample Profile
func ProfileAsPointer(fullName, email string, age int, phoneNumbers []string, address Address) *Profile {
	return &Profile{
		FullName:     fullName,
		Age:          age,
		Email:        email,
		PhoneNumbers: phoneNumbers,
		DateOfBirth:  time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Address:      address,
		Preferences:  map[string]bool{"newsletter": true, "sms_alerts": false},
		Friends:      []*Profile{}, // Add references to other profiles if needed
		ExtraInfo:    map[string]interface{}{"hobbies": []string{"reading", "gaming"}, "profession": "Engineer"},
		LastLogin:    time.Now(),
	}
}
