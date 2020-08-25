package initialization

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type configuration struct {
	Title string
	FilePath filePath
	Category category
}

type filePath struct {
	ArticlePath string
}

type category struct {
	Label []interface{}
	Default string
	Separator string
}

var Conf configuration

func init() {
	confFile := "conf.toml"
	if _, err := toml.DecodeFile(confFile, &Conf); err != nil {
		log.Panic(err)
	}
	log.Println("文件夹路径初始化：", Conf.FilePath.ArticlePath)
	pwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	path := pwd + string(os.PathSeparator) + Conf.FilePath.ArticlePath
	state, err := os.Stat(path)
	if err != nil {
		// 路径存在
		if os.IsExist(err) {
			// 不是文件夹，创建
			if !state.IsDir() {
				err := os.MkdirAll(path, os.ModePerm)
				if err != nil {
					log.Panic("无法创建存放文件的文件夹，请手动创建", err)
				} else {
					log.Println("创建文件夹：", path)
				}
			}
		} else if os.IsNotExist(err) {
			// 不存在
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				log.Panic("无法创建存放文件的文件夹，请手动创建", err)
			} else {
				log.Println("创建文件夹：", path)
			}
		} else {
			// 其他错误
			log.Panic(err)
		}
	}

}