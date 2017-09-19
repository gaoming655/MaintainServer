local cjson = require "cjson"
local f = io.popen("curl -s http://yourserver/api")
local res = f:read("*a")
f:close()
local lua_value = cjson.decode(res)
local weihu = ngx.shared.weihu
weihu:flush_all()
for k,v in pairs(lua_value) do
        local key = string.format("%s|%s",v["host"],v["method"])
        weihu:set(key,v["result"])
end