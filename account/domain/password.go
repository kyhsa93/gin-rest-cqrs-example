package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// PasswordAnemic anemic password model
type PasswordAnemic struct {
	Hashed     string
	Cost       int
	CreatedAt  time.Time
	ComparedAt time.Time
}

type passwordOptionsImplement struct {
	hashed string
	cost   int
}

type passwordOptions interface {
	Hashed() string
	Cost() int
}

func (p *passwordOptionsImplement) Hashed() string {
	return p.hashed
}

func (p *passwordOptionsImplement) Cost() int {
	return p.cost
}

// PasswordImplement password domain model struct
type PasswordImplement struct {
	hashed     string
	cost       int
	createdAt  time.Time
	comparedAt time.Time
}

// Password password domain model interface
type Password interface {
	ToAnemic() PasswordAnemic
	Compare(password string) error
}

// NewPassword create new password domain object instance
func NewPassword(options passwordOptions) Password {
	return &PasswordImplement{
		hashed:     options.Hashed(),
		cost:       options.Cost(),
		createdAt:  time.Now(),
		comparedAt: time.Now(),
	}
}

// ReconstitutePassword reconstitute password
func ReconstitutePassword(anemic PasswordAnemic) Password {
	return &PasswordImplement{
		hashed:     anemic.Hashed,
		cost:       anemic.Cost,
		createdAt:  anemic.CreatedAt,
		comparedAt: anemic.ComparedAt,
	}
}

// ToAnemic return anemic password model
func (p *PasswordImplement) ToAnemic() PasswordAnemic {
	return PasswordAnemic{
		Hashed:     p.hashed,
		Cost:       p.cost,
		CreatedAt:  p.createdAt,
		ComparedAt: p.comparedAt,
	}
}

// Compare compare password
func (p *PasswordImplement) Compare(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.hashed), []byte(password))
}
