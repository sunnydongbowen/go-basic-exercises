package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()

// pipelineDemo 获取pipeline结果的第一种方式
// 适合命令不多，命令返回的结果格式不太一样
func pipeline() {
	pipe := rdb.Pipeline()
	c1 := pipe.Get(ctx, "name")
	c2 := pipe.HMGet(ctx, "博文", "score", "weight")
	incr := pipe.Incr(ctx, "pipeline_counter") //自动创建了这个key Redis Incr 命令将 key 中储存的数字值增一。
	pipe.Expire(ctx, "pipeline_counter", time.Hour)
	// 执行命令
	_, err := pipe.Exec(ctx)
	// 取结果，方式1
	fmt.Println(c1.String()) // 这里面注意类型要对
	fmt.Println(c2.Val())
	fmt.Println(incr.Val(), err)
}

func TestPipeline(t *testing.T) {
	if err := initClient(); err != nil {
		fmt.Println("init redis client failed,err:", err)
		return
	}
	fmt.Println("连接redis成功啦！")
	pipeline()
}

// pipelineDemo2 获取pipeline结果的第二种方式
// 适合批量执行相同类型的命令（返回值类型一致）
func pipeline2() {
	pipe := rdb.Pipeline()
	// 输入命令
	for i := 0; i < 10; i++ {
		pipe.Get(ctx, fmt.Sprintf("key%d", i))
	}
	// 执行命令
	cmders, err := pipe.Exec(ctx)

	//if err != nil {
	//	// 如果是这个错，就弹出来key不存在
	//	if err == redis.Nil {
	//		fmt.Println("key不存在")
	//		return
	//	}
	//	//其他错就直接打印原本的错误
	//	fmt.Println(err)
	//	return
	//}
	// 有错误，并且不是这个错误
	if err != nil && err != redis.Nil {
		fmt.Println(err)
		return
	}
	// 取结果
	for _, cmder := range cmders {
		v := cmder.String()
		// 因为拿到的cmder是一个接口类型
		// 需要自己根据上面的命令的返回值进行类型断言
		switch cmder.(type) {
		case *redis.StringCmd:
		case *redis.IntCmd:
			//。。。。

		}
		// cmder.(*redis.StringCmd).Result()
		// cmder.(*redis.StringSliceCmd).Result()
		fmt.Println(v)

	}
}

func TestPipeline2(t *testing.T) {
	if err := initClient(); err != nil {
		fmt.Println("init redis client failed,err:", err)
		return
	}
	fmt.Println("连接redis成功啦！")
	pipeline2()
}
