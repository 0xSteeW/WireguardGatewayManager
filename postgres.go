package wgm

import (
	"errors"
	"time"

	"github.com/go-pg/pg/v10"
)

func readConfig() {
}

// Temporary placeholder
type User struct {
	id            int
	user_name     string
	user_email    string
	created_at    time.Time
	updated_at    time.Time
	domain_id     int
	unique_id     string
	total_clients int
	auth_provider int
	provider_id   string
	token         string
	role          string
}

type Domains struct {
	id            int
	name          string
	_type         string `pg:type`
	created_at    time.Time
	updated_at    time.Time
	unique_id     string
	total_users   int
	referral_code string
}

func loadPG() {
	db := pg.Connect(&pg.Options{
		User:     "",
		Password: "",
		Database: "",
	})
	// TODO defer elsewhere
	defer db.Close()
	if db == nil {
		panic("could not connect to postgresql")
	}
}

func authorize(db *pg.DB, token string) error {
	user := new(User)
	domains := new(Domains)
	db.Model(user).Where("token = ?", token).Select()
	if user == nil {
		return errors.New("Could not find an user with the requested token")
	}
	db.Model(domains).Where("id = ?", user.domain_id)
	if db == nil {
		return errors.New("Encountered a domain error")
	}
	if domains.id
	return nil

}
