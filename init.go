package logconf

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tracer0tong/kafkalogrus"
	"github.com/wangxianzhuo/filehook"

	log "github.com/sirupsen/logrus"
)

var (
	FilePath        = flag.String("log-path", "./logs/", "logrus will store to this path")
	LineBreak       = flag.String("line-break", "\n", "line break")
	SegmentInterval = flag.Int64("segment-interval", 86400, "file segment interval")
	FileNamePattern = flag.String("file-name-pattern", "%YY-%MM-%DD_%HH-%mm-%SS.log", "log file name pattern")
	Level           = flag.String("level", "info", "log level, can be: info, warn, debug, error, fatal, panic")
	Debug           = flag.Bool("debug", false, "debug mode")
	KafkaTopic      = flag.String("log-kafka-topic", "log_msg", "kafka topic for log")
	KafkaBrokers    = flag.String("log-kafka-brokers", "localhost:9092", "kafka brokers for log, it can be like '192.168.1.100:9092,192.168.1.101:9092'")
	ToFileSystem    = flag.Bool("to-file-sys", false, "switch of sending log to local file system")
	ToKafka         = flag.Bool("to-kafka", false, "switch of sending log to kafka")
)

func defalutFormatter() log.Formatter {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.999999999"
	customFormatter.FullTimestamp = true
	return customFormatter
}

func defalutJSONFormatter() log.Formatter {
	customFormatter := new(log.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.999999999"
	return customFormatter
}

// Configure logrus basic represent, log level and add a file hook
func Configure() {
	log.SetFormatter(defalutFormatter())

	LogLevel(*Level)
	if *Debug {
		log.SetLevel(log.DebugLevel)
	}

	// ConfigureLocalFileHook()

	// ConfigureKafkaHook()
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

// ConfigureKafkaHook config kafka hook for logrus
func ConfigureKafkaHook() {
	if !*ToKafka {
		return
	}
	if *KafkaTopic == "" {
		fmt.Println("Error: --log-kafka-topic can't be empty")
		flag.PrintDefaults()
		os.Exit(2)
	}
	if *KafkaTopic == "" {
		fmt.Println("Error: --log-kafka-brokers can't be empty")
		flag.PrintDefaults()
		os.Exit(2)
	}
	l := strings.Split(*KafkaBrokers, ",")
	if len(l) < 1 {
		fmt.Println("Error: --log-kafka-brokers can't be empty")
		flag.PrintDefaults()
		os.Exit(2)
	}

	var brokers []string

	for _, i := range l {
		brokers = append(brokers, strings.TrimSpace(i))
	}

	hook, err := kafkalogrus.NewKafkaLogrusHook(
		"kh",
		log.AllLevels,
		defalutJSONFormatter(),
		brokers,
		*KafkaTopic,
		true, nil)

	if err != nil {
		panic(err)
	}

	log.AddHook(hook)
}

// ConfigureLocalFileHook config local file system hook for logrus
func ConfigureLocalFileHook() {
	if !*ToFileSystem {
		return
	}
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
