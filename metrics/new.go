package metrics

import (
	"fmt"

	"github.com/kris-nova/novaarchive/filesystem"
	"github.com/kris-nova/novaarchive/metrics/backends"
)

func NewDynamicDirectoryResource(dir *filesystem.Directory) (*MetricResource, error) {
	err := dir.Ensure(backends.DynamicDirectoryMode)
	if err != nil {
		return nil, fmt.Errorf("unable to ensure directory: %s: %v", dir.String(), err)
	}
	resource := &MetricResource{
		backend: backends.NewDynamicDirectoryBackend(dir),
	}
	return resource, nil
}

const (
	ZFSMetricsDataDirectory = "/data/metrics/"
)

func NewDataDynamicDirectoryResource() (*MetricResource, error) {
	dir := filesystem.NewDirectory(ZFSMetricsDataDirectory)
	resource, err := NewDynamicDirectoryResource(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to create metrics resource: %v", err)
	}
	return resource, nil
}
