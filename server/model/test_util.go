package model

import (
	"database/sql"
	"fmt"

	"github.com/blendlabs/spiffy"
	"github.com/wcharczuk/instabot/server/core"
)

// CreateTestUser creates a test user.
func CreateTestUser(tx *sql.Tx) (*User, error) {
	u := NewUser(fmt.Sprintf("__test_user_%s__", core.UUIDv4().ToShortString()))
	u.FirstName = "Test"
	u.LastName = "User"
	err := spiffy.DefaultDb().CreateInTransaction(u, tx)
	return u, err
}

// CreateTestUserAuth creates a test user auth.
func CreateTestUserAuth(userID int64, token, secret string, tx *sql.Tx) (*UserAuth, error) {
	ua := NewUserAuth(userID, token, secret)
	ua.Provider = "test"
	err := spiffy.DefaultDb().CreateInTransaction(ua, tx)
	return ua, err
}

// CreateTestUserSession creates a test user session.
func CreateTestUserSession(userID int64, tx *sql.Tx) (*UserSession, error) {
	us := NewUserSession(userID)
	err := spiffy.DefaultDb().CreateInTransaction(us, tx)
	return us, err
}
