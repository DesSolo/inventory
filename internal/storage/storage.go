package storage

import "inventory/internal/collector"

type Storage interface {
	Send(*collector.HostInfo) error
}
