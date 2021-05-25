package storage

import "inventory/internal/collector"

type Storage interface {
	Save(collector.HostInfo) error
	IsExist(collector.HostInfo) (bool, error)
	GetAll() ([]collector.HostInfo, error)
	GetByWH(string) (*collector.HostInfo, error)
	GetBySerial(string) (*collector.HostInfo, error)
	DeleteBySerial(string) error
}
