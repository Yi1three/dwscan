package input

import (
	"github.com/projectdiscovery/goflags"
)

type Options struct {
	Urls    string
	UrlFile string

	Ports string

	Threads int
}

func ParseOptions() *Options {
	options := &Options{}

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("this is dwscan.")

	flagSet.CreateGroup("parse", "parse targets",
		flagSet.StringVarP(&options.Urls, "url", "u", "", "target urls (separated by ',' )"),
		flagSet.StringVarP(&options.UrlFile, "url-file", "uf", "", "target urls' file. Urls in file separated by ','"),
		flagSet.StringVarP(&options.Ports, "port", "p", "80", "target ports, separated by ','"),
	)

	flagSet.CreateGroup("config", "Configuration",
		flagSet.IntVarP(&options.Threads, "thread", "t", 25, "num of threads"),
	)

	flagSet.Parse()

	return options
}
