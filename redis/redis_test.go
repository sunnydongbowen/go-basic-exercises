package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var rdb *redis.Client // 声明一个全局的redis连接对象

// 初始化连接
func initClient() (err error) {
	// 此处应该是初始化全局的redis连接对象
	rdb = redis.NewClient(&redis.Options{
		Addr: "192.168.72.130:6379",
		// 下面都是默认值
		Password: "000415",
		DB:       0,   //use default DB
		PoolSize: 100, // 连接池大小
	})
	ctx, concel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer concel()
	_, err = rdb.Ping(ctx).Result()
	return err
}

func demo1() {
	//设置值
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	// 最后那个要改成0，不然设置不成功的，因为docker里的时区关系
	err := rdb.Set(ctx, "name", "bowen", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	//取值,获取值 类的命令后面一般用 Result()
	v, err := rdb.Get(context.Background(), "name").Result()
	if err != nil {
		// 排除掉key不存在的场景
		if err == redis.Nil {
			// 返回的err是key不存在时...
		}
		fmt.Println(err)
		return
	}
	fmt.Println(v, err)

	//我只想用value,如果出错了就用默认值
	fmt.Println("------")
	fmt.Printf("Err()==redis.Nil:%#v\n", rdb.Get(context.Background(), "namewwww").Err() == redis.Nil)
	fmt.Printf("Err()==nil:%#v\n", rdb.Get(context.Background(), "namexxxxx").Err() == nil)
	fmt.Printf("Val():%#v\n", rdb.Get(context.Background(), "nameeee").Val())
	nv, nerr := rdb.Get(context.Background(), "nameewww").Result()
	fmt.Printf("Result():%#v %#v\n", nv, nerr)
}

func TestDemo(t *testing.T) {
	//初始化全局的redis连接对象rdb
	if err := initClient(); err != nil {
		fmt.Println("init redis client failed,err:", err)
		return
	}
	fmt.Println("连接redis成功啦！")
	demo1()
}

func Hset() {
	err := rdb.HSet(context.Background(), "博文", "score", 60).Err()
	if err != nil {
		fmt.Println(err)
	}

	err = rdb.HSet(context.Background(), "博文", "weight", 37).Err()
	if err != nil {
		fmt.Println(err)
	}

	// 取hash中的一个键值对
	v, err := rdb.HGet(context.Background(), "博文", "weight").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	//fmt.Println(v, err)

	//
	v2, err := rdb.HMGet(context.Background(), "博文", "score", "weight").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v2)

	v3, err := rdb.HGetAll(context.Background(), "博文").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v3)

}

func TestHSet(t *testing.T) {
	//初始化全局的redis连接对象rdb
	if err := initClient(); err != nil {
		fmt.Println("init redis client failed,err:", err)
		return
	}
	fmt.Println("连接redis成功啦！")
	Hset()
}

func redisExample2() {
	zsetKey := "language:rank"
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	//ZADD
	err := rdb.ZAdd(context.Background(), zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	// 把GoLang的分数加10
	err = rdb.ZIncrBy(context.Background(), zsetKey, 10.0, "Golang").Err()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}

	//取分数最高的3个
	fmt.Println("==================================")
	ret, err := rdb.ZRevRangeWithScores(context.Background(), zsetKey, 0, 3).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	//取95到100分的
	fmt.Println("==================================")
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(context.Background(), zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func TestZset(t *testing.T) {
	//初始化全局的redis连接对象rdb
	if err := initClient(); err != nil {
		fmt.Println("init redis client failed,err:", err)
		return
	}
	fmt.Println("连接redis成功啦！")
	redisExample2()
}
