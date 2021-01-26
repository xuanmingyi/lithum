local json = require("json")
local regexp = require("regexp")

local body = event:get("body")
local result, err = regexp.find_all_string_submatch("你的外网IP地址是：(\\d+\\.\\d+\\.\\d+\\.\\d+)", body)

print(result[1][2])

sql = "INSERT INTO "
event:set("sqls", {})