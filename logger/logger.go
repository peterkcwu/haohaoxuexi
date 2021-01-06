package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var Web *logrus.Logger
var Report *logrus.Logger

func InitWebLogger(conf *LogConf) error {
	logger, err := InitLogger(conf)
	if err != nil {
		return err
	}
	Web = logger
	return nil
}

//初始化
func InitReportLogger(conf *LogConf) error {
	logger, err := InitLogger(conf)
	if err != nil {
		return err
	}
	Report = logger
	return nil
}

func NewLoggerInstance() *logrus.Logger {
	l := logrus.New()
	l.SetLevel(logrus.ErrorLevel)
	return l
}

type Logger struct {
	mu     sync.RWMutex
	Logger *logrus.Logger
}

func NewLogger() *Logger {
	l := &Logger{}
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	l.SetLogger(log)
	return l
}

func (c *Logger) SetLogger(logger *logrus.Logger) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if logger == nil {
		// 防止空指针
		if c.Logger == nil {
			c.Logger = logrus.New()
		}
		return
	}
	c.Logger = logger
}

func (c *Logger) GetLogger() *logrus.Logger {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.Logger == nil {
		// 防止空指针
		c.Logger = logrus.New()
	}
	return c.Logger
}

// default
func defaultValue(logConf *LogConf) *LogConf {
	if logConf.LogPath == "" {
		logConf.LogPath = fmt.Sprintf("../logger/%s_golib.logger", os.Args[0])
	}
	if logConf.LogLevel == "" {
		logConf.LogLevel = "info"
	}
	if logConf.LogReserveDay == 0 {
		logConf.LogReserveDay = 7
	}
	return logConf
}

// 初始化
func InitLogger(logConf *LogConf) (*logrus.Logger, error) {
	logConf = defaultValue(logConf)
	level, err := logrus.ParseLevel(logConf.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("LogLevel err:%w\n", err)
	}

	var logger = logrus.New()
	if !filepath.IsAbs(logConf.LogPath) {
		logConf.LogPath = filepath.Join(filepath.Dir(os.Args[0]), logConf.LogPath)
	}
	writer, err := rotatelogs.New(
		logConf.LogPath+".%Y%m%d",
		rotatelogs.WithLinkName(logConf.LogPath),
		rotatelogs.WithMaxAge(time.Duration(logConf.LogReserveDay)*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		return nil, err
	}
	logger.AddHook(lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
		logrus.TraceLevel: writer,
	}, &logrus.TextFormatter{}))
	logger.SetOutput(ioutil.Discard)
	logger.SetLevel(level)
	logger.SetReportCaller(logConf.ReportCaller)
	return logger, nil
}

// 初始化
type LogConf struct {
	LogLevel      string
	LogPath       string
	LogReserveDay int
	ReportCaller  bool
}
