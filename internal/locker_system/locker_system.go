package locker_system

import (
	"github.com/amazon-locker-design/internal/item"
	"github.com/amazon-locker-design/internal/locker"
)

type ILockerSystem interface {
	FindLockerForItem(item item.IItem) (locker.ILocker, error)
	AddItem(item item.IItem, locker locker.ILocker, code string) error
	RemoveItem(locker locker.ILocker, code string) (item.IItem, error)
	//can be a different class altogether
	GenerateCode(locker locker.ILocker) string
}
