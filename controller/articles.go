package controller

import (
	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
	"goBlog/initialization"
	"goBlog/service"
	"goBlog/util"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

func Articles(c *gin.Context)  {
	log.Println("method:", c.Request.Method, "url:", c.Request.URL.Path)

	filename := c.Query("name")
	category := c.DefaultQuery("category", initialization.Conf.Category.Default)
	fTotal := filename + initialization.Conf.Category.Separator + category + ".md"
	fpath := ""
	if util.Exists(initialization.Conf.FilePath.ArticlePath + string(os.PathSeparator) + fTotal) {
		fpath = initialization.Conf.FilePath.ArticlePath + string(os.PathSeparator) + fTotal
	} else if util.Exists(initialization.Conf.FilePath.ArticlePath + string(os.PathSeparator) + filename + ".md") {
		fpath = initialization.Conf.FilePath.ArticlePath + string(os.PathSeparator) + filename + ".md"
	} else {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"title": "发生错误",
			"error": "未找到该文章",
		})
	}
	log.Println("path:", fpath)
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Panicln(err)
	}

	html := lute.New().Markdown("", b)
	c.HTML(http.StatusOK, "markdown.tmpl", gin.H{
		"title": filename,
		"info": filename,
		"body": template.HTML(html),
	})
}

func Total(c *gin.Context) {
	log.Println("method:", c.Request.Method, "url:", c.Request.URL.Path)
	list := new(service.MdList)
	searchCategory := c.Query("category")
	searchArticle := c.Query("q")
	categoryMap := make(map[string]int)
	if searchCategory != "" && searchArticle != "" {
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"title": "发生错误",
			"error": "仅支持单个参数查询",
		})
	} else if searchCategory == "" && searchArticle == ""{
		categoryMap = list.GetArticleList()
	} else if searchCategory != "" {
		categoryMap = list.GetArticleListByCategory(searchCategory)
	} else {
		categoryMap = list.GetArticleListByQuery(searchArticle)
	}
	sort.Sort(sort.Reverse(list))
	c.HTML(http.StatusOK, "list.tmpl", gin.H{
		"title": "Go Blog",
		"content": list,
		"categoryMap": categoryMap,
	})
}