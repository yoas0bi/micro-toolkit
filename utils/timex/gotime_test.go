package timex_test

import (
	"testing"
	"time"

	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
	"github.com/yoas0bi/micro-toolkit/utils/timex"
)

func TestAddSeconds(t *testing.T) {
	nw := time.Now()

	nt := timex.AddSec(nw, 10)
	assert.True(t, nt.After(nw))
	nt = timex.AddSeconds(nw, 10)
	assert.True(t, nt.After(nw))

	// add hour
	nt = timex.AddHour(nw, 1)
	assert.True(t, nt.After(nw))

	// add day
	nt = timex.AddDay(nw, 1)
	assert.True(t, nt.After(nw))

	// add minutes
	nt = timex.AddMinutes(nw, 1)
	assert.True(t, nt.After(nw))
}
