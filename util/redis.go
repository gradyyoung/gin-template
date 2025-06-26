package util

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"ygang.top/gin-template/internal/config"
)

// RedisClient redis 操作封装客户端
type RedisClient struct {
	Client    *redis.Client
	KeyPrefix string
}

// NewRedisClient 创建 RedisClient
func NewRedisClient(client *redis.Client, appConfig *config.ApplicationConfig) *RedisClient {
	return &RedisClient{
		Client:    client,
		KeyPrefix: appConfig.Redis.KeyPrefix,
	}
}

var ctx = context.Background()

// getKey 拼接 Key 前缀
func (r *RedisClient) getKey(key string) string {
	return fmt.Sprintf("%s:%s", r.KeyPrefix, key)
}

/*------------------------------------ 通用 操作 ------------------------------------*/

// Keys 获取所有匹配的 key
func (r *RedisClient) Keys(pattern string) ([]string, error) {
	key := r.getKey(pattern)
	return r.Client.Keys(ctx, key).Result()
}

/*------------------------------------ 字符 操作 ------------------------------------*/

// Set 设置 key的值
func (r *RedisClient) Set(key, value string) error {
	key = r.getKey(key)
	return r.Client.Set(ctx, key, value, 0).Err()

}

// SetEX 设置 key 的值并指定过期时间
func (r *RedisClient) SetEX(key, value string, ex time.Duration) error {
	key = r.getKey(key)
	return r.Client.Set(ctx, key, value, ex).Err()
}

// Get 获取 key 的值
func (r *RedisClient) Get(key string) (string, error) {
	key = r.getKey(key)
	return r.Client.Get(ctx, key).Result()
}

// GetSet 设置新值获取旧值
func (r *RedisClient) GetSet(key, value string) (string, error) {
	key = r.getKey(key)
	return r.Client.GetSet(ctx, key, value).Result()
}

// Incr key值每次加一 并返回新值
func (r *RedisClient) Incr(key string) (int64, error) {
	key = r.getKey(key)
	return r.Client.Incr(ctx, key).Result()
}

// IncrBy key值每次加指定数值 并返回新值
func (r *RedisClient) IncrBy(key string, incr int64) (int64, error) {
	key = r.getKey(key)
	return r.Client.IncrBy(ctx, key, incr).Result()
}

// IncrByFloat key值每次加指定浮点型数值 并返回新值
func (r *RedisClient) IncrByFloat(key string, incrFloat float64) (float64, error) {
	key = r.getKey(key)
	return r.Client.IncrByFloat(ctx, key, incrFloat).Result()
}

// Decr key值每次递减 1 并返回新值
func (r *RedisClient) Decr(key string) (int64, error) {
	key = r.getKey(key)
	return r.Client.Decr(ctx, key).Result()
}

// DecrBy key值每次递减指定数值 并返回新值
func (r *RedisClient) DecrBy(key string, incr int64) (int64, error) {
	key = r.getKey(key)
	return r.Client.DecrBy(ctx, key, incr).Result()
}

// Del 删除 key
func (r *RedisClient) Del(key string) error {
	key = r.getKey(key)
	return r.Client.Del(ctx, key).Err()
}

// Expire 设置 key的过期时间
func (r *RedisClient) Expire(key string, ex time.Duration) error {
	key = r.getKey(key)
	return r.Client.Expire(ctx, key, ex).Err()
}

/*------------------------------------ list 操作 ------------------------------------*/

// LPush 从列表左边插入数据，并返回列表长度
func (r *RedisClient) LPush(key string, date ...interface{}) (int64, error) {
	key = r.getKey(key)
	return r.Client.LPush(ctx, key, date).Result()
}

