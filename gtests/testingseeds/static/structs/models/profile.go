// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package models

import (
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"time"
)

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

func ProfileAsRandValue() Profile {
	now := time.Now()
	profile := Profile{
		FullName: random.Alphabet(1, 3, 15).At(0) + " " +
			random.Alphabet(1, 3, 15).At(0),
		Age:   random.Int(1, 1, 100).At(0),
		Email: random.Alphabet(1, 3, 15).At(0) + "@example.com",
		PhoneNumbers: random.StringOf(random.Int(1, 1, 10).At(0), "1234567890", 8, 9).
			Values(),
		DateOfBirth: random.Timestamp(1, now.AddDate(-100, 0, 0),
			now.AddDate(-10, 0, 0)).At(0),
		Address:     AddressAsRandValue(),
		Preferences: random.Of[map[string]bool](1).At(0),
		ExtraInfo:   random.Of[map[string]any](1).At(0),
		LastLogin:   random.Timestamp(1, now.AddDate(-2, 0, 0), now).At(0),
	}

	// Generate a random profile with various fields
	size := random.Int(1, 1, 10).At(0)
	friends := make([]*Profile, size)
	for i := 0; i < size; i++ {
		friends = append(friends, randFriend(&profile))
	}
	profile.Friends = friends
	return profile
}

func randFriend(of *Profile) *Profile {
	return &Profile{
		FullName:     random.Alphabet(1, 3, 15).At(0) + " " + random.Alphabet(1, 3, 15).At(0),
		Age:          random.Int(1, 1, 100).At(0),
		Email:        random.Alphabet(1, 3, 15).At(0) + "@example.com",
		PhoneNumbers: random.StringOf(random.Int(1, 1, 3).At(0), "1234567890", 8, 9).Values(),
		DateOfBirth:  random.Timestamp(1, time.Now().AddDate(-100, 0, 0)).At(0),
		Address:      AddressAsRandValue(),
		Preferences:  random.Of[map[string]bool](1).At(0),
		Friends:      []*Profile{of},
	}
}

func ProfileAsRandRef() *Profile {
	now := time.Now()
	profile := Profile{
		FullName: random.Alphabet(1, 3, 15).At(0) + " " +
			random.Alphabet(1, 3, 15).At(0),
		Age:          random.Int(1, 1, 100).At(0),
		Email:        random.Alphabet(1, 3, 15).At(0) + "@example.com",
		PhoneNumbers: random.StringOf(random.Int(1, 1, 3).At(0), "1234567890", 8, 9).Values(),
		DateOfBirth: random.Timestamp(1, now.AddDate(-100, 0, 0),
			now.AddDate(-10, 0, 0)).At(0),
		Address:     AddressAsRandValue(),
		Preferences: random.Of[map[string]bool](1).At(0),
		ExtraInfo:   random.Of[map[string]any](1).At(0),
		LastLogin:   random.Timestamp(1, now.AddDate(-2, 0, 0), now).At(0),
	}

	// Generate a random profile with various fields
	size := random.Int(1, 1, 10).At(0)
	friends := make([]*Profile, size)
	for i := 0; i < size; i++ {
		friends = append(friends, randFriend(&profile))
	}
	profile.Friends = friends
	return &profile
}
