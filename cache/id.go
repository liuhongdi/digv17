package cache

import (
	"context"
	"github.com/liuhongdi/digv17/global"
	"strconv"
	"github.com/go-redis/redis/v8"
)

//得到cache的名字
func getIdCacheName(idType string) (string) {
	return "id_"+idType
}

//得到一个id
func GetOneId(idType string) (int64,error) {
	key := getIdCacheName(idType);
	luaId := redis.NewScript(`
local id_key = KEYS[1]
local current = redis.call('get',id_key)
if current == false then
    redis.call('set',id_key,1)
    return '1'
end
--redis.log(redis.LOG_NOTICE,' current:'..current..':')
local result =  tonumber(current)+1
--redis.log(redis.LOG_NOTICE,' result:'..result..':')
redis.call('set',id_key,result)
return tostring(result)
	`)
	var ctx = context.Background()
	n, err := luaId.Run(ctx, global.RedisDb, []string{key}, 2).Result()

    if (err != nil) {
    	return -1,err
	} else {
		var ret string = n.(string)
		retint,err := strconv.ParseInt(ret, 10, 64)
		if (err == nil) {
			return retint,err
		} else {
			return -1,err
		}
	}
}
