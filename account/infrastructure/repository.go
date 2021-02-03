package infrastructure

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/domain"
)

func test(container container) {

}

// NewRepository create new repository
func NewRepository(connection *gorm.DB) domain.AccountRepository {
	return &MySQL{DB: connection}
}

// MySQL mysql
type MySQL struct {
	*gorm.DB
}

// Save save given domain object
func (m *MySQL) Save(account domain.Account) {
	m.DB.Save(convertDomainToEntity(account))
}

// FindNewID find new domain object id
func (m *MySQL) FindNewID() (string, error) {
	newID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	m.DB.Save(accountEntity{id: newID.String()})
	return newID.String(), nil
}

// FindByID find domain object using id
func (m *MySQL) FindByID(id string) (domain.Account, error) {
	var entity accountEntity
	if err := m.DB.First(&entity).Error; err != nil {
		return nil, err
	}

	return convertEntityToDomain(entity), nil
}

func convertEntityToDomain(entity accountEntity) domain.Account {
	return domain.ReconstituteAccount(domain.AccountAnemic{
		ID:   entity.id,
		Name: entity.name,
		Password: domain.PasswordAnemic{
			Hashed:     entity.password.hashed,
			Cost:       entity.password.cost,
			CreatedAt:  entity.password.createdAt,
			ComparedAt: entity.password.comparedAt,
		},
	})
}

func convertDomainToEntity(account domain.Account) accountEntity {
	return accountEntity{
		id:   account.ToAnemic().ID,
		name: account.ToAnemic().Name,
		password: passwordEntity{
			accountID:  account.ToAnemic().ID,
			hashed:     account.ToAnemic().Password.Hashed,
			cost:       account.ToAnemic().Password.Cost,
			createdAt:  account.ToAnemic().Password.CreatedAt,
			comparedAt: account.ToAnemic().Password.ComparedAt,
		},
		balance:   account.ToAnemic().Balance,
		openedAt:  account.ToAnemic().OpenedAt,
		updatedAt: account.ToAnemic().UpdatedAt,
		closeAt:   account.ToAnemic().ClosedAt,
	}
}

type accountEntity struct {
	id        string
	name      string
	password  passwordEntity
	balance   int
	openedAt  time.Time
	updatedAt time.Time
	closeAt   time.Time
}

type passwordEntity struct {
	accountID  string
	hashed     string
	cost       int
	createdAt  time.Time
	comparedAt time.Time
}
