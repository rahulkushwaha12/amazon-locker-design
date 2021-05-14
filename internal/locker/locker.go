package locker

import (
	"github.com/amazon-locker-design/internal/enum/size"
	"github.com/amazon-locker-design/internal/enum/status"
)

type ILocker interface {
	UpdateStatus(status status.Status) error
	GetSize() size.Size
	GetStatus() status.Status
	GetId() int
}
