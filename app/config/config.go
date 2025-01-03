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
	Use:  "config",
	Long: `zng: zng build`,
}

type Config struct {
	Version string `yaml:"version"`
}

func Init() *cobra.Command {
	command.AddCommand(
		List(),
		Set(),
	)
	return command
}
func Set() *cobra.Command {
	return &cobra.Command{
		Use:  "set",
		Long: `zng: zng set`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) <= 2 {
				fmt.Println("参数错误")
				return
			}
			viper.Set(args[1], args[2])
		},
	}
}
func List() *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Long: `zng: zng list`,
		Run: func(cmd *cobra.Command, args []string) {
			var version = viper.Get("version")
			fmt.Println(version)
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
	viper.AddConfigPath(dir)      // 配置文件夹路径
	viper.SetConfigName("config") // 配置文件名，假设为 config.yaml
	viper.SetConfigType("yaml")   // 配置文件类型是 YAML
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
		err := os.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return configDir, nil
}
func CreateDefaultConfig(configDir string) (err error) {
	// 创建并写入默认配置文件
	configFilePath := configDir + "/config.yaml"
	//判断文件是否存在
	_, err = os.Stat(configFilePath)
	if err == nil {
		fmt.Println("config.yaml 配置文件已存在")
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
	// 写入默认配置内容
	defaultConfig := `version: 1.0.2`
	_, err = f.WriteString(defaultConfig)
	return
}
