package user

import (
	"github.com/amazon-locker-design/internal/item"
	"github.com/amazon-locker-design/internal/locker"
)

type IUser interface {
	OptForLockerDelivery(item item.IItem) (locker.ILocker, error)
	GetItemFromLocker(code string, locker locker.ILocker) (item.IItem, error)
	OptForLockerReturn(item item.IItem) (locker.ILocker, error)
	ReturnItemToLocker(item item.IItem, code string, locker locker.ILocker) error
}
