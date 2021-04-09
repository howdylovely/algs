<!--
 * @Descripttion: 
 * @version: 
 * @Author: WangShuaibing
 * @Date: 2020-09-27 09:26:23
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-10-22 16:52:31
-->
# Linux 操作系统

> 系统的分类


> 系统的启动过程

> 计算机网络
点到点是物理拓扑，如光纤，就必须是点到点连接，DDN专线也是，即两头各一个机器中间不能有机器。
点到点是网络层的，你传输层只认为我的数据是从a直接到e的，但实际不是这样的，打个比方，传输层好象领导，他发布命令：要干什么什么事，但真正干的不是他，真正干的是员工，也许领导认为很简单一句话就可以干好的事，在员工眼里却是难于登天，手续极其烦琐，所以传输层是发布命令的领导，他说的是命令，也就是最终的目的，所以他只看到最初的地址和最终的地址，既一个任务的两个端点，网络层就相当于员工，领导的任务我要一步一步的作完，先从a到b,在从b到c...,所以他看到的只是整个任务的一个阶段，a到b,b到c...这就是点到点。

端到端是网络连接。网络要通信，必须建立连接，不管有多远，中间有多少机器，都必须在两头（源和目的）间建立连接，一旦连接建立起来，就说已经是端到端连接了，即端到端是逻辑链路，这条路可能经过了很复杂的物理路线，但两端主机不管，只认为是有两端的连接，而且一旦通信完成，这个连接就释放了，物理线路可能又被别的应用用来建立连接了。TCP就是用来建立这种端到端连接的一个具体协议，SPX也是。
端到端是传输层的，你比如你要将数据从A传送到E，中间可能经过A->B->C->D->E,对于传输层来说他并不知道b,c,d的存在，他只认为我的报文数据是从a直接到e的，这就叫做端到端。
总之，一句话概括就是端到端是由无数的点到点实现和组成的。



> 进程中的堆栈
https://blog.csdn.net/guoping16/article/details/6579434

> ZStack 装机过程
https://www.zstack.io/docs/2.3.2/pdf2/a_print_ZStack_Baremetal_Deploy_Tutorial.pdf

> linuxli.com
http://linuxli.com/2018/08/RHEL6vsRHEL7/


> 日期设置
GMT(Greenwich Mean Time，格林威治标准时间): 是指位于英国伦敦郊区的皇家格林尼治天文台的标准时间，因为本初子午线被定义在通过那里的经线。
UTC(Universal Time/Temps Cordonné 世界标准时间)
CST(Central Standard Time ); 中国标准时间（China Standard Time）
GMT + 8 = UTC + 8 = CST




### ssh 工具
https://www.smarthomebeginner.com/best-ssh-clients-windows-putty-alternatives/