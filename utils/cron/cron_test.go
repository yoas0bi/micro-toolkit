package cron

import (
	"github.com/yoas0bi/micro-toolkit/utils/helper"
	"testing"
)

func TestNewWithSecond(t *testing.T) {
	c := helper.TCorn.NewWithSecond()
	spec := "*/5 * * * * ?"
	i := 0
	_, err := c.AddFunc(spec, func() {
		i++
	})
	if err != nil {
		t.Errorf("cron errors: %v\n", err)
	}
	c.Start()
	c.Stop()
}
