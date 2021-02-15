package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"engine/global"

	"engine/log"
)

func waitSingals(ctx context.Context) error {
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, os.Kill)

	select {
	case <-ctx.Done():
	case sig := <-osSignalChan:
		fmt.Println(sig)
	}
	return nil
}

func main() {
	log.Log.Info("启动程序")

	defer global.Global.Cancel()

	global.Global.StartAll()

	global.Global.WaitAll()

	waitSingals(global.Global.GlobalCtx)

	log.Log.Info("程序退出")
}
