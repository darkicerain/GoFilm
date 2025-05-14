package config

import (
	"bytes"
	_ "github.com/joho/godotenv/autoload"
	"go.beyondstorage.io/v5/types"
	"log"
	"mime/multipart"
	"os"
	"reflect"
	"server/plugin/storage"
)

var (
	Store = GetStore()
)

func GetStore() types.Storager {
	var store types.Storager = nil

	if os.Getenv("STORAGE_TYPE") == "minio" {
		store, err = storage.NewMinio()
		if err != nil {
			log.Print("minio err:", err)
		}
	}
	if os.Getenv("STORAGE_TYPE") == "fs" || os.Getenv("STORAGE_TYPE") == "" || reflect.ValueOf(store).IsZero() {
		store, err = storage.NewFs()
		if err != nil {
			log.Print("fs err:", err)
		}
	}
	return store
}

func SaveFile(file *multipart.FileHeader, fileName string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	var b = make([]byte, file.Size)
	_, err = src.Read(b)
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var pair types.Pair = types.Pair{Key: "content_type", Value: file.Header.Get("content-type")}
	_, err = Store.Write(fileName, r, file.Size, pair)
	if err != nil {
		return err
	}
	return nil
}

func GetUrl(fileName string) string {
	var url string
	if os.Getenv("STORAGE_TYPE") == "fs" || os.Getenv("STORAGE_TYPE") == "" {
		url = "/api/static" + fileName
	} else {
		url, err = Store.(types.Reacher).Reach(fileName)
		if err != nil {
			log.Print("get url err:", err)
			return ""
		}
	}

	return url
}
