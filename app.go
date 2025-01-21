package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/net"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.getNetworkUsage()
}

// send network data to frontend
func (a *App) getNetworkUsage() {
	for {
		status, err := net.IOCounters(false)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(status) > 0 {
			data := map[string]interface{}{
				"bytesSent":   status[0].BytesSent,
				"bytesRecv":   status[0].BytesRecv,
				"packetsSent": status[0].PacketsSent,
				"packetsRecv": status[0].PacketsRecv,
			}

			runtime.EventsEmit(a.ctx, "networkData", data)
		}
		time.Sleep(1 * time.Second)
	}
}
