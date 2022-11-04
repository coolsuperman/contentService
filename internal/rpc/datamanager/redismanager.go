package datamanager

import (
	"contentService/pkg/config"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var once sync.Once
var client *RedisHelper

//RedisHelper is a global redis client object
type RedisHelper struct {
	Db *redis.Client
}

func GetRedisInstance() *RedisHelper {
	once.Do(func() {
		client = &RedisHelper{}
		addr := config.Config.Redis.Host
		db := config.Config.Redis.DB
		passwd := config.Config.Redis.Passwd
		if addr == "" {
			panic("Init redis fail")
		}

		client.Db = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: passwd,
			DB:       db,
		})
	})

	return client
}

func NewRedisHelper(conf config.GConfig) (*RedisHelper, func(), error) {
	client = &RedisHelper{}
	addr := conf.Redis.Host
	db := conf.Redis.DB
	passwd := conf.Redis.Passwd
	if addr == "" {
		return nil, func() {}, fmt.Errorf("Init redis fail")
	}

	client.Db = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	return client, func() {}, nil
}

func (r *RedisHelper) HSet(key, field string, value interface{}) (bool, error) {
	return r.Db.HSet(key, field, value).Result()
}

func (r *RedisHelper) HGet(key string, field string) (string, error) {
	return r.Db.HGet(key, field).Result()
}

func (r *RedisHelper) HMGet(key string, field ...string) ([]interface{}, error) {
	return r.Db.HMGet(key, field...).Result()
}

func (r *RedisHelper) HMSet(key string, fields map[string]interface{}) (string, error) {
	return r.Db.HMSet(key, fields).Result()
}

func (r *RedisHelper) HGetAll(key string, field ...string) (map[string]string, error) {
	return r.Db.HGetAll(key).Result()
}

func (r *RedisHelper) ZAdd(key string, field ...redis.Z) (int64, error) {
	return r.Db.ZAdd(key, field...).Result()
}

func (r *RedisHelper) ZCard(key string) (int64, error) {
	return r.Db.ZCard(key).Result()
}

func (r *RedisHelper) ZIncr(key string, field redis.Z) (float64, error) {
	return r.Db.ZIncr(key, field).Result()
}

func (r *RedisHelper) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return r.Db.ZIncrBy(key, increment, member).Result()
}

func (r *RedisHelper) ZRevRangeByLex(key string, opt redis.ZRangeBy) ([]string, error) {
	return r.Db.ZRevRangeByLex(key, opt).Result()
}

func (r *RedisHelper) ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return r.Db.ZRevRangeWithScores(key, start, stop).Result()
}

func (r *RedisHelper) ZRevRange(key string, start, stop int64) ([]string, error) {
	return r.Db.ZRevRange(key, start, stop).Result()
}

func (r *RedisHelper) RPush(key string, values ...interface{}) (int64, error) {
	return r.Db.RPush(key, values...).Result()
}

func (r *RedisHelper) LRange(key string, start, stop int64) ([]string, error) {
	return r.Db.LRange(key, start, stop).Result()
}

func (r *RedisHelper) Expire(key string, expiration time.Duration) (bool, error) {
	return r.Db.Expire(key, expiration).Result()
}

func (r *RedisHelper) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return r.Db.SetNX(key, value, expiration).Result()
}

func (r *RedisHelper) Exists(key string) (int64, error) {
	return r.Db.Exists(key).Result()
}

func (r *RedisHelper) Get(key string) (string, error) {
	return r.Db.Get(key).Result()
}

func (r *RedisHelper) Del(key string) (int64, error) {
	return r.Db.Del(key).Result()
}

func (r *RedisHelper) EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return r.Db.EvalSha(sha1, keys, args...).Result()
}

func (r *RedisHelper) ScriptLoad(script string) (string, error) {
	return r.Db.ScriptLoad(script).Result()
}

func (r *RedisHelper) ZScore(key, member string) (float64, error) {
	return r.Db.ZScore(key, member).Result()
}

func (r *RedisHelper) TxPipeline() redis.Pipeliner {
	return r.Db.TxPipeline()
}

func (r *RedisHelper) SAdd(key string, members ...interface{}) (int64, error) {
	return r.Db.SAdd(key, members...).Result()
}

func (r *RedisHelper) SMembers(key string) ([]string, error) {
	return r.Db.SMembers(key).Result()
}

func (r *RedisHelper) SRandMember(key string) (string, error) {
	return r.Db.SRandMember(key).Result()
}

func (r *RedisHelper) SRandMemberN(key string, num int64) ([]string, error) {
	return r.Db.SRandMemberN(key, num).Result()
}

func (r *RedisHelper) SRem(key string, members ...interface{}) (int64, error) {
	return r.Db.SRem(key, members...).Result()
}

func (r *RedisHelper) SCard(key string) (int64, error) {
	return r.Db.SCard(key).Result()
}

func (r *RedisHelper) ZRange(key string, start, stop int64) ([]string, error) {
	return r.Db.ZRange(key, start, stop).Result()
}
