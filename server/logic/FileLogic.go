package logic

import (
	"os"
	"path/filepath"
	"server/config"
	"server/model/system"
	"strings"
)

type FileLogic struct {
}

var FileL FileLogic

func (fl *FileLogic) SingleFileUpload(fileName string, uid int) string {
	// 生成图片信息
	var link = config.GetUrl(fileName)
	var f = system.FileInfo{Link: link, FilePath: fileName, Uid: uid, Type: 0}
	f.Fid = strings.TrimSuffix(filepath.Base(fileName), filepath.Ext(fileName))
	f.FileType = strings.TrimPrefix(filepath.Ext(fileName), ".")
	f.FileStorage = "fs"
	if os.Getenv("STORAGE_TYPE") != "" {
		f.FileStorage = os.Getenv("STORAGE_TYPE")
	}
	// 记录图片信息到系统表中
	system.SaveGallery(f)
	return f.Link
}

// GetPhotoPage 获取系统内的图片分页信息
func (fl *FileLogic) GetPhotoPage(page *system.Page) []system.FileInfo {
	// 设置必要参数
	var tl = []string{"jpeg", "jpg", "png", "webp"}
	return system.GetFileInfoPage(tl, page)
}

// RemoveFileById 删除文件信息
func (fl *FileLogic) RemoveFileById(id uint) error {
	// 首先获取对应图片信息
	f := system.GetFileInfoById(id)
	// 通过f删除本地图片
	err := f.DelFile()
	if err != nil {
		return err
	}
	// 删除图片的关联信息
	system.DelFileInfo(id)
	return err
}
