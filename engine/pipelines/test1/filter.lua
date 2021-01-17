local json = require("json")
local regexp = require("regexp")

local result, err = regexp.find_all_string_submatch("\\d+\\.\\d+\\.\\d+\\.\\d+", message)

ips = {}

for k, v in ipairs(result) do
    if (v[1] ~= "255.255.255.0") then
        ips[k] = {ip=v[1]}
    end
end

output, err=[[ [{"ip": "192.168.6.1"}, {"ip": "127.0.0.1'"}]
]]