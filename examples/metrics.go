package main

import (
	"fmt"
	"os"

	"github.com/kris-nova/novaarchive/filesystem"

	"github.com/kris-nova/logger"
	"github.com/kris-nova/novaarchive/metrics"
)

func main() {
	dir := filesystem.NewDirectory("/tmp/novametrics")
	resource, err := metrics.NewDynamicDirectoryResource(dir)
	if err != nil {
		if err != nil {
			logger.Critical(err.Error())
			os.Exit(1)
		}
	}
	err = resource.Set("key", "value")
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}
	err = resource.Set("key", "newValue")
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}
	record, err := resource.Get("key")
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}

	fmt.Println(record.Value)
}
