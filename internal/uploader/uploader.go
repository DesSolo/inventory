package uploader

import "inventory/internal/collector"

type Uploader interface {
	Upload(*collector.HostInfo) error
	IsExist(string) (bool, error)
}
