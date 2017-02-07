package tool

import "github.com/zheng-ji/goSnowFlake"

func GenNewID() (int64, error) {
	iw, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		return 0, err
	}
	if id, err := iw.NextId(); err != nil {
		return 0, err
	} else {
		return id, nil
	}
}

func NewID() int64 {
	newid, _ := GenNewID()
	return newid
}

func NewStrID() string {
	return ToString(NewID())
}
