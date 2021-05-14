package locker_system

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/amazon-locker-design/internal/enum/size"
	"github.com/amazon-locker-design/internal/enum/status"
	"github.com/amazon-locker-design/internal/item"
	"github.com/amazon-locker-design/internal/locker"
)

var (
	lockersSizeStatusMap = make(map[size.Size]map[status.Status]map[locker.ILocker]struct{})
	codeLockerMap        = make(map[string]codeSystem) //assuming these code are non-expired
	//for expiring the code a cron will run to delete from this map
	lockerItemMap = make(map[int]item.IItem)
)

type codeSystem struct {
	assignedDate time.Time
	locker       locker.ILocker
}

type LockerSystem struct {
	lockers []locker.ILocker
}

func NewLockerSystem(lockers []locker.ILocker) *LockerSystem {
	for _, sz := range []size.Size{size.SMALL, size.MEDIUM, size.LARGE, size.EXTRALARGE} {
		lockersSizeStatusMap[sz] = make(map[status.Status]map[locker.ILocker]struct{})
		for _, status := range []status.Status{status.FREE, status.OCCUPIED, status.RESERVED} {
			lockersSizeStatusMap[sz][status] = make(map[locker.ILocker]struct{})
		}
	}
	for _, locker := range lockers {
		lockersSizeStatusMap[locker.GetSize()][locker.GetStatus()][locker] = struct{}{}
	}
	return &LockerSystem{lockers: lockers}
}

func getLockerBySizeAndStatus(sz size.Size, st status.Status) (locker.ILocker, error) {
	if lockerByStatus, szExists := lockersSizeStatusMap[sz]; szExists {
		if lockersMap, freeExists := lockerByStatus[st]; freeExists && len(lockersMap) > 0 {
			for locker := range lockersMap {
				return locker, nil
			}
		}
	}
	return nil, errors.New("unable to find locker")
}

func (l LockerSystem) FindLockerForItem(item item.IItem) (locker.ILocker, error) {
	switch item.GetSize() {
	case size.SMALL:
		for _, sz := range []size.Size{size.SMALL, size.MEDIUM, size.LARGE, size.EXTRALARGE} {
			if locker, err := getLockerBySizeAndStatus(sz, status.FREE); err == nil {
				locker.UpdateStatus(status.RESERVED)
				lockersSizeStatusMap[sz][status.RESERVED][locker] = struct{}{}
				delete(lockersSizeStatusMap[sz][status.FREE], locker)
				return locker, nil
			}
		}
	case size.MEDIUM:
		for _, sz := range []size.Size{size.MEDIUM, size.LARGE, size.EXTRALARGE} {
			if locker, err := getLockerBySizeAndStatus(sz, status.FREE); err == nil {
				locker.UpdateStatus(status.RESERVED)
				lockersSizeStatusMap[sz][status.RESERVED][locker] = struct{}{}
				delete(lockersSizeStatusMap[sz][status.FREE], locker)
				return locker, nil
			}
		}
	case size.LARGE:
		for _, sz := range []size.Size{size.LARGE, size.EXTRALARGE} {
			if locker, err := getLockerBySizeAndStatus(sz, status.FREE); err == nil {
				locker.UpdateStatus(status.RESERVED)
				lockersSizeStatusMap[sz][status.RESERVED][locker] = struct{}{}
				delete(lockersSizeStatusMap[sz][status.FREE], locker)
				return locker, nil
			}
		}
	case size.EXTRALARGE:
		if locker, err := getLockerBySizeAndStatus(size.EXTRALARGE, status.FREE); err == nil {
			locker.UpdateStatus(status.RESERVED)
			lockersSizeStatusMap[size.EXTRALARGE][status.RESERVED][locker] = struct{}{}
			delete(lockersSizeStatusMap[size.EXTRALARGE][status.FREE], locker)
			return locker, nil
		}

	}
	return nil, errors.New("unable to find locker")
}

func (l LockerSystem) AddItem(item item.IItem, locker locker.ILocker, code string) error {
	if codeData, exists := codeLockerMap[code]; exists {
		if codeData.locker.GetId() == locker.GetId() {
			//assign to locker item map
			lockerItemMap[locker.GetId()] = item
			//update locker status
			locker.UpdateStatus(status.OCCUPIED)
			//update lockersSizeStatusMap
			delete(lockersSizeStatusMap[locker.GetSize()][status.RESERVED], locker)
			lockersSizeStatusMap[locker.GetSize()][status.OCCUPIED][locker] = struct{}{}
			//expire code
			delete(codeLockerMap, code)
			return nil
		}

		return errors.New("incorrect code locker combination")
	}
	return errors.New("code has expired")
}

func (l LockerSystem) RemoveItem(locker locker.ILocker, code string) (item.IItem, error) {
	if codeData, exists := codeLockerMap[code]; exists {
		if codeData.locker.GetId() == locker.GetId() {
			item := lockerItemMap[locker.GetId()]
			//delete from lockerItemMap
			delete(lockerItemMap, locker.GetId())
			//update locker status
			locker.UpdateStatus(status.FREE)
			//update lockersSizeStatusMap
			delete(lockersSizeStatusMap[locker.GetSize()][status.OCCUPIED], locker)
			lockersSizeStatusMap[locker.GetSize()][status.FREE][locker] = struct{}{}
			//expire code
			delete(codeLockerMap, code)
			return item, nil
		}

		return nil, errors.New("incorrect code locker combination")
	}
	return nil, errors.New("code has expired")
}

func (l LockerSystem) GenerateCode(locker locker.ILocker) string {
	code := uuid.New().String()
	//add code to codeLockerMap
	codeLockerMap[code] = codeSystem{
		assignedDate: time.Now(),
		locker:       locker,
	}
	return code
}
