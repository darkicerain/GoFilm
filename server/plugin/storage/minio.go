package storage

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	minio "go.beyondstorage.io/services/minio"
	"go.beyondstorage.io/v5/pairs"
	"go.beyondstorage.io/v5/services"
	"go.beyondstorage.io/v5/types"
	"os"
)

func NewMinio() (types.Storager, error) {
	var store, err = minio.NewStorager(
		// credential: https://beyondstorage.io/docs/go-storage/pairs/credential
		//
		// Credential could be fetched from service's console.
		//
		// Example Value: hmac:access_key_id:secret_access_key
		pairs.WithCredential(os.Getenv("STORAGE_MINIO_CREDENTIAL")),
		// endpoint: https://beyondstorage.io/docs/go-storage/pairs/endpoint
		//
		// Example Value: https:host:port
		pairs.WithEndpoint(os.Getenv("STORAGE_MINIO_ENDPOINT")),
		// name: https://beyondstorage.io/docs/go-storage/pairs/name
		//
		// name is the bucket name.
		pairs.WithName(os.Getenv("STORAGE_MINIO_NAME")),
		// work_dir: https://beyondstorage.io/docs/go-storage/pairs/work_dir
		//
		// Relative operations will be based on this WorkDir.
		pairs.WithWorkDir(os.Getenv("STORAGE_MINIO_WORKDIR")),
	)
	return store, err
}

func NewMinioFromString() (types.Storager, error) {
	str := fmt.Sprintf(
		"minio://%s%s?credential=%s&endpoint=%s",
		os.Getenv("STORAGE_MINIO_NAME"),
		os.Getenv("STORAGE_MINIO_WORKDIR"),
		os.Getenv("STORAGE_MINIO_CREDENTIAL"),
		os.Getenv("STORAGE_MINIO_ENDPOINT"),
	)
	return services.NewStoragerFromString(str)
}
