package auth

var (
	// GetIAuth is a function that is call to obtain a instance of the
	// IAuth use by this application
	GetIAuth = func() IAuth {
		return nil
	}
)

// IAuth is the interface for authentification and user management
type IAuth interface {
	LoginUser(username string, password string) (*User, error)

	IsValidUser(id int, hash string) error

	UpdateAccountPassword(id int, oldhash string, newpw string) error

	CreateAccount(username string, password string, role Role) (*User, error)

	DeleteAccount(id int) error

	GetListAccount() ([]*User, error)

	Close()
}
