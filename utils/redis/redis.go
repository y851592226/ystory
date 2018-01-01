package redis

import "github.com/astaxie/beego"
import "github.com/garyburd/redigo/redis"
import "time"

func newPool(addr string, MaxIdle, MaxActive int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	pool *redis.Pool
)

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	rc := pool.Get()
	if err!=nil{
		panic(err)
	}  
	defer rc.Close()
	reply, err = rc.Do(commandName, args...)
	return
}

func init() {
	redisaddr := beego.AppConfig.String("redisaddr")
	redisidle, err := beego.AppConfig.Int("redisidle")
	if err != nil {
		panic(err)
	}
	redisacti, err := beego.AppConfig.Int("redisacti")
	if err != nil {
		panic(err)
	}
	pool = newPool(redisaddr, redisidle, redisacti)
}

func Int(commandName string, args ...interface{}) (int, error) {
	return redis.Int(Do(commandName, args...))
}
func Int64(commandName string, args ...interface{}) (int64, error) {
	return redis.Int64(Do(commandName, args...))
}
func Uint64(commandName string, args ...interface{}) (uint64, error) {
	return redis.Uint64(Do(commandName, args...))
}
func Float64(commandName string, args ...interface{}) (float64, error) {
	return redis.Float64(Do(commandName, args...))
}
func String(commandName string, args ...interface{}) (string, error) {
	return redis.String(Do(commandName, args...))
}
func Bytes(commandName string, args ...interface{}) ([]byte, error) {
	return redis.Bytes(Do(commandName, args...))
}
func Bool(commandName string, args ...interface{}) (bool, error) {
	return redis.Bool(Do(commandName, args...))
}
func Values(commandName string, args ...interface{}) ([]interface{}, error) {
	return redis.Values(Do(commandName, args...))
}
func Strings(commandName string, args ...interface{}) ([]string, error) {
	return redis.Strings(Do(commandName, args...))
}
func ByteSlices(commandName string, args ...interface{}) ([][]byte, error) {
	return redis.ByteSlices(Do(commandName, args...))
}
func Ints(commandName string, args ...interface{}) ([]int, error) {
	return redis.Ints(Do(commandName, args...))
}
func StringMap(commandName string, args ...interface{}) (map[string]string, error) {
	return redis.StringMap(Do(commandName, args...))
}
func IntMap(commandName string, args ...interface{}) (map[string]int, error) {
	return redis.IntMap(Do(commandName, args...))
}
func Int64Map(commandName string, args ...interface{}) (map[string]int64, error) {
	return redis.Int64Map(Do(commandName, args...))
}
func Positions(commandName string, args ...interface{}) ([]*[2]float64, error) {
	return redis.Positions(Do(commandName, args...))
}
