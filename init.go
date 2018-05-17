package logconf

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/wangxianzhuo/filehook"

	log "github.com/sirupsen/logrus"
)

var (
	FilePath        = flag.String("log-path", "./logs/", "logrus will store to this path")
	LineBreak       = flag.String("line-break", "\n", "line break")
	SegmentInterval = flag.Int64("segment-interval", 86400, "file segment interval")
	FileNamePattern = flag.String("file-name-pattern", "%YY-%MM-%DD_%HH-%mm-%SS.log", "log file name pattern")
	Level           = flag.String("level", "info", "log level, can be: info, warn, debug, error, fatal, panic")
)

// Configure logrus basic represent, log level and add a file hook
func Configure() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)

	LogLevel(*Level)

	h, err := filehook.New(&filehook.Option{
		Path:            *FilePath,
		SegmentInterval: *SegmentInterval,
		NamePattern:     *FileNamePattern,
		LineBreak:       *LineBreak,
	})

	if err != nil {
		panic(err)
	}

	log.AddHook(h)
}

// CustomFormatConifgure ...
func CustomFormatConifgure(formater log.Formatter) {
	log.SetFormatter(formater)
}

// LogLevel set log level
func LogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

// PrintConfigs print logrus config to io.writer
func PrintConfigs(w io.Writer) {
	configs := make(map[string]interface{})
	configs["log-path"] = *FilePath
	configs["segment-interval"] = *SegmentInterval
	configs["file-name-pattern"] = *FileNamePattern
	configs["level"] = *Level
	c, _ := json.Marshal(configs)
	fmt.Fprintf(w, "%v", string(c))
}

// Configurations print logrus configs to stdout
func Configurations() {
	PrintConfigs(os.Stdout)
	fmt.Fprint(os.Stdout, "\n")
}
