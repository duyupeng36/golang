# Redis数据库介绍
是一种高性能的`Key-Value`数据库

## NoSQL介绍
`NoSQL`：一类新出现的数据库(`not only sql`)，它的特点：
1. 不支持`SQL`语法
   
2. 存储结构跟传统关系型数据库中的那种关系表完全不同，
   `nosql`中存储的数据都是`Key-Value`形式

3. `NoSQL`的世界中没有一种通用的语言，每种`nosql`数据库都有自己的`api`和语法，
   以及擅长的业务场景

### NoSQL中的产品种类相当多：
1. Mongodb
2. Redis
3. Hbase hadoop
4. Cassandra hadoop

### NoSQL和SQL数据库的比较
* 适用场景不同：`sql`数据库适合用于关系特别复杂的数据查询场景，`nosql`反之

* 两者在不断地取长补短，呈现融合趋势


## Redis简介
* `Redis`是一个开源的使用`ANSI C`语言编写、支持网络、可基于内存亦可持久化的日志型、
  `Key-Value`数据库，并提供多种语言的`API`。从`2010`年`3`月`15`日起，
  `Redis`的开发工作由`VMware`主持。从`2013`年`5`月开始，`Redis`的开发由`Pivotal`赞助。

* `Redis`是 `NoSQL`技术阵营中的一员，它通过多种键值数据类型来适应不同场景下的存储需求，
  借助一些高层级的接口使用其可以胜任，如缓存、队列系统的不同角色

## Redis特性
`Redis` 与其他 `key - value` 缓存产品有以下三个特点：
* `Redis`支持数据的持久化，可以将内存中的数据保存在磁盘中，
  重启的时候可以再次加载进行使用。
* `Redis`不仅仅支持简单的`key-value`类型的数据，同时还提供`list，set，zset，hash`等
  数据结构的存储。
* `Redis`支持数据的备份，即`master-slave`模式的数据备份。

## Redis 优势
* 性能极高: `Redis`能读的速度是`110000次/s`,写的速度是`81000次/s` 。
* 丰富的数据类型: `Redis`支持二进制案例的 `Strings, Lists, Hashes, Sets` 及 
  `Ordered Sets` 数据类型操作。
* 原子性: `Redis`的所有操作都是原子性的，同时`Redis`还支持对几个操作全并后的原子性执行。
* 丰富的特性: `Redis`还支持 `publish/subscribe`, 通知, `key` 过期等等特性。


## redis应用场景
* 用来做缓存(`ehcache/memcached`)——`redis`的所有数据是放在内存中的（内存数据库）
  
* 可以在某些特定应用场景下替代传统数据库——比如社交类的应用

* 在一些大型系统中，巧妙地实现一些特定的功能：`session共享`、`购物车`


# Redis安装

## Windows安装
1. 在https://github.com/tporadowski/redis/releases 下载`.msi`文件
2. 依次下一步安装

## Linux安装

1. 下载
    ```shell
    wget https://download.redis.io/releases/redis-3.2.8.tar.gz
    ```

2. 解压缩

    ```shell
    mkdir /usr/local/redis
    tar -zxvf redis-3.2.8.tar.gz
    ```

3. 编译
   
    ```shell
    cd redis-3.2.8
    make
    ```

4. 等待编译完成，测试环境
   
    ```shell
    make test
    ```

5. 如果出现保存，但报错提示下载需要的文件，然后在进行测试

6. 测试通过后，然后安装到指定目录

    ```shell
    make PREFIX=/usr/local/redis install
    ```
    该命令会将编译好的二进制文件复制到`/usr/local/redis`目录下

**文件介绍**
1. redis-server      redis服务器
2. redis-cli         redis命令行客户端
3. redis-benchmark   redis性能测试工具
4. redis-check-aof   AOF文件修复工具
5. redis-check-rdb   RDB文件检索工具


### 基本配置
在`/usr/local/redis`目录下新建如下几个目录
* `etc`: 保存配置文件
* `data`: 保存数据
* `log`: 保存日志


* 复制配置文件到`/usr/local/redis/etc`目录下
    ```shell
    mkdir /etc/redis
    cp redis.conf /usr/local/redis/etc
    mv redis.conf redis_6379.conf
    ```

* 查看配置文件
    
    ```shell
    vi /usr/local/redis/etc/redis_6379.conf
    ```

* 绑定ip：如果需要远程访问，可将此⾏注释，或绑定⼀个真实ip
    ```shell
    bind 0.0.0.0
    ```
    * 接收所有ip的连接

* 端⼝，默认为6379
    ```shell
    port 6379
    ```

* 是否以守护进程运⾏

    * 如果以守护进程运⾏，则不会在命令⾏阻塞，类似于服务  
    * 如果以⾮守护进程运⾏，则当前终端被阻塞
    * 设置为`yes`表示守护进程，设置为`no`表示⾮守护进程  
    * 推荐设置为`yes`

    ```shell
    daemonize yes
    ```

* 数据⽂件
    ```shell
    dbfilename dump.rdb
    ```

* 数据⽂件存储路径
    ```shell
    dir /usr/local/redis/data
    ```

* ⽇志⽂件
    ```shell
    logfile /usr/local/redis/log/redis_6379.log
    ```

* 数据库，默认有16个
    ```shell
    databases 16
    ```

### 自启动配置
在系统服务目录`/usr/lib/systemd/system/`里创建`redis.service`文件，
在文件写入如下内容。重启系统即可开机启动
```shell
[Unit]
Description=redis-server
After=network.target

[Service]
Type=forking
ExecStart=/usr/local/redis/bin/redis-server   /usr/local/redis/etc/redis_6379.conf
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```
`[Unit]`各项参数说明：
* Description:描述服务
* After:描述服务在哪些基础服务启动后再启动

`[Service]`服务运行参数的设置
* Type=forking 是最简单和速度最快的选择
* ExecStart 为启动服务的具体运行命令
* ExecReload 为重启命令
* ExecStop 为停止命令
* PrivateTmp=True 表示给服务分配独立的临时空间

注意：`[Service]`的启动、重启、停止命令全部要求使用绝对路径

`[Install]`运行级别下服务安装的相关设置，可设置为多用户，即系统运行级别为3

**配置开机自启**
```shell
systemctl enable redis
```

**测试**
```shell
systemctl start redis  # 启动
systemctl stop redis  # 停止
systemctl restart redis  # 重启

systemctl status redis  # 查看信息
```


