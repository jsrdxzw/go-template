package app

import (
	"errors"
	entschema "example/be/gen/entschema"
	"example/be/service/extensions"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hashicorp/go-multierror"
)

var ErrAppInit = errors.New("app initialize error")

var (
	Log       *extensions.LoggerExt
	Ent       *extensions.EntExt
	EntClient *entschema.Client
)

type Ext struct {
	Key       string
	Extension extensions.Extension
}

func getExtensions(conf Config, env string) []Ext {
	Log = extensions.NewLogExt(conf.LOG)
	Ent = extensions.NewEntExt(conf.DB)

	return []Ext{
		{Key: "log", Extension: Log},
		{Key: "ent", Extension: Ent},
	}
}

func InitExtensions(conf Config, env string) error {
	var allErrs error
	for _, ext := range getExtensions(conf, env) {
		if err := ext.Extension.Init(); err != nil {
			allErrs = multierror.Append(allErrs, fmt.Errorf("ErrAppInit %w: %s", ErrAppInit, ext.Key), err)
		}
	}
	if allErrs == nil {
		var ok bool
		EntClient, ok = Ent.Client.(*entschema.Client)
		if !ok {
			return fmt.Errorf("ErrAppInit %w: fail to initialize ent client", ErrAppInit)
		}
	}
	return allErrs
}
