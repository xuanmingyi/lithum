package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
	"testing"
)

func NewEvent(t *testing.T) {
	L := lua.NewState()
	newEvent(L)
	err := L.DoString(`
event:set("key1", "val1")
event:set("key2", 222)
event:set("key3", true)

event:dump()

event:set("[key1]", "val2")

event:dump()
`)
	if err != nil {
		fmt.Println(err)
	}
}

func TestSplitName(t *testing.T) {
	assert := assert.New(t)
	assert.Len(splitName("sss"), 1)
	assert.Len(splitName("[sss][bbb]"), 2)
	assert.Len(splitName("[sss]"), 1)

	assert.Len(splitName("bbb[sss]"), 0)
	assert.Len(splitName("[sss]bbb"), 0)
	assert.Len(splitName("[sss"), 0)
	assert.Len(splitName("[sss]]"), 0)
	assert.Len(splitName("[[sss]]"), 0)
	assert.Len(splitName("[s[s]s]"), 0)
	assert.Len(splitName(""), 0)
	assert.Len(splitName("   "), 0)
}
