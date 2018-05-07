package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/wangxianzhuo/logrus-conf"
)

func main() {
	logconf.Configure()
	log.Infof("log file path: %s, line break: %s, segment interval: %d, file name pattern: %s, level: %s", *logconf.FilePath, *logconf.LineBreak, *logconf.SegmentInterval, *logconf.FileNamePattern, *logconf.Level)
}
