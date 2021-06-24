package gopgtesting

import (
	"fmt"
)

// User
type User struct {
	ID        string `json:"id" pg:"id,pk,uuid_generate_v4()"`
	FirstName string `json:"first_name" pg:"first_name"`
	LastName  string `json:"last_name" pg:"last_name"`
	Email     string `json:"email" pg:"email,unique"`
}

type Users []User

func (u User) String() string {
	return fmt.Sprintf("User<%s %s %s %s>", u.ID, u.FirstName, u.LastName, u.Email)
}

// Story
type Story struct {
	ID       string `json:"id" pg:"id,pk,uuid_generate_v4()"`
	Title    string `json:"title" pg:"title"`
	AuthorId string `json:"author_id" pg:"author_id,fk"`
	Author   *User  `json:"author" pg:"rel:has-one"`
}

type Stories []Story

func (s Story) String() string {
	return fmt.Sprintf("Story<%s %s %s>", s.ID, s.Title, s.Author)
}
