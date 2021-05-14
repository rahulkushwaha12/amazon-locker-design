package agent

import (
	"github.com/amazon-locker-design/internal/item"
	"github.com/amazon-locker-design/internal/locker"
	"github.com/amazon-locker-design/internal/locker_system"
	"github.com/amazon-locker-design/internal/user"
)

type Agent struct {
	id         int
	lockerSys  locker_system.ILockerSystem
	masterCode string
}

func NewAgent(lockerSys locker_system.ILockerSystem) *Agent {
	return &Agent{lockerSys: lockerSys}
}

func (a Agent) ExecuteDelivery(item item.IItem, locker locker.ILocker, user user.IUser) error {
	err := a.lockerSys.AddItem(item, locker, a.masterCode)
	if err != nil {
		return err
	}
	//generate code or user
	_ = a.lockerSys.GenerateCode(locker)
	//send this code to user vis sms
	return nil

}
