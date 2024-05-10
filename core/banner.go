package core

import (
	"github.com/boy-hack/ksubdomain/core/conf"
	"github.com/boy-hack/ksubdomain/core/gologger"
)

const banner = ``

func ShowBanner() {
	gologger.Printf(banner)
	gologger.Infof("Current Version: %s\n", conf.Version)
}
