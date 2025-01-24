package template

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zngue/zng_tool/app/proto/proto"
	"github.com/zngue/zng_tool/app/proto/proto/kratos/biz"
	"github.com/zngue/zng_tool/app/proto/proto/kratos/data"
	"github.com/zngue/zng_tool/app/proto/proto/kratos/service"
	"github.com/zngue/zng_tool/app/proto/proto/types"
	"github.com/zngue/zng_tool/app/util"
	"os"
	"path/filepath"
	"strings"
)

var kratosTmp = &cobra.Command{
	Use:   "kt",
	Long:  `zng: zng proto`,
	Short: "template zng tmp kt",
	Run: func(cmd *cobra.Command, args []string) {
		inPath := "api/gin-pb/v1/gin-pb.proto"
		filepath.Join()
		var (
			out     string
			outPath []string
		)
		if len(args) >= 2 {
			out = args[1]
			split := strings.Split(out, "/")
			for _, s := range split {
				if s != "" {
					outPath = append(outPath, s)
				}
			}
		} else {
			outPath = []string{"demo"}
		}
		err := Mkdir(filepath.Join(outPath...))
		if err != nil {
			fmt.Println(err)
			return
		}
		//判断下面的 biz data service 文件夹是否存在不存在则创建
		var bizPathSlice = append(outPath, "biz")
		var bizPath = filepath.Join(bizPathSlice...)
		if !IsDir(bizPath) {
			err := Mkdir(bizPath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		var dataPathSlice = append(outPath, "data")
		var dataPath = filepath.Join(dataPathSlice...)
		if !IsDir(dataPath) {
			err := Mkdir(dataPath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		var servicePathSlice = append(outPath, "service")
		var servicePath = filepath.Join(servicePathSlice...)
		if !IsDir(servicePath) {
			err := Mkdir(servicePath)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		scs := proto.Inits(inPath)
		//根据服务生成文件
		for _, sc := range scs {
			DataFile(sc, dataPath)
			BizFile(sc, bizPath)
			ServiceFile(sc, servicePath)
		}
	},
}

// ServiceFile 生成service文件
func ServiceFile(sc *types.ServiceDesc, outPath string) {
	//将content 写入到文件
	tmp := service.NewDataTemplate(sc).Execute()
	var fileName = util.UpperLineToLower(sc.ServiceName)
	//生成data 数据
	var outFileName = fmt.Sprintf("%s/%s.go", outPath, fileName)
	err := util.WriteFile(outFileName, tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// BizFile 生成biz文件
func BizFile(sc *types.ServiceDesc, outPath string) {
	//将content 写入到文件
	tmp := biz.NewDataTemplate(sc).Execute()
	var fileName = util.UpperLineToLower(sc.ServiceName)
	//生成data 数据
	var outFileName = fmt.Sprintf("%s/%s.go", outPath, fileName)
	//判断文件是否存在
	if _, err := os.Stat(outFileName); err == nil {
		fmt.Println(outFileName, "文件已存在")
		//outFileName = fmt.Sprintf("%s/%s.d.go", outPath, fileName)
	} else {
		err = nil
	}
	err := util.WriteFile(outFileName, tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// DataFile 生成data文件
func DataFile(sc *types.ServiceDesc, outPath string) {
	//将content 写入到文件
	tmp := data.NewDataTemplate(sc).Execute()
	var fileName = util.UpperLineToLower(sc.ServiceName)
	//生成data 数据
	var outFileName = fmt.Sprintf("%s/%s.go", outPath, fileName)
	//判断文件是否存在
	if _, err := os.Stat(outFileName); err == nil {
		fmt.Println(outFileName, "文件已存在")
		//outFileName = fmt.Sprintf("%s/%s.d.go", outPath, fileName)
	} else {
		err = nil
	}

	err := util.WriteFile(outFileName, tmp)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func KratosInit() *cobra.Command {
	kratosTmp.AddCommand()
	return kratosTmp
}
func KratosBiz() *cobra.Command {
	return &cobra.Command{
		Use:   "biz",
		Long:  `zng: zng proto`,
		Short: "template zng tmp k ",
	}
}

// IsDir 判断文件夹是否存在
func IsDir(path string) bool {
	if f, err := os.Stat(path); err == nil {
		return true
	} else {
		if f == nil {
			return false
		}
		return f.IsDir()
	}
}

// Mkdir 文件夹不存在就创建
func Mkdir(path string) (err error) {
	if !IsDir(path) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return
}
