package item

import (
	"github.com/amazon-locker-design/internal/enum/size"
)

type Item struct {
	size size.Size
}

func NewItem(size size.Size) *Item {
	return &Item{size: size}
}

func (i Item) GetSize() size.Size {
	return i.size
}
