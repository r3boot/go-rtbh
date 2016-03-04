package pipeline

import (
	"github.com/r3boot/go-rtbh/config"
	"github.com/r3boot/go-rtbh/lists"
	"github.com/r3boot/rlib/logger"
)

type Pipeline struct {
	Control chan int
	Done    chan bool
}

var Log logger.Log
var Config *config.Config

var Whitelist *lists.Whitelist
var Blacklist *lists.Blacklist

func Setup(l logger.Log, cfg *config.Config) (err error) {
	Log = l
	Config = cfg

	return
}

func NewPipeline() (pl *Pipeline, err error) {
	pl = &Pipeline{}
	Whitelist = lists.NewWhitelist()
	Blacklist = lists.NewBlacklist()

	return
}

func (pl *Pipeline) Startup(input chan []byte) (err error) {
	var stop_loop bool
	var event *Event

	stop_loop = false
	for {
		if stop_loop {
			break
		}

		select {
		case data := <-input:
			{
				if event, err = NewEvent(data); err != nil {
					Log.Warning("[Pipeline] NewEvent: " + err.Error())
					continue
				}

				if event.Address == "" {
					// Log.Debug("[Pipeline]: Failed to parse event: " + string(data))
					continue
				}

				if Whitelist.Listed(event.Address) {
					Log.Warning("[Pipeline]: Host " + event.Address + " is on whitelist")
					continue
				}

				if Blacklist.Listed(event.Address) {
					Log.Warning("[Pipeline]: Host " + event.Address + " is already listed")
					continue
				}

				Log.Debug(event)
			}
		case cmd := <-pl.Control:
			{
				switch cmd {
				case config.CTL_SHUTDOWN:
					{
						Log.Debug("Shutting down pipeline")
						stop_loop = true
						continue
					}
				}
			}
		}
	}

	pl.Done <- true

	return
}