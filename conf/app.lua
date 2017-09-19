local weihu =  ngx.shared.weihu
local host = ngx.var.host
local method = ngx.var.request_method
local key = string.format("%s|%s",host,method)
local s,_ = weihu:get(key)
if s then
 ngx.header.content_type="application/json;charset=utf8"
 ngx.say(s)
else
 ngx.say("not match")
end