[TOC]

# ip2region开源库

## Ip2region是什么？

ip2region - 准确率99.9%的离线IP地址定位库，0.0x毫秒级查询，ip2region.db数据库只有数MB，提供了java,php,c,python,nodejs,golang,c#等查询绑定和Binary,B树,内存三种查询算法。

## Ip2region特性

### 99.9%准确率

数据聚合了一些知名ip到地名查询提供商的数据，这些是他们官方的的准确率，经测试着实比经典的纯真IP定位准确一些。
ip2region的数据聚合自以下服务商的开放API或者数据(升级程序每秒请求次数2到4次):
01, >80%, 淘宝IP地址库, http://ip.taobao.com/
02, ≈10%, GeoIP, https://geoip.com/
03, ≈2%, 纯真IP库, http://www.cz88.net/
**备注：**如果上述开放API或者数据都不给开放数据时ip2region将停止数据的更新服务。

### 标准化的数据格式

每条ip数据段都固定了格式：

```
_城市Id|国家|区域|省份|城市|ISP_
```

只有中国的数据精确到了城市，其他国家有部分数据只能定位到国家，后前的选项全部是0，已经包含了全部你能查到的大大小小的国家（请忽略前面的城市Id，个人项目需求）。

### 体积小

包含了全部的IP，生成的数据库文件ip2region.db只有几MB，最小的版本只有1.5MB，随着数据的详细度增加数据库的大小也慢慢增大，目前还没超过8MB。

### 查询速度快

全部的查询客户端单次查询都在0.x毫秒级别，内置了三种查询算法

1. memory算法：整个数据库全部载入内存，单次查询都在0.1x毫秒内，C语言的客户端单次查询在0.00x毫秒级别。
2. binary算法：基于二分查找，基于ip2region.db文件，不需要载入内存，单次查询在0.x毫秒级别。
3. b-tree算法：基于btree算法，基于ip2region.db文件，不需要载入内存，单词查询在0.x毫秒级别，比binary算法更快。

**任何客户端b-tree都比binary算法快，当然memory算法固然是最快的！**



## golang 实现ip地址查询

### 获取

```bash
go get github.com/lionsoul2014/ip2region/binding/golang
```

### 使用

```go
package main

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func main() {
	fmt.Println("err")
	region, err := ip2region.New("ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	ip, err := region.MemorySearch("127.0.0.1")
	fmt.Println(ip, err)
	ip, err = region.BinarySearch("127.0.0.1")
	fmt.Println(ip, err)
	ip, err = region.BtreeSearch("127.0.0.1")
	fmt.Println(ip, err)
}
```

### 返回对象

```go
type IpInfo struct {
	CityId   int64
	Country  string
	Region   string
	Province string
	City     string
	ISP      string
}
```

### 性能

| 名称                    | 次数    | 平均耗时    |
| ----------------------- | ------- | ----------- |
| BenchmarkBtreeSearch-4  | 200000  | 7715 ns/op  |
| BenchmarkMemorySearch-4 | 2000000 | 840 ns/op   |
| BenchmarkBinarySearch-4 | 30000   | 42680 ns/op |

### 测试程序

```bash
cd /binging/golang

go run main.go ../../data/ip2region.db

Or

go build -o ip2region main.go
./ip2region ../../data/ip2region.db
```

会看到如下界面

```bash
initializing
+-------------------------------------------------------+
| ip2region test script                                 |
| format 'ip type'                                      |
| type option 'b-tree','binary','memory' default b-tree |
| Type 'quit' to exit program                           |
+-------------------------------------------------------+
ip2reginon >> 127.0.0.1 memory
0|未分配或者内网IP|0|0|0|0  960.5µs
```