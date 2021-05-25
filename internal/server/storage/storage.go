package storage

import "inventory/internal/collector"

type Storage interface {
	Save(collector.HostInfo) error
	GetAll() ([]collector.HostInfo, error)
	IsExist(collector.HostInfo) (bool, error)
	SearchByWH(string) (*collector.HostInfo, error)
	SearchBySerial(string) (*collector.HostInfo, error)
}
