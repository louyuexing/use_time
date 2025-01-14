package use_time

import (
	"testing"
	"time"
)

func TestUseTime(t *testing.T) {
	ut := New(time.Now())
	time.Sleep(2 * time.Second)
	t.Log(ut.Step().String())
	time.Sleep(2 * time.Second)
	t.Log(ut.Cross().String())
}
