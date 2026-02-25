package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("Setup tests...")
	start := time.Now()

	m.Run()
	fmt.Println("Teardown tests...")

	end := time.Now()
	fmt.Println("Tests took", end.Sub(start))
}

func TestUsersDamage(t *testing.T) {
	user1 := User{
		ID:      1,
		Name:    "Roma",
		Healpth: 100,
		Damage:  20,
	}

	user2 := User{
		ID:      2,
		Name:    "Nick",
		Healpth: 200,
		Damage:  13,
	}

	t.Run("user1 get damage", func(t *testing.T) {
		want := 180
		user1.attack(&user2)
		if want != user2.Healpth {
			t.Errorf("Want %d, but has %d", want, user2.Healpth)
		}
	})
}
