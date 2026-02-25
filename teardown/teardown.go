package main

type UserInterface interface {
	attack(user *User)
	getDamage(damage int)
	isDie() bool
}

type User struct {
	ID      int
	Name    string
	Healpth int
	Damage  int
	UserInterface
}

func (u *User) getDamage(damage int) {
	u.Healpth = u.Healpth - damage
}

func (u *User) isDie() bool {
	return u.Healpth <= 0
}

func (u *User) attack(targetUser *User) {
	targetUser.getDamage(u.Damage)
}
