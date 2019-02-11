package auth

import "errors"

type DumpAuth struct {
	filepath string
	Account  []*User
}

func (d *DumpAuth) LoginUser(username string, password string) (*User, error) {
	for _, a := range d.Account {
		if a.Username == username {
			if err := ValidPassword(password, a.SaltedPW); err != nil {
				println("Invalid password")
				return nil, err
			} else {
				return a, nil
			}
		}
	}
	return nil, errors.New("Not found")
}

func (d *DumpAuth) IsValidUser(id int, hash string) error {
	for _, a := range d.Account {
		if a.ID == id {
			return nil
		}
	}

	return errors.New("Not found")
}

func (d *DumpAuth) UpdateAccountPassword(id int, oldhash string, newpw string) error {
	return nil
}

func (d *DumpAuth) CreateAccount(username string, password string, role Role) (*User, error) {
	return nil, nil
}

func (d *DumpAuth) DeleteAccount(id int) error {
	return nil
}

func (d *DumpAuth) GetListAccount() ([]*User, error) {
	return nil, nil
}

func (d *DumpAuth) Close() {
}

func OpenDumpAuth(filepath string) (*DumpAuth, error) {

	return nil, nil
}
