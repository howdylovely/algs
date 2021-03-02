## 自动化装机

系统运维：
- 使用 Django + Angular 开发基于 iPXE 和 kickstart 的装机平台：
- 适配 DeLL / HP / 超微 / 联想 / 华为 / 曙光 / H3C 等主流厂商的机型
- 装机实现自动配置 RAID 、IPMI 以及部分厂商的 BIOS
- 制作 LiveOS 系统通过 iPXE 引导后采集主机信息，以及服务器的故障排查

CMDB 相关功能扩展开发：
- 使用 Flask + ipmitool 开发通过 BMC 硬重启物理机的 API
- 开发终端命令行工具查询 CMDB 服务器信息
- 使用阿里云最新版 SDK 实现对公网 DNS 记录的增删改查操作

服务器日常的故障诊断及维修：
- 升级 docker 存储引擎为 overlay 解决 devicemapper 存储出现的 IO 异常问题
- 通过升级 BIOS 固件和系统内核，解决 nvidia 显卡机型频繁重启的问题
- 批量更新联想 BMC 固件，解决机器频繁出现的 BMC 死机问题
- 服务器采购到货后的 IDC 上架规划，以及装机交付前的所有相关工作
- 维护现有的 VMware 虚拟化平台的虚拟机相关的资源管理及日常维护

Developed iPXE and kickstart based clone platform by Django and Angular :

- Compatible with DeLL, HP, Supermicro, Lenovo, Huawei, H3C, Sugon servers.
- Configured BIOS, IPMI and RAID automatically at pre install stage.
- Create LiveOS to collect server info and troubleshoot broken server.

Extending then CMDB functions:

- CMDB could reboot server by IPMI API developed by Flask and ipmitool.
- Developed a shell command to search CMDB info with Flask API.
- Updated DNS records with Aliyun lastest python SDK.

Maintained Linux system and rack servers:

- Updated Docker storage driver from devicemapper to OverlayFS fixed high IO issue.
- Upgrade BIOS firmware and OS kernel fixed reboot issue with Nvidia video cards.
- Upgrade Lenovo BMC firmware fixed IPMI no response when server was down.


使用 Flask + Bootstrap 开发基于 iPXE 和 kickstart 的装机平台：
- 编写 shell 脚本，装机实现自动配置 DeLL 服务器的 BIOS，IPMI，RAID

使用 ansible 批量初始化私有云环境：
- 配置 keepalived 服务，主备切换时 shell 脚本加锁，解决 VPN 路由配置异常问题
- 系统安装成功后，批量进行磁盘并行格式化挂载

负责公有云 KVM 虚拟机，以及青云虚拟机的安装部署
培训 IDC 工程师现场实施系统相关的操作，以及常见的系统故障排查
协助网络及 IDC 同事做一些自动运维相关的工具开发

Developed iPXE and kickstart-based installation platform by Flask and Bootstrap:
- Configuring BIOS, IPMI, RAID of DeLL server by bash shell script.

Use ansible to batch initialize the private cloud environment:
- Configure keepalived service
- Solved VPN routing abnormality by shell script with flock, when master, slave switch.
- After installed CentOS system, parted and formatted disks in parallel.

Responsible for the installation and deployment of public cloud KVM and Qingyun VM.
Training IDC engineers to implement system and troubleshootings on-site.
Assist network and IDC colleagues to do some automatic tool development.


系统运维：

- 维护 PXE 装机平台的 KS 模板，shell 配置脚本
- 将各个机房的 PXE 装机平台从虚拟机迁移到物理机上
- 修复超微、曙光服务器关机情况下 IPMI PXE 重启失败的问题
- 为 PXE 装机平台，增加 CentOS 7 和 ubuntu preseed 自动装机模板
- 自动完成 openstack 平台 VM 初始化扩展 LVM 逻辑卷和文件系统操作
- 特定机型 PXE 装机时，识别到 SSD 磁盘，自动创建 flashcache 设备
- 修复系统网卡 eth0，eth1 装机后默认命名为 em0，em1 的问题
- 为下一代装机系统，创建基于 centos 6 的 PXE LiveOS 镜像

运维开发：

- 使用 python + PostgreSQL 开发 IPMI 监控系统，并提供 RESTful 查询 API
- 使用 Flask + Bootstrap 扩展 CMDB 根据主机名，IP 等自定义字段批量查询的功能
- 扩展 CMDB 系统，根据不同机房和厂商，批量查询故障报修信息

IDC 运维：
- 维护办公网机房的所有 IDC 运维工作
- 为新采购或搬迁的服务器，规划所在机房的 IP 地址和机柜使用情况
- 批量修改跨机房搬迁服务器的网络和 IPMI 配置
- 故障机器的维修，以及服务器硬件设备的更换、升级

Maintaining PXE clone system backend (ks config and bash shell script) :

- fixed default network device name from em0, em1 to eth0, eth1
- fixed Supermicro server IPMI boot failed at power off state
- autocreated flashcache device with SSD disk for specific servers
- added new kickstart template for ubuntu and CentOS 7 system
- automated openstack VM disk extended LVM and FS at initial setup
- created LiveOS for Next generation clone system

DevOPS:

- developed a batch search page with IP or host in CMDB system with flask, bootstrap
- developed IPMI monitor and RESTful API with python, postgresql

IDC management:

- planing the use of rackspace in colocated datacenters
- deployed OS for rack servers
- changed servers config with ansible for migration between datacenters
- repaired broken components or add devices for rack servers

Daily maintenance and troubleshooting for Linux system

负责阿里云，淘宝新采购服务器的系统安装和资源分配
使用 libvirt 的 python API 封装新、旧 VM 装机平台的差异化操作
使用 python 封装装机系统 API 实现批量安装和分配虚拟机
使用 python 通过 snmp 批量采集 IDC 机柜的温湿度、功率等监控信息

Deploy OS for rack servers and XEN vm servers
Maintain XEN vm clone system backend with libvirt python API
Develop a temp and power monitor for rackspace by SNMP with web.py
Repair broken components or add devices for rack servers