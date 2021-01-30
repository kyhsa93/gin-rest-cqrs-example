package domain

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// AccountAnemic anemic account model
type AccountAnemic interface {
	ID() string
	Name() string
	Passord() PasswordAnemic
	Balance() int
	OpenedAt() time.Time
	UpdatedAt() time.Time
	ClosedAt() time.Time
}

type accountAnemicImplement struct {
	id        string
	name      string
	password  PasswordAnemic
	balance   int
	openedAt  time.Time
	updatedAt time.Time
	closedAt  time.Time
}

func (a *accountAnemicImplement) ID() string {
	return a.id
}

func (a *accountAnemicImplement) Name() string {
	return a.name
}

func (a *accountAnemicImplement) Passord() PasswordAnemic {
	return a.password
}

func (a *accountAnemicImplement) Balance() int {
	return a.balance
}

func (a *accountAnemicImplement) OpenedAt() time.Time {
	return a.openedAt
}

func (a *accountAnemicImplement) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *accountAnemicImplement) ClosedAt() time.Time {
	return a.closedAt
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
	return &accountAnemicImplement{
		id:        a.id,
		name:      a.name,
		password:  a.password.ToAnemic(),
		balance:   a.balance,
		openedAt:  a.openedAt,
		updatedAt: a.updatedAt,
		closedAt:  a.closedAt,
	}
}

// AccountRepository account repository
type AccountRepository interface {
	Save(account Account) error
	FindNewID() (string, error)
	FindByID(id string) (Account, error)
}
