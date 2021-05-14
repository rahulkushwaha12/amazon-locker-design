package item

import (
	"github.com/amazon-locker-design/internal/enum/size"
)

type IItem interface {
	GetSize() size.Size
}
