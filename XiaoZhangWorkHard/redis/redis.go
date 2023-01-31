package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	ctx, calfunc := context.WithTimeout(context.Background(), 10*time.Second)
	_, err = redisdb.Ping(ctx).Result()
	calfunc()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("connect redis failed,err:", err)
	}
	ctx, canl := context.WithTimeout(context.Background(), 10*time.Second)

	//执行redis 的set和get
	err = redisdb.Set(ctx, "score", 100, 10*time.Second).Err() //将score设置为100，超时时间为10秒
	defer canl()
	cmder := redisdb.Get(ctx, "score")                //得到score的值
	fmt.Println(cmder.Val())                          //获取值
	fmt.Println(cmder.Err())                          //获取错误
	cmder2, err := redisdb.Get(ctx, "score").Result() //获取值
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) { //redis自带nil来表示key不存在的错误
			return defaultValue, nil
		}
		// 出其他错了
		return "", err
	}
	fmt.Println(cmder2)
}

// doDemo rdb.Do 方法使用示例
func doDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 直接执行命令获取错误
	err := redisdb.Do(ctx, "set", "key", 10, "EX", 3600).Err()
	fmt.Println(err)

	// 执行命令获取结果
	val, err := redisdb.Do(ctx, "get", "key").Result()
	fmt.Println(val, err)
}

// zest
func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("connect redis failed,err:", err)
	}
	//zest
	key := "rank"
	items := []*redis.Z{ //redis.Z为众多redis数据结构中的一种
		&redis.Z{Score: 99, Member: "PHP"},
		&redis.Z{Score: 96, Member: "Golang"},
		&redis.Z{Score: 97, Member: "Python"},
		&redis.Z{Score: 99, Member: "Java"},
	}
	ctx, canl := context.WithTimeout(context.Background(), 10*time.Hour)
	defer canl()
	//把这些元素都追加到key
	redisdb.ZAdd(ctx, key, items...)
	newScore, err := redisdb.ZIncrBy(ctx, key, 10, "Golang").Result() //给golang加10分
	if err != nil {
		fmt.Println("zincrby failed,err:", err)
		return
	}
	fmt.Println("Golang's score is ", newScore)
	// 取分数最高的3个
	ret := redisdb.ZRevRangeWithScores(ctx, key, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(ctx, key, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
