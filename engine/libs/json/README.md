### 介绍

JSON解析库,Fork From [gpher-json](https://github.com/layeh/gopher-json)

```bash
CommitID: 552bb3c4c3bf4b9df70a2ad8ec555143f7e1b812
```

### 使用
解码
```lua
local json = require("json")
local data = json.decode('{"a": "b"}')
print(data["a"])
```

编码
```lua
local json = require("json")
local data = json.encode({"apple", "banana", "orange"})
print(data)

local data = json.encode({
    a="apple", b="banana", o="orange"
})
print(data)
```