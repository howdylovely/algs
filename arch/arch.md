<!--
 * @Descripttion: 
 * @version: 
 * @Author: WangShuaibing
 * @Date: 2020-10-24 15:44:54
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-12-18 10:25:34
-->
# 架构设计

- 云原生
- 微服务


### 微服务技术栈
- 网关实现
```text
zuul
kong
nginx+lua
```

- 服务注册、服务发现配置中心
```text
consul 
etcd
```

- 微服务框架
```text
 go-kit/kit // 工具
 Go Micro // 微服务框架
```

- 链路追踪
```text
Zipkin²和Jaeger³
```


参考
https://www.cnblogs.com/jiujuan/p/13301055.html

### 视图

- 逻辑架构
    - 功能需求
        - 提供的服务
        - 职责和行为
        - 用户可见功能
        - 辅助功能
        - 交互与协作
- 开发架构
    - 开发期质量属性
        - 系统软件
        - 中间件
        - 程序组件
        - 模块组织方式
        - 框架类库
- 运行架构
    - 运行期质量属性
        - 并发与同步
        - 对象、线程
        - 系统通讯
        - 运行时交互
        - 单元结构
- 物理架构
    - 安装和部署需求
        - 部署方式
        - 物理设备
        - 网络环境
        - 安全性
        - 其它性能
- 数据架构
    - 数据需求
        - 数据传递
        - 数据复制
        - 数据同步
        - 实体对象
        - 存储格式
参考文档
> https://blog.csdn.net/qq_21794823/article/details/103409495

### 绘图
- 组织架构
    - 人员的分工

- 系统架构
    - 项目分层
    - 分层+关系箭头

- 业务架构
    - 粒度更细，功能模块


- 项目管理
    - 质量屋
    - 甘特图
    - 日历
    - 状态图
    - 思维导图
    - 工作分解图
    - 时间线
    - PERT图


### 并行计算


### 分布式计算


### 微服务解决的问题
部署

交付

API

版本控制

合同

缩放/自动缩放

服务发现

负载均衡

路由/自适应路由

健康检查

配置

熔断器

bulk-heads

TTL / deadlining

延迟跟踪

服务因果跟踪

分布式日志

度量操作与收集


https://github.com/binhnguyennus/awesome-scalability


文档生成工具

loppo: 非常简单的静态站点生成器
idoc：简单的文档生成工具
gitbook：大名鼎鼎的文档协作工具
docsify：一个神奇的文档站点生成器，简单轻巧，无需静态构建html



ipmitool 介绍
> https://www.cnblogs.com/zhangxinglong/p/5012441.html


如何设计一个秒杀系统
> https://gongfukangee.github.io/2019/06/09/SecondsKill/#0-%E5%9F%BA%E6%9C%AC%E7%A7%92%E6%9D%80%E9%80%BB%E8%BE%91