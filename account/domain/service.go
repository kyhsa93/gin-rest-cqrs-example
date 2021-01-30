package domain

// RemittanceOptions options for remittance
type RemittanceOptions struct {
	Password string
	Sender   Account
	Receiver Account
	Amount   int
}

// AccountDomainServiceImplement account domain service struct
type AccountDomainServiceImplement struct{}

// AccountDomainService account domain service interface
type AccountDomainService interface {
	Remit(options RemittanceOptions) error
}

// Remit remit to receiver account from sender account
func (a *AccountDomainServiceImplement) Remit(options RemittanceOptions) error {
	if err := options.Sender.Withdraw(options.Amount, options.Password); err != nil {
		return err
	}
	if err := options.Receiver.Deposit(options.Amount); err != nil {
		return err
	}
	return nil
}
