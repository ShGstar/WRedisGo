local result = redis.call('get',KEYS[1])
return result