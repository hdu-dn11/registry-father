package services

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"registry-father/model"
	"strconv"
)

func GetASInfoList() ([]*model.ASInfo, error) {
	asPath := "./data"
	list := make([]*model.ASInfo, 0)
	files, err := os.ReadDir(asPath)
	if err != nil {
		if err != os.ErrNotExist {
			_ = os.MkdirAll(asPath, 0644)
			return list, nil
		}
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fName := file.Name()
		if len(fName) <= len("AS.yaml") || fName[:2] != "AS" || path.Ext(fName) != ".yaml" {
			continue // todo: 考虑是否报告异常配置
		}
		asn, err := strconv.ParseUint(fName[2:len(fName)-len(".yaml")], 10, 32)
		if err != nil {
			continue
		}
		fByte, err := os.ReadFile(path.Join(asPath, file.Name()))
		if err != nil {
			continue
		}
		asInfo := &model.ASInfo{
			Path: path.Join(asPath, file.Name()),
		}
		err = yaml.Unmarshal(fByte, asInfo)
		if err != nil || uint32(asn) != asInfo.ASN {
			continue
		}
		fInfo, _ := file.Info()
		asInfo.UpdatedAt = fInfo.ModTime()
		list = append(list, asInfo)
	}
	return list, nil
}

func SaveASInfo(info *model.ASInfo) error {
	if fmt.Sprintf("data/AS%d.yaml", info.ASN) != info.Path {
		_ = os.Remove(info.Path)
		return os.WriteFile(fmt.Sprintf("data/AS%d.yaml", info.ASN), info.YAML(), 0644)
	}
	return os.WriteFile(fmt.Sprintf("data/AS%d.yaml", info.ASN), info.YAML(), 0644)
}

func DeleteASInfo(info *model.ASInfo) error {
	return os.Remove(info.Path)
}
