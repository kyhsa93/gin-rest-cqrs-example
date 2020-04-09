package query

// ReadAccountByIDQuery read account query by accountId
type ReadAccountByIDQuery struct {
	AccountID string
}

// ReadAccountQuery read account query
type ReadAccountQuery struct {
	Email    string
	Provider string
	SocialID string
	Password string
	Deleted  bool
}

// ReadAccountByEmailQuery read account by email
type ReadAccountByEmailQuery struct {
	Email string
}
