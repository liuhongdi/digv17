package service

import (
	//"github.com/go-redis/redis/v8"
	"github.com/liuhongdi/digv17/cache"
)

//得到一个id
func GetOneId(idType string) (int64, error) {
	//get from cache
	id,err := cache.GetOneId(idType);
	return id,err
}