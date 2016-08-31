package orm

import (
	"github.com/r3boot/go-rtbh/config"
	"github.com/r3boot/rlib/logger"
	"gopkg.in/pg.v4"
)

const MYNAME string = "ORM"

var Config config.Config
var Log logger.Log

type ORM struct {
	Db *pg.DB
}

func Setup(l logger.Log, c config.Config) (err error) {
	Log = l
	Config = c

	return
}

func New() *ORM {
	var orm *ORM

	orm = &ORM{}
	orm.Connect()

	return orm
}