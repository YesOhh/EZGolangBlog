package service

import (
	"goBlog/initialization"
	"goBlog/mylog"
	"io/ioutil"
	"strings"
)

type MdList struct {
	MdInfos []mdInfo
}

type mdInfo struct {
	Title string
	Category string
	ModTime string
}

var filesPath = initialization.Conf.FilePath.ArticlePath
var categories = initialization.Conf.Category.Label
var defaultCategory = initialization.Conf.Category.Default
var separator = initialization.Conf.Category.Separator

func (list *MdList) GetArticleList() map[string]int {
	if separator == "" {
		separator = "#"
	}
	dir, err := ioutil.ReadDir(filesPath)
	if err != nil {
		mylog.MyLogger.Panic("读取文件夹内文档失败")
	}

	categoryMap := make(map[string]int)
	for _, fileInfo := range dir {
		if fileInfo.IsDir() {
			// 目前不递归文件夹
			continue
		}
		ok := strings.HasSuffix(fileInfo.Name(), ".md")
		if ok {
			fName, fCategory := getNameCategoryOrDefault(fileInfo.Name())
			curInfo := mdInfo{Title: fName, Category: defaultCategory, ModTime: fileInfo.ModTime().Format("2006-01-02 15:04:05")}
			if isInCategories(categories, fCategory) {
				curInfo.Category = fCategory
				categoryMap[fCategory] += 1
			} else {
				categoryMap[defaultCategory] += 1
			}
			list.MdInfos = append(list.MdInfos, curInfo)
		}
	}
	return categoryMap
}

func (list *MdList) GetArticleListByCategory(category string) map[string]int {
	if separator == "" {
		separator = "#"
	}
	dir, err := ioutil.ReadDir(filesPath)
	if err != nil {
		mylog.MyLogger.Panic("读取文件夹内文档失败")
	}

	categoryMap := make(map[string]int)
	for _, fileInfo := range dir {
		if fileInfo.IsDir() {
			// 目前不递归文件夹
			continue
		}
		ok := strings.HasSuffix(fileInfo.Name(), ".md")
		if ok {
			fName, fCategory := getNameCategoryOrDefault(fileInfo.Name())
			if isInCategories(categories, fCategory) {
				categoryMap[fCategory] += 1
			} else {
				categoryMap[defaultCategory] += 1
			}
			if fCategory == category {
				list.MdInfos = append(list.MdInfos, mdInfo{Title: fName, Category: fCategory, ModTime: fileInfo.ModTime().Format("2006-01-02 15:04:05")})
			}
		}
	}
	return categoryMap
}

func (list *MdList) GetArticleListByQuery(q string) map[string]int {
	if separator == "" {
		separator = "#"
	}
	dir, err := ioutil.ReadDir(filesPath)
	if err != nil {
		mylog.MyLogger.Panic("读取文件夹内文档失败")
	}

	categoryMap := make(map[string]int)
	for _, fileInfo := range dir {
		if fileInfo.IsDir() {
			// 目前不递归文件夹
			continue
		}
		ok := strings.HasSuffix(fileInfo.Name(), ".md")
		if ok {
			fName, fCategory := getNameCategoryOrDefault(fileInfo.Name())
			if isInCategories(categories, fCategory) {
				categoryMap[fCategory] += 1
			} else {
				categoryMap[defaultCategory] += 1
			}
			if strings.Contains(fName, q) {
				list.MdInfos = append(list.MdInfos, mdInfo{Title: fName, Category: fCategory, ModTime: fileInfo.ModTime().Format("2006-01-02 15:04:05")})
			}
		}
	}
	return categoryMap
}

func getNameCategoryOrDefault(filename string) (string, string) {
	fInfo := strings.Split(strings.TrimRight(filename,".md"), separator)
	n := len(fInfo)
	// 只要有指定的文件名和分类间的分隔符，就把最后的字段当成分类，所以不要在文件名里乱加分隔符
	if n > 1 {
		return strings.Join(fInfo[:n-1], separator), fInfo[n-1]
	}
	return fInfo[0], defaultCategory
}

func isInCategories(categories []interface{}, fCategory string) bool {
	for _, category := range categories {
		if fCategory == category.(string) {
			return true
		}
	}
	return false
}

// 实现 MdList 的 sort 接口
func (list *MdList) Len() int {
	return len(list.MdInfos)
}

func (list *MdList) Less(i, j int) bool {
	return list.MdInfos[i].ModTime < list.MdInfos[j].ModTime
}

func (list *MdList) Swap(i, j int) {
	list.MdInfos[i], list.MdInfos[j] = list.MdInfos[j], list.MdInfos[i]
}