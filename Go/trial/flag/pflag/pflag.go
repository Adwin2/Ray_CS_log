package main

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/pflag"
)

type Config struct {
	Dir        string
	DBFilename string
}

func ParseFlags() *Config {
	var cfg Config

	// 1. 定义命令行参数并绑定到结构体
	pflag.StringVar(&cfg.Dir, "dir", "/var/lib/app", "持久化数据存储目录")
	pflag.StringVar(&cfg.DBFilename, "dbfilename", "dump.rdb", "数据库文件名")

	// 2. 支持 POSIX 风格短参数（可选）
	pflag.StringVarP(&cfg.Dir, "directory", "d", cfg.Dir, "短参数别名")

	// 3. 解析参数
	pflag.Parse()

	// 4. 参数后置校验（真实项目常用模式）
	if !filepath.IsAbs(cfg.Dir) {
		cfg.Dir = filepath.Join("/", cfg.Dir)
	}

	return &cfg
}

func main() {
	var cfg Config

	fmt.Println(cfg.Dir, cfg.DBFilename)

	cfg = *ParseFlags()

	fmt.Println(cfg.Dir, cfg.DBFilename)
}
