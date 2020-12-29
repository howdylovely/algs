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