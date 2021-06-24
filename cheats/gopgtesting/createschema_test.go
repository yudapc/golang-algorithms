package gopgtesting

import "testing"

func TestCreateSchemas(t *testing.T) {
	db := NewDBConnection()
	CreateSchemas(db)
}

func TestCreateSingleSchema(t *testing.T) {
	db := NewDBConnection()
	CreateSingleSchema(db)
}
