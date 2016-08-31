package pipeline

import (
	"github.com/r3boot/go-rtbh/config"
	"github.com/r3boot/go-rtbh/lib/blacklist"
	"github.com/r3boot/go-rtbh/lib/history"
	"github.com/r3boot/go-rtbh/lib/whitelist"
	"github.com/r3boot/rlib/logger"
)

const MYNAME string = "Pipeline"

var Config config.Config
var Log logger.Log

type Pipeline struct {
	ruleset   []*regexp.Regexp
	blacklist *blacklist.Blacklist
	whitelist *whitelist.Whitelist
	history   *history.History
	Control   chan int
	Done      chan bool
}

func Setup(l logger.Log, c config.Config) (err error) {
	Log = l
	Config = c

	return
}

func New(r []*regexp.Regexp, bl *blacklist.Blacklist, wl *whitelist.Whitelist, h *history.History) *Pipeline {
	var pipeline *Pipeline

	pipeline = &Pipeline{
		ruleset:   r,
		blacklist: bl,
		whitelist: wl,
		history:   h,
		Control:   make(chan int, config.D_CONTROL_BUFSIZE),
		Done:      make(chan bool, config.D_DONE_BUFSIZE),
	}

	return pipeline
}