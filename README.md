# logrus conf

## Dependency

```shell
dep ensure
```

## Usage

```go
package main

import (
    log "github.com/sirupsen/logrus"
    "github.com/wangxianzhuo/logrus-conf"
)

func main() {
    // auto configure logrus and it's file system hook
    logconf.Configure()

    // print log configs to stdout
    logconf.Configurations()

    // start a log conf admin server
    s := logconf.Server(1)
    s.Start(":9090")

    log.Infof("log file path: %s, line break: %s, segment interval: %d, file name pattern: %s, level: %s", *logconf.FilePath, *logconf.LineBreak, *logconf.SegmentInterval, *logconf.FileNamePattern, *logconf.Level)
}
```