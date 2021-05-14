package agent

import (
	"github.com/amazon-locker-design/internal/item"
	"github.com/amazon-locker-design/internal/locker"
	"github.com/amazon-locker-design/internal/user"
)

type IAgent interface {
	ExecuteDelivery(item item.IItem, locker locker.ILocker, user user.IUser) error
}
