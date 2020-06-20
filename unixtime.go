package unixtime

import "peterdekok.nl/gotools/logger"

var log logger.Logger

func init() {
    log = logger.New("unixtime")
}
