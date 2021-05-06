<!--
 * @Descripttion: 
 * @version: 
 * @Author: WangShuaibing
 * @Date: 2020-09-27 09:26:23
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-12-17 10:16:16
-->
# Go 语言编程
> 静态语言
> Golang非常适合用来进行：服务器编程、分布式系统、数据库代理器、网络编程、内存数据库、云平台、微服务、区块链等这些领域或者行业的应用和实施。


### 测试驱动开发
* 重要
doc
```text
https://chenfh5.github.io/2019/09/04/00_go-test.htm
```
Lib
```text
单元测试
https://geektutu.com/post/quick-go-test.htm
https://github.com/stretchr/testify

golang interface 测试
http://goconvey.co/

静态代码分析
govet 、gofix 、gofmt、Golint
https://github.com/360EntSecGroup-Skylar/goreporter 

mock
https://github.com/golang/mock
```

### Golang APP性能监控
```text
https://github.com/divan/expvarmon
```


### Go 开发者路线图
```text
https://github.com/Quorafind/golang-developer-roadmap-cn/blob/master/ReadMe.md
```
### 编发编程模型



### 


### 消息队列
url https://nsq.io/
```text
NSQ
A realtime distributed messaging platform
```

### 网络模型
```text
新型网络库设计
netpoll 核心是 Reactor 事件监听调度器
```

### 编译原理
```text
https://godbolt.org/z/s9sWc4
```

### 内核调试工具
```text
strace
ptrace
```

### 编译方式


### 内存分配
- TCMalloc
- 回收机制

### 参考文献
程序在内存中的分布 https://www.cnblogs.com/Lynn-Zhang/p/5449199.html
从内存分配开始 https://mp.weixin.qq.com/s/EyWKFRu1xryoHY386QUcuA
译文：Go 内存分配器可视化指南 https://www.linuxzen.com/go-memory-allocator-visual-guide.html
图解Go语言内存分配 https://juejin.im/post/5c888a79e51d456ed11955a8
Golang源码探索(三) GC的实现原理 https://www.cnblogs.com/zkweb/p/7880099.html
简单易懂的 Go 内存分配原理解读 https://yq.aliyun.com/articles/652551
雨痕<<Go源码解析>>
go内存分配(英文) https://andrestc.com/post/go-memory-allocation-pt1/
https://xie.infoq.cn/article/4abf9bf90d431b717726ea1ab
https://qcon.infoq.cn/2019/guangzhou/schedule

Go 学习笔记
> https://github.com/qyuhen/book/blob/master/Go%201.5%20%E6%BA%90%E7%A0%81%E5%89%96%E6%9E%90%20%EF%BC%88%E4%B9%A6%E7%AD%BE%E7%89%88%EF%BC%89.pdf


Go 在线教程
> https://goplay.tools/

Go语言开源项目
> https://www.zhihu.com/question/20801814

Go 设计模式
> https://golangbyexample.com/all-design-patterns-golang/

Go blog 翻译
https://learnku.com/docs/go-blog/ismmkeynote/6499

Go 应用场景
https://my.oschina.net/editorial-story/blog/967845
https://tonybai.com/2017/04/20/go-coding-in-go-way/

Go 语言编程规范
https://www.infoq.cn/article/g6c95vyu5telnxxcc9yo


### 指针的用法
推荐在方法上使用指针（前提是这个类型不是 map、slice 等引用类型）
当结构体较大的时候使用指针会更高效，可以避免内存拷贝，“结构较大” 到底多大才算大可能需要自己或团队衡量，如超过 5 个字段或者根据结构体内存占用来计算
如果要修改结构体内部的数据或状态必须使用指针
如果方法的receiver是map、slice 、channel等引用类型不要使用指针
小数据类型如 bool、int 等没必要使用指针传递
如果该函数会修改receiver或变量等，使用指针

### golang 
日期
https://golangtc.com/t/572af85ab09ecc050a000063

### git go 项目 总结
https://github.com/hwholiday/learning_tools


### 幂等
https://juejin.cn/post/6906290538761158670

### 磊磊落落博客
https://leileiluoluo.com/

### https
https://caddyserver.com/


### golang高级特性
- 高阶函数
- 匿名函数

### golang原理方面的资料
https://draveness.me/golang/docs/part1-prerequisite/ch01-prepare/golang-debug/

### golang 系列文档
https://www.liwenzhou.com/posts/Go/about_golang/

### golang 官方blog 翻译
https://learnku.com/docs/go-blog/gos-declaration-syntax/6587

### 鸟窝文档
https://colobu.com/2016/04/14/Golang-Channels/

### 李周文
https://www.liwenzhou.com/archives/


### 面试、底层内容数据
http://www.topgoer.com/


### gin validator 参数校验
https://juejin.cn/post/6863765115456454664


### go并发实战git代码
https://github.com/ltqthu/gopcp.v2


### go defer 总结
https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html

### testing 单元测试、基准测试、并发基准测试
https://techlog.cn/note


### 字符串处理
https://www.gwalker.cn/article-b3f647bge5815e1dce2e6bf31027e8de.html

### 线程池处理
https://github.com/panjf2000/ants/blob/master/README_ZH.md


### upx 二进制压缩
https://abelsu7.top/2019/10/24/go-build-compress-using-upx/

### http/net 简单流程
https://www.kancloud.cn/digest/batu-go/153528


### golang 查找Cpu在做什么
### golang 查找Cpu开销

https://juejin.cn/post/6844903887757901831

https://juejin.cn/post/6844903680232144910

```text
当然，go tool trace不能解决一切问题。 如果您想跟踪运行缓慢的函数，或者找到大部分CPU时间花费在哪里，这个工具就是不合适的。 为此，您应该使用go tool pprof，它可以显示在每个函数中花费的CPU时间的百分比。 go tool trace更适合于找出程序在一段时间内正在做什么，而不是总体上的开销。 此外，还有“view trace”链接提供的其他可视化功能，这些对于诊断争用问题特别有用。 了解您的程序在理论上的表现（使用老式Big-O分析）也是无可替代的。
```

## https://jingwei.link/tags#k8s
不错的博客 里面有喝多golang知识



## TDD 
https://tech.youzan.com/you-zan-go-xiang-mu-dan-ce-ji-cheng-zeng-liang-fu-gai-lu-tong-ji-yu-fen-xi/

单元测试
基准测试
覆盖测试
静态代码分析


### json 序列化技巧
https://www.jianshu.com/p/f2e6c8c4bb90


#### 数据结构
https://github.com/emirpasic/gods#arraylist



#### golang 调试工具
```text
dlv.exe     go语言调试工具
gocode.exe  go语言代码检查，自动补全
godef.exe   go语言代码定义和引用的跳转
golint.exe  go语言代码规范检查
go-outline.exe 用于在Go源文件中提取JSON形式声明的简单工具
gopkgs.exe  快速列出可用包的工具
gorename.exe 在Go源代码中执行标识符的精确类型安全重命名
goreturns.exe 类似fmt和import的工具，使用零值填充Go返回语句以匹配func返回类型
go-symbols.exe 从go源码树中提取JSON形式的包符号的工具
gotour.exe  go语言指南网页版
guru.exe    go语言源代码有关工具，如代码高亮等
```

#### golang 面试
```text
https://www.cnblogs.com/wpgraceii/p/10528183.html
https://www.cnblogs.com/Survivalist/p/11527949.html
```