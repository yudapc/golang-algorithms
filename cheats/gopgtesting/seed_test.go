package gopgtesting

import (
	"testing"
)

func TestSeed(t *testing.T) {
	db := NewDBConnection()
	i := 0
	for i < 5 {
		SeedData(db)
		i++
	}
	users := &Users{}
	db.Model(users).Select()
	t.Log(users)
}
