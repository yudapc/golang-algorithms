package gopgtesting

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

func SeedData(db *pg.DB) {
	user, err := SeedDataUser(db)
	if err != nil {
		panic(err)
	}
	seedDataStory(db, user)
}

func SeedDataUser(db *pg.DB) (*User, error) {
	user := &User{
		ID:        uuid.New().String(),
		FirstName: randomdata.FirstName(randomdata.Female),
		LastName:  randomdata.LastName(),
		Email:     randomdata.Email(),
	}
	_, err := db.Model(user).Insert()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func seedDataStory(db *pg.DB, u *User) {
	story := &Story{
		ID:       uuid.New().String(),
		Title:    "Cool story",
		AuthorId: u.ID,
	}
	_, err := db.Model(story).Insert()
	if err != nil {
		panic(err)
	}
}
