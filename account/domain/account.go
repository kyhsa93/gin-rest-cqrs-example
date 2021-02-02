package domain

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// AccountAnemic anemic account model
type AccountAnemic struct {
	ID        string
	Name      string
	Password  PasswordAnemic
	Balance   int
	OpenedAt  time.Time
	UpdatedAt time.Time
	ClosedAt  time.Time
}

// AccountOptions options for create account
type AccountOptions struct {
	ID       string
	Name     string
	Password string
}

// AccountImplement account domain model struct
type AccountImplement struct {
	id        string
	name      string
	password  Password
	balance   int
	openedAt  time.Time
	updatedAt time.Time
	closedAt  time.Time
	events    []interface{}
}

// Account account domain model interface
type Account interface {
	comparePasswrod(password string) error
	UpdatePassword(new, password string) error
	Withdraw(amount int, password string) error
	Deposit(amount int) error
	Close(password string) error
	ToAnemic() AccountAnemic
	Events() []interface{}
}

// NewAccount create new account domain object instance
func NewAccount(options AccountOptions) Account {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(options.Password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return &AccountImplement{
		id:   options.ID,
		name: options.Name,
		password: NewPassword(&passwordOptionsImplement{
			hashed: string(hashedByte),
			cost:   bcrypt.MinCost,
		}),
		balance:   0,
		openedAt:  time.Now(),
		updatedAt: time.Now(),
		closedAt:  time.Now(),
		events:    []interface{}{AccountOpenedDomainEventImplement{accountID: options.ID}},
	}
}

// ReconstituteAccount reconstitute account domain object
func ReconstituteAccount(anemic AccountAnemic) Account {
	return &AccountImplement{
		id:   anemic.ID,
		name: anemic.Name,
		password: ReconstitutePassword(PasswordAnemic{
			Hashed:     anemic.Password.Hashed,
			Cost:       anemic.Password.Cost,
			CreatedAt:  anemic.Password.CreatedAt,
			ComparedAt: anemic.Password.ComparedAt,
		}),
		balance:   anemic.Balance,
		openedAt:  anemic.OpenedAt,
		updatedAt: anemic.UpdatedAt,
		closedAt:  anemic.ClosedAt,
	}
}

// Events account domain events
func (a *AccountImplement) Events() []interface{} {
	return a.events
}

// UpdatePassword update account's password
func (a *AccountImplement) UpdatePassword(new, password string) error {
	if err := a.comparePasswrod(password); err != nil {
		return err
	}

	a.updatedAt = time.Now()
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	a.password = NewPassword(&passwordOptionsImplement{
		hashed: string(hashedByte),
		cost:   bcrypt.MinCost,
	})
	a.events = append(a.events, AccountPasswordUpdatedDomainEventImplement{accountID: a.id})
	return nil
}

// Withdraw withdraw amount from account
func (a *AccountImplement) Withdraw(amount int, password string) error {
	if err := a.comparePasswrod(password); err != nil {
		return err
	}

	if amount < 0 {
		return errors.New("Can not withdraw under 0")
	}

	if a.balance < amount {
		return errors.New("Requested amount exceeds your withdrawal limits")
	}

	a.updatedAt = time.Now()
	a.balance = a.balance - amount
	a.events = append(a.events, WithdrawnDomainEventImplement{accountID: a.id})
	return nil
}

// Deposit deposit amount to account
func (a *AccountImplement) Deposit(amount int) error {
	if amount < 0 {
		return errors.New("Can not deposit under 0")
	}

	a.updatedAt = time.Now()
	a.balance = a.balance + amount
	a.events = append(a.events, DepositedDomainEventImplement{accountID: a.id})
	return nil
}

// Close close account
func (a *AccountImplement) Close(password string) error {
	if err := a.comparePasswrod(password); err != nil {
		return err
	}

	a.updatedAt = time.Now()
	a.closedAt = time.Now()
	a.events = append(a.events, AccountClosedDomainEventImplement{accountID: a.id})
	return nil
}

func (a *AccountImplement) comparePasswrod(password string) error {
	return a.password.Compare(password)
}

// ToAnemic convert account to anemic object
func (a *AccountImplement) ToAnemic() AccountAnemic {
	return AccountAnemic{
		ID:        a.id,
		Name:      a.name,
		Password:  a.password.ToAnemic(),
		Balance:   a.balance,
		OpenedAt:  a.openedAt,
		UpdatedAt: a.updatedAt,
		ClosedAt:  a.closedAt,
	}
}

// AccountRepository account repository
type AccountRepository interface {
	Save(account Account)
	FindNewID() (string, error)
	FindByID(id string) (Account, error)
}
