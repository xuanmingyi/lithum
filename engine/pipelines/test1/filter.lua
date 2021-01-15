local regexp = require("regexp")

local result, err = regexp.find_all_string_submatch("\\d+\\.\\d+\\.\\d+\\.\\d+", message)

output={}

for k, v in ipairs(result) do
    if (v[1] ~= "255.255.255.0") then
        output[k] = {ip=v[1]}
    end
end
