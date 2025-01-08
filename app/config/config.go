package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
)

var command = &cobra.Command{
	Use:   "c",
	Long:  `zng: zng config`,
	Short: "配置信息管理",
}

type Config struct {
	Version string `yaml:"version"`
}

func Init() *cobra.Command {
	command.AddCommand(
		List(),
		Set(),
		Remove(),
		AllList(),
	)
	return command
}

// Remove
func Remove() *cobra.Command {
	return &cobra.Command{
		Use:   "rm",
		Long:  `zng: zng config remove`,
		Short: "remove 删除配置",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("参数错误")
				return
			}
			viper.Set(args[0], nil)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println("写入配置文件失败:", err)
			}
		},
	}
}

func AllList() *cobra.Command {
	return &cobra.Command{
		Use:   "w",
		Long:  `zng: zng list all`,
		Short: "all info 获取所有配置信息",
		Run: func(cmd *cobra.Command, args []string) {
			keys := viper.AllKeys()
			for _, key := range keys {
				fmt.Println(fmt.Sprintf("%s=%s", key, viper.Get(key)))
			}
		},
	}
}

func Set() *cobra.Command {
	return &cobra.Command{
		Use:   "s",
		Long:  `zng: zng set`,
		Short: "set info 设置配置信息",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("参数错误")
				return
			}
			viper.Set(args[0], args[1])
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println("写入配置文件失败:", err)
				return
			}
		},
	}
}
func List() *cobra.Command {
	return &cobra.Command{
		Use:   "l",
		Long:  `zng: zng list`,
		Short: "list info 获取配置信息",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("参数错误")
				return
			}
			content := viper.Get(args[0])
			fmt.Println(content)
		},
	}
}

func Run() {
	dir, err2 := getConfigDir()
	if err2 != nil {
		fmt.Println("获取配置目录失败:", err2)
		return
	}
	err2 = CreateDefaultConfig(dir)
	if err2 != nil {
		fmt.Println("创建配置文件失败:", err2)
		return
	}
	viper.AutomaticEnv()
	viper.AddConfigPath(dir)      // 配置文件夹路径
	viper.SetConfigName("config") // 配置文件名，假设为 config.yaml
	//viper.SetConfigType("yaml")   // 配置文件类型是 YAML
	viper.SetConfigType("toml") // 配置文件类型是 YAML
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("配置文件读取错误，创建配置文件:", err)
		return
	}
}
func getConfigDir() (string, error) {
	var configDir string

	switch runtime.GOOS {
	case "windows":
		// Windows 上，放在 C:\Users\Administrator\.zng 目录下
		configDir = filepath.Join("C:", "Users", "Administrator", ".zng")
	case "linux", "darwin":
		// Linux 和 macOS 上，放在用户的主目录下的 .zng 文件夹
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configDir = filepath.Join(homeDir, ".zng")
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	// 检查目录是否存在，如果不存在则创建
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return configDir, nil
}
func CreateDefaultConfig(configDir string) (err error) {
	// 创建并写入默认配置文件
	configFilePath := configDir + "/config.toml"
	//判断文件是否存在
	_, err = os.Stat(configFilePath)
	if err == nil {
		//fmt.Println("config.yaml 配置文件已存在", configFilePath)
		fmt.Println(configFilePath)
		return
	}
	var f *os.File
	f, err = os.Create(configFilePath)
	if err != nil {
		return
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			return
		}
	}(f)
	return
}
