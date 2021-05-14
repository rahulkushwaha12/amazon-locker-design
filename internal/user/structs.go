package user

import (
	"github.com/amazon-locker-design/internal/item"
	"github.com/amazon-locker-design/internal/locker"
	"github.com/amazon-locker-design/internal/locker_system"
)

type User struct {
	id        int
	mobile    string
	lockerSys locker_system.ILockerSystem
}

func NewUser(mobile string, lockerSys locker_system.ILockerSystem) *User {
	return &User{mobile: mobile, lockerSys: lockerSys}
}

func (u User) OptForLockerDelivery(item item.IItem) (locker.ILocker, error) {
	return u.lockerSys.FindLockerForItem(item)
}

func (u User) GetItemFromLocker(code string, locker locker.ILocker) (item.IItem, error) {
	return u.lockerSys.RemoveItem(locker, code)
}

func (u User) OptForLockerReturn(item item.IItem) (locker.ILocker, error) {
	var (
		err    error
		locker locker.ILocker
	)
	if locker, err = u.lockerSys.FindLockerForItem(item); err == nil {
		_ = u.lockerSys.GenerateCode(locker)
		//send this code vis sms to user
		return locker, nil
	}
	return nil, err
}

func (u User) ReturnItemToLocker(item item.IItem, code string, locker locker.ILocker) error {
	return u.lockerSys.AddItem(item, locker, code)
}
