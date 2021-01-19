package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	lua "github.com/yuin/gopher-lua"
	"testing"
)

func TestNewEvent(t *testing.T) {
	assert := assert.New(t)
	L := lua.NewState()
	event := NewEvent(L)
	err := L.DoString(`
event:set("key1", "val1")
event:set("key2", 222)
event:set("key3", true)
event:set("[key4][subkey]", "subval1")
event:set("[key1]", "val2")
event:set("key5", {"v1", "v2", "v3"})
event:set("key6", {v1="vv1", v2="vv2", v3="vv3"})
event:set("key6", {v1="vv1", v2="vv2", v3={"vvv1", "vvv2", "vvv3"}})

print(event:get("key1"))
--event:dump()
`)
	if err != nil {
		fmt.Println(err)
	}

	assert.EqualValues("val2", event.Get("key1"))
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
