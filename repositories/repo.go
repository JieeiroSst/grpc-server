package repositories

import (
	"github.com/JIeeiroSst/go-app/inventory"
)

type IventoryRepository interface {
	CheckAndOrder([]*inventory.Item) (bool, error)
}