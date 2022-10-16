package service

import (
	corn "github.com/robfig/cron/v3"
)

func init() {
	go scheduled()
}

func newWithSeconds() *corn.Cron {
	secondParser := corn.NewParser(
		corn.Second | corn.Minute | corn.Hour | corn.Dom | corn.Month | corn.DowOptional | corn.Descriptor,
	)
	return corn.New(corn.WithParser(secondParser), corn.WithChain())
}

func scheduled() {
	c := newWithSeconds()
	c.AddFunc("0 0 * * * ?", userService.statisticalUserArea)
	c.AddFunc("0 0 0 * * ?", uniqueService.saveUniqueView)
	c.AddFunc("0 1 0 * * ?", uniqueService.clear)
	c.Start()
	select {}
}
