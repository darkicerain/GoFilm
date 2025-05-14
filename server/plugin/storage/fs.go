package storage

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	fs "go.beyondstorage.io/services/fs/v4"
	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
	"os"
)

func NewFs() (types.Storager, error) {
	return fs.NewStorager(
		// WorkDir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_FS_WORKDIR")),
	)
}

func NewFsFromString() (types.Storager, error) {
	connStr := fmt.Sprintf(
		"fs://%s",
		os.Getenv("STORAGE_FS_WORKDIR"),
	)
	return services.NewStoragerFromString(connStr)
}
