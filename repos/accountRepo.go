package repos

import "rovaBankProject/models"

type Repos interface {
	CreateAccount(account models.Account) error
}
