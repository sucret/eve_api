package logger

import (
	"eve/internal/config"
	"eve/internal/global"
	"eve/internal/tool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var (
	options []zap.Option
	conf    config.Logger
)

func New() (logger *zap.Logger, err error) {
	// 获取配置文件
	conf = global.Config.Logger

	// 创建日志存放目录
	rootDir, _ := tool.GetRootDir()
	logDir := rootDir + conf.FilePath
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	loggerConf := genConfig()

	loggerConf.EncoderConfig = genEncodeConfig()

	writer, err := genWriteSyncer()
	if err != nil {
		return nil, err
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(loggerConf.EncoderConfig),
		writer,
		loggerConf.Level,
	)

	// 添加调用位置
	options = append(options, zap.AddCaller())

	logger = zap.New(core, options...)

	return
}

// 生成WriteSyncer
func genWriteSyncer() (writeSyncer zapcore.WriteSyncer, err error) {
	// 创建日志存放目录
	rootDir, _ := tool.GetRootDir()
	logDir := rootDir + conf.FilePath
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	lumberJack := &lumberjack.Logger{
		Filename:   logDir + conf.FileName, // 文件名称
		MaxSize:    conf.MaxSize,           // 切割的文件大小 单位MB
		MaxBackups: conf.MaxBackups,        // 保留的最大文件数量
		MaxAge:     conf.MaxAge,            // 保留的文件最大天数
	}

	writeSyncer = zapcore.AddSync(lumberJack)
	return
}

// 生成配置
func genConfig() (config zap.Config) {
	config = zap.NewProductionConfig()

	config.EncoderConfig = genEncodeConfig()
	config.Level = zap.NewAtomicLevelAt(getLevel())

	return

}

// 生成编码配置
func genEncodeConfig() (c zapcore.EncoderConfig) {
	c = zap.NewProductionEncoderConfig()

	c.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}

	c.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(strings.ToUpper(l.String()))
	}

	c.TimeKey = "time"

	return
}

// 配置文件的level转换为zapcore的level
func getLevel() (level zapcore.Level) {
	switch conf.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	return
}

//
//func init2() {
//	rootDir, _ := tool.GetRootDir()
//
//	logDir := rootDir + "/tmp/log/"
//	err := os.MkdirAll(logDir, os.ModePerm)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	file, err := os.OpenFile(logDir+"eve.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	config := zap.NewProductionConfig()
//	config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
//
//	//config.OutputPaths = []string{"stdout", logDir + "logfile.log"}
//
//	logger, err := config.Build()
//	if err != nil {
//		panic(err)
//	}
//
//	//defer func(logger *zap.Logger) {
//	//	err := logger.Sync()
//	//	if err != nil {
//	//
//	//	}
//	//}(logger)
//
//	options := []zap.Option{zap.AddCaller()}
//
//	core := zapcore.NewCore(
//		zapcore.NewJSONEncoder(config.EncoderConfig),
//		zapcore.AddSync(file),
//		config.Level,
//	)
//
//	logger = zap.New(core, options...)
//
//	logger.Info("info", zap.String("name", "eve"),
//		zap.Any("user", struct {
//			Name string
//			Age  int
//		}{Name: "eve", Age: 18}))
//
//	logger.Debug("debug", zap.String("name", "eve"),
//		zap.Any("user", struct {
//			Name string
//			Age  int
//		}{Name: "eve", Age: 18}))
//
//	logger.Error("error", zap.String("name", "eve"),
//		zap.Any("user", struct {
//			Name string
//			Age  int
//		}{Name: "eve", Age: 18}))
//}

//func init() {
//	rootDir, _ := tool.GetRootDir()
//
//	logDir := rootDir + "/tmp/log/"
//	err := os.MkdirAll(logDir, os.ModePerm)
//	if err != nil {
//		fmt.Println("fweofw")
//		return
//	}
//
//	file, err := os.OpenFile(logDir+"eve.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	lumberJack := &lumberjack.Logger{
//		Filename:   logDir + "eve.log",
//		MaxSize:    10, // megabytes
//		MaxBackups: 3,
//		MaxAge:     28, //days
//	}

//	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
//
//	logger.Println("hello world")
//	logger.Print("hello world", Options{})
//
//}
