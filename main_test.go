package main

import (
	"testing"

	"github.com/poy/onpar"
)

func TestSpecs(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	// Connection Test
	o.BeforeEach(func(t *testing.T) (*testing.T, *DB) {
		db := DB{
			ConnectionString: "",
		}

		return t, &db
	})

	o.AfterEach(func(t *testing.T, db *DB) {
		// ...
	})

	o.Spec("something informative", func(t *testing.T, db *DB) {
		if db.DB != nil {
			t.Error("Database connection was nil!!")
		}
	})

	// Unmarshal Tests
	// o.BeforeEach(func(t *testing.T) (*testing.T, int, float64) {
	// 	return nil, 0, 0
	// })

	// o.AfterEach(func(t *testing.T, a int, b float64) {
	// 	// ...
	// })

	// o.Spec("something informative", func(t *testing.T, a int, b float64) {
	// 	if a != 99 {
	// 		t.Errorf("%d != 99", a)
	// 	}
	// })

}
