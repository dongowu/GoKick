package ratelimit

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Rule 限流规则
type Rule struct {
	Method  string        // HTTP 方法：GET/POST/PUT/DELETE
	Path    string        // 路由路径
	Limit   int           // 时间窗口内最大请求数
	Window  time.Duration // 时间窗口
}

// Store 限流存储（基于 Redis）
type Store struct {
	rdb *redis.Client
}

// NewStore 创建限流存储
func NewStore(rdb *redis.Client) *Store {
	return &Store{rdb: rdb}
}

// Allow 检查是否允许请求通过
func (s *Store) Allow(ctx context.Context, rule Rule, key string) (bool, error) {
	// 使用 Redis Lua 脚本实现原子限流
	// 格式：gokick:ratelimit:{method}:{path}:{key}
	redisKey := fmt.Sprintf("gokick:ratelimit:%s:%s:%s", rule.Method, rule.Path, key)

	// Lua 脚本：实现滑动窗口限流
	// 使用 Redis Sorted Set，score 为时间戳，member 为随机值
	luaScript := `
		local key = KEYS[1]
		local limit = tonumber(ARGV[1])
		local window = tonumber(ARGV[2])
		local now = tonumber(ARGV[3])

		-- 移除窗口外的旧记录
		local minScore = now - window
		redis.call('ZREMRANGEBYSCORE', key, '-inf', minScore)

		-- 获取当前窗口内的请求数
		local count = redis.call('ZCARD', key)

		if count >= limit then
			return 0 -- 拒绝
		else
			-- 添加当前请求记录（使用随机 member 避免重复）
			local member = redis.call('TIME')[1] .. redis.call('TIME')[2]
			redis.call('ZADD', key, now, member)
			-- 设置过期时间（避免内存泄漏）
			redis.call('EXPIRE', key, window)
			return 1 -- 允许
		end
	`

	result, err := s.rdb.Eval(ctx, luaScript, []string{redisKey},
		rule.Limit,
		int64(rule.Window.Seconds()),
		time.Now().Unix(),
	).Result()

	if err != nil {
		return false, err
	}

	allowed, ok := result.(int64)
	if !ok {
		return false, fmt.Errorf("unexpected result type from lua script")
	}

	return allowed == 1, nil
}

// Reset 重置某个 key 的限流计数
func (s *Store) Reset(ctx context.Context, rule Rule, key string) error {
	redisKey := fmt.Sprintf("gokick:ratelimit:%s:%s:%s", rule.Method, rule.Path, key)
	return s.rdb.Del(ctx, redisKey).Err()
}
