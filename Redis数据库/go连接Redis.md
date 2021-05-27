# go连接Redis

## 使用go-redis客户端

`go-redis`客户端封装较为完善，对于每种数据类型

### 安装驱动

 ```shell
 go get -u github.com/go-redis/redis
 ```

### 连接数据库
#### 普通连接(连接一个redis数据库)
最新版本的`go-redis`库的相关命令都需要传递`context.Context`参数

```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "121.5.72.146:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
```

#### 连接哨兵
```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

func initClient()(err error){
	rdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
```

#### 连接集群
```go
package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// 声明一个全局的rdb变量
var rdb *redis.Client
func initClient()(err error){
	rdb = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
```

### 常用函数
#### 基本指令
* Keys():根据正则获取keys
* Type():获取key对应值得类型
* Del():删除缓存项
* Exists():检测缓存项是否存在
* Expire(),ExpireAt():设置有效期
* TTL(),PTTL():获取有效期
* DBSize():查看当前数据库key的数量
* FlushDB():清空当前数据
* FlushAll():清空所有数据库
#### 字符串(string)类型
* Set():设置
* SetEX():设置并指定过期时间
* SetNX():设置并指定过期时间
* Get():获取
* GetRange():字符串截取
* Incr():增加+1
* IncrBy():按指定步长增加
* Decr():减少-1
* DecrBy():按指定步长减少
* Append():追加
* StrLen():获取长度
#### 列表(list)类型
* LPush():将元素压入链表
* LInsert():在某个位置插入新元素
* LSet():设置某个元素的值
* LLen():获取链表元素个数
* LIndex():获取链表下标对应的元素
* LRange():获取某个选定范围的元素集
* LPop(): 从链表左侧弹出数据
* LRem():根据值移除元素
#### 集合(set)类型
* SAdd():添加元素
* SPop():随机获取一个元素
* SRem():删除集合里指定的值
* SSMembers():获取所有成员
* SIsMember():判断元素是否在集合中
* SCard():获取集合元素个数
* SUnion():并集,SDiff():差集,SInter():交集
#### 有序集合(zset)类型
* ZAdd():添加元素
* ZIncrBy():增加元素分值
* ZRange()、ZRevRange():获取根据score排序后的数据段
* ZRangeByScore()、ZRevRangeByScore():获取score过滤后排序的数据段
* ZCard():获取元素个数
* ZCount():获取区间内元素个数
* ZScore():获取元素的score
* ZRank()、ZRevRank():获取某个元素在集合中的排名
* ZRem():删除元素
* ZRemRangeByRank():根据排名来删除
* ZRemRangeByScore():根据分值区间来删除
#### 哈希(hash)类型
* HSet():设置
* HMset():批量设置
* HGet():获取某个元素
* HGetAll():获取全部元素
* HDel():删除某个元素
* HExists():判断元素是否存在
* HLen():获取长度

### 指令使用示例

#### Set/Get
```go
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient(ctx context.Context) (err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "121.5.72.146:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	err := initClient(ctx)
	if err != nil {
		fmt.Println("连接是失败")
		return
	}

	fmt.Println("连接成功")
	// Set/Get使用
	err = rdb.Set(ctx, "name", "dyp", 0).Err()
	if err != nil {
		fmt.Println("添加值失败")
		return
	}
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println("获取值失败")
		return
	}
	fmt.Println(val)
}
```

#### zadd示例
```go
func zadd(ctx context.Context) {
	zsetKey := "language_rank"
	languages := []*redis.Z{
		&redis.Z{Score: 90.0, Member: "Golang"},
		&redis.Z{Score: 98.0, Member: "Java"},
		&redis.Z{Score: 95.0, Member: "Python"},
		&redis.Z{Score: 97.0, Member: "JavaScript"},
		&redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := rdb.ZAdd(ctx, zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, &op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
```

### 直接执行命令
```go
res, err := rdb.Do(ctx, "set", "key", "value").Result()
```
