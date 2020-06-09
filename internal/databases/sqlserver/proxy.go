package sqlserver

import (
	"context"
	"github.com/wal-g/storages/storage"
	"github.com/wal-g/tracelog"
	"github.com/wal-g/wal-g/internal/databases/sqlserver/blob"
	"github.com/wal-g/wal-g/utility"
	"os"
	"syscall"
)

func RunProxy(folder storage.Folder) {
	ctx, cancel := context.WithCancel(context.Background())
	signalHandler := utility.NewSignalHandler(ctx, cancel, []os.Signal{syscall.SIGINT, syscall.SIGTERM})
	defer func() { _ = signalHandler.Close() }()
	bs, err := blob.NewServer(folder)
	tracelog.ErrorLogger.FatalfOnError("blob proxy error: %v", err)
	err = bs.Run(ctx)
	tracelog.ErrorLogger.FatalfOnError("blob proxy error: %v", err)
}