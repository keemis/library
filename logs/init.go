package logs

import (
	"encoding/json"

	"github.com/astaxie/beego"
	beeLogs "github.com/astaxie/beego/logs"
)

var (
	// 全局日志配置
	logger = beeLogs.GetBeeLogger()
)

// Conf 用法参考 https://beego.me/docs/module/logs.md
type Conf struct {
	FileName string `json:"filename"`
	MaxLines int64  `json:"maxlines"`
	MaxSize  int64  `json:"maxsize"`
	Daily    bool   `json:"daily"`
	Maxdays  int64  `json:"maxdays"`
	Rotate   bool   `json:"rotate"`
	Level    int    `json:"level"`
}

// defaultConf 默认配置
func defaultConf() *Conf {
	return &Conf{
		FileName: beego.AppConfig.DefaultString("LogFile", "logs/app.log"),
		MaxLines: beego.AppConfig.DefaultInt64("LogMaxLines", 0),
		MaxSize:  beego.AppConfig.DefaultInt64("LogMaxSize", 0),
		Daily:    beego.AppConfig.DefaultBool("LogDaily", true), // 按天切割
		Maxdays:  beego.AppConfig.DefaultInt64("LogMaxDays", 7), // 保存7天
		Rotate:   beego.AppConfig.DefaultBool("LogRotate", true),
		Level:    beego.AppConfig.DefaultInt("LogLevel", beego.LevelDebug),
	}
}

// Init 初始化
func Init(conf *Conf) {
	if conf == nil {
		conf = defaultConf()
	}
	byt, err := json.Marshal(conf)
	if err != nil {
		panic(err)
	}
	_ = logger.SetLogger("file", string(byt))
	logger.EnableFuncCallDepth(false)
	// dev模式 输出到控制台
	mode := beego.AppConfig.String("RunMode")
	if mode != "dev" {
		_ = logger.DelLogger("console")
	}
}