// RPush 从列表右边插入数据，并返回列表长度
func (r *RedisClient) RPush(key string, date ...interface{}) (int64, error) {
	key = r.getKey(key)
	return r.Client.RPush(ctx, key, date).Result()
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func (r *RedisClient) LPop(key string) (string, error) {
	key = r.getKey(key)
	return r.Client.LPop(ctx, key).Result()
}

// RPop 从列表右边删除第一个数据，并返回删除的数据
func (r *RedisClient) RPop(key string) (string, error) {
	key = r.getKey(key)
	return r.Client.RPop(ctx, key).Result()
}

// LIndex 根据索引坐标，查询列表中的数据
func (r *RedisClient) LIndex(key string, index int64) (string, error) {
	key = r.getKey(key)
	return r.Client.LIndex(ctx, key, index).Result()
}

// LLen 返回列表长度
func (r *RedisClient) LLen(key string) (int64, error) {
	key = r.getKey(key)
	return r.Client.LLen(ctx, key).Result()
}

// LRange 返回列表的一个范围内的数据，也可以返回全部数据
func (r *RedisClient) LRange(key string, start, stop int64) ([]string, error) {
	key = r.getKey(key)
	return r.Client.LRange(ctx, key, start, stop).Result()
}

// LRem 从列表左边开始，删除元素data， 如果出现重复元素，仅删除 count次
func (r *RedisClient) LRem(key string, count int64, data interface{}) error {
	key = r.getKey(key)
	return r.Client.LRem(ctx, key, count, data).Err()
}

// LInsert 在列表中 pivot 元素的后面插入 data
func (r *RedisClient) LInsert(key string, pivot int64, data interface{}) error {
	key = r.getKey(key)
	err := r.Client.LInsert(ctx, key, "after", pivot, data).Err()
	return err
}

/*------------------------------------ set 操作 ------------------------------------*/

// SAdd 添加元素到集合中
func (r *RedisClient) SAdd(key string, data ...interface{}) error {
	key = r.getKey(key)
	return r.Client.SAdd(ctx, key, data).Err()
}

// SCard 获取集合元素个数
func (r *RedisClient) SCard(key string) (int64, error) {
	key = r.getKey(key)
	return r.Client.SCard(ctx, key).Result()
}

// SIsMember 判断元素是否在集合中
func (r *RedisClient) SIsMember(key string, data interface{}) (bool, error) {
	key = r.getKey(key)
	return r.Client.SIsMember(ctx, key, data).Result()
}

// SMembers 获取集合所有元素
func (r *RedisClient) SMembers(key string) ([]string, error) {
	key = r.getKey(key)
	return r.Client.SMembers(ctx, key).Result()
}

// SRem 删除 key集合中的 data元素
func (r *RedisClient) SRem(key string, data ...interface{}) error {
	key = r.getKey(key)
	return r.Client.SRem(ctx, key, data).Err()
}

// SPopN 随机返回集合中的 count个元素，并且删除这些元素
func (r *RedisClient) SPopN(key string, count int64) ([]string, error) {
	key = r.getKey(key)
	return r.Client.SPopN(ctx, key, count).Result()
}

/*------------------------------------ hash 操作 ------------------------------------*/

// HSet 根据 key和 field字段设置，field字段的值
func (r *RedisClient) HSet(key, field, value string) error {
	key = r.getKey(key)
	return r.Client.HSet(ctx, key, field, value).Err()
}

// HGet 根据 key 和 field 字段，查询 field 字段的值
func (r *RedisClient) HGet(key, field string) (string, error) {
	key = r.getKey(key)
	return r.Client.HGet(ctx, key, field).Result()
}

// HMGet 根据key和多个字段名，批量查询多个 hash 字段值
func (r *RedisClient) HMGet(key string, fields ...string) ([]interface{}, error) {
	key = r.getKey(key)
	return r.Client.HMGet(ctx, key, fields...).Result()
}

// HGetAll 根据 key 查询所有字段和值
func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	key = r.getKey(key)
	return r.Client.HGetAll(ctx, key).Result()
}

// HKeys 根据 key 返回所有字段名
func (r *RedisClient) HKeys(key string) ([]string, error) {
	key = r.getKey(key)
	return r.Client.HKeys(ctx, key).Result()
}

// HLen 根据 key，查询hash的字段数量
func (r *RedisClient) HLen(key string) (int64, error) {
	key = r.getKey(key)
	return r.Client.HLen(ctx, key).Result()
}

// HMSet 根据 key和多个字段名和字段值，批量设置 hash字段值
func (r *RedisClient) HMSet(key string, data map[string]interface{}) (bool, error) {
	key = r.getKey(key)
	return r.Client.HMSet(ctx, key, data).Result()
}

// HSetNX 如果 field字段不存在，则设置 hash字段值
func (r *RedisClient) HSetNX(key, field string, value interface{}) error {
	key = r.getKey(key)
	return r.Client.HSetNX(ctx, key, field, value).Err()
}

// HDel 根据 key和字段名，删除 hash字段，支持批量删除
func (r *RedisClient) HDel(key string, fields ...string) error {
	key = r.getKey(key)
	return r.Client.HDel(ctx, key, fields...).Err()
}

// HExists 检测 hash 字段名是否存在
func (r *RedisClient) HExists(key, field string) (bool, error) {
	key = r.getKey(key)
	return r.Client.HExists(ctx, key, field).Result()
}
