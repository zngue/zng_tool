package main

import (
	_ "embed"
	"fmt"
	"github.com/zngue/zng_tool/app/util"
	"os"
	"regexp"
	"strings"
)

//go:embed wire_template.tpl
var wireTemplate string

func dbReplace(dir, fileName, content string) {
	util.IsDir(dir)
	_ = util.WriteFile(fmt.Sprintf("%s/%s.go", dir, fileName), content)
}

func ReplaceWire(dir, fileName, serverName string, pkg string) {
	fileName = fmt.Sprintf("%s/%s.go", dir, fileName)
	//文件不存在则创建文件
	if !util.FileExists(fileName) {
		tmp := strings.ReplaceAll(wireTemplate, "{{PKG}}", pkg)
		tmp = strings.ReplaceAll(tmp, "{{CONTENT}}", serverName)
		err := util.WriteFile(fileName, tmp)
		if err != nil {
			return
		}
		return
	}
	readFile, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	re := regexp.MustCompile(`wire\.NewSet\(([\s\S]*?)\)`)
	matches := re.FindStringSubmatch(string(readFile))
	if len(matches) > 1 {
		// matches[0] 是整个匹配的字符串，matches[1] 是括号内的内容
		//将字符使用逗号分隔 并且去掉空格，和换行
		var params []string
		for _, param := range strings.Split(matches[1], ",") {
			param = strings.TrimSpace(param)
			if param != "" && !util.InArray(param, params) {
				params = append(params, param)
			}
		}
		////将新的加入
		if !util.InArray(serverName, params) {
			params = append(params, serverName)
		}
		////将新的替换
		newContent := "\n\t" + strings.Join(params, ",\n\t") + ",\n"
		newContent = strings.Replace(string(readFile), matches[1], newContent, 1)
		err = util.WriteFile(fileName, newContent)
		if err != nil {
			return
		}
	} else {
		tmp := strings.ReplaceAll(wireTemplate, "{{PKG}}", pkg)
		tmp = strings.ReplaceAll(tmp, "{{CONTENT}}", serverName)
		err = util.WriteFile(fileName, tmp)
		if err != nil {
			return
		}
	}
}

func WireFile(dir, fileName, content string) (err error) {
	util.IsDir(dir)
	return util.WriteFile(fmt.Sprintf("%s/%s.go", dir, fileName), content)
}
