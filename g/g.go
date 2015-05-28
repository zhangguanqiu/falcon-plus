package g

import (
	"log"
	"runtime"
)

// changelog:
// 0.0.1: init project
// 0.0.3: add readme, add gitversion, modify proc, add config reload
// 0.0.4: make collector configurable, add monitor cron, adjust index db
// Changes: send turning-ok only after alarm happens
// TODO set http-get conn&rw timeout
const (
	VERSION = "0.0.4"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
