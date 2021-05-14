package locker

import (
	"github.com/amazon-locker-design/internal/enum/size"
	"github.com/amazon-locker-design/internal/enum/status"
)

type Locker struct {
	id     int
	size   size.Size
	status status.Status
}

func NewLocker(size size.Size, status status.Status) *Locker {
	return &Locker{size: size, status: status}
}

func (l *Locker) UpdateStatus(status status.Status) error {
	l.status = status
	return nil
}

func (l Locker) GetSize() size.Size {
	return l.size
}

func (l Locker) GetStatus() status.Status {
	return l.status
}

func (l Locker) GetId() int {
	return l.id
}
