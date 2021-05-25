package storage

import (
	"inventory/internal/collector"
)

type MemoryStorage struct {
	st map[string]*collector.HostInfo
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		st: make(map[string]*collector.HostInfo),
	}
}

func (s *MemoryStorage) Save(hi collector.HostInfo) error {
	s.st[hi.SerialNumber] = &hi
	return nil
}

func (s *MemoryStorage) GetAll() ([]collector.HostInfo, error) {
	var hil []collector.HostInfo
	for _, v := range s.st {
		hil = append(hil, *v)
	}

	return hil, nil
}

func (s *MemoryStorage) IsExist(hi collector.HostInfo) (bool, error) {
	_, exist := s.st[hi.SerialNumber]
	return exist, nil
}

func (s *MemoryStorage) SearchByWH(wh string) (*collector.HostInfo, error) {
	hil, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	for _, hi := range hil {
		if hi.WH == wh {
			return &hi, nil
		}
	}

	return nil, nil
}

func (s *MemoryStorage) SearchBySerial(serial string) (*collector.HostInfo, error) {
	hi := s.st[serial]
	return hi, nil
}
