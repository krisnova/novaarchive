package main

import (
	lol "github.com/kris-nova/lolgopher"
	"github.com/kris-nova/novaarchive/logger"
)

func main() {
	//
	logger.Writer = lol.NewLolWriter()          // Sometimes this will work better
	logger.Writer = lol.NewTruecolorLolWriter() // Comment one of these out
	//

	logger.BitwiseLevel = logger.LogEverything
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Always("Always logging")
	logger.Debug("Debug logging")
	logger.Critical("Stderr logging")
}
