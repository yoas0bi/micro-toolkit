package textscan_test

import (
	"testing"

	"github.com/yoas0bi/micro-toolkit/utils/dump"
	"github.com/yoas0bi/micro-toolkit/utils/strutil/textscan"
	"github.com/yoas0bi/micro-toolkit/utils/testutil/assert"
)

func TestParser_ParseText(t *testing.T) {
	p := textscan.NewParser(func(t textscan.Token) {
		dump.P(t)
	})

	err := p.ParseText(`
# comments 1
# comments 1.1
# comments 1.2
name = inhere

// comments 2
age = 28

/*
multi line
comments 3
*/
desc = '''
a multi
line string
'''
`)
	assert.NoErr(t, err)

}
