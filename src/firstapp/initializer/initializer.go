package initializer

import (
	. "firstapp/common/logger"
	"firstapp/dassignal"
	"firstapp/runconfig"
	"fmt"
)

func Initlize() error {
	var err error

	if err = runconfig.Config.Parse(); err != nil {
		LogTemp("runconfig.Config.Parse error: %s", err)
		return err
	}

	if err = initLog(); err != nil {
		LogTemp("initLog error: %s", err)
		return err
	}
	dassignal.Trap(cleanup)
	return nil
}



func initLog() error {
	logConfig := fmt.Sprintf(`{
		"filename": %q,
		"perm":     "0775"
	}`, runconfig.Config.Log.FilePath)

	return Log.Initialize(runconfig.Config.Log.Level, logConfig)
}


// 清理工作
func cleanup() {
	Log.Info("==========cleanup")
}
