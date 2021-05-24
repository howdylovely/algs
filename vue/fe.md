<!--
 * @Descripttion: 
 * @version: 
 * @Author: WangShuaibing
 * @Date: 2020-11-19 20:00:12
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-12-18 11:36:45
-->
## FE


### Install
 https://www.cnblogs.com/goldlong/p/8027997.html

### 前端规范
https://lq782655835.github.io/blogs/team-standard/recommend-vue-project-structure.html

### vue cms
https://gitee.com/guashi/vue_cms

### 淘宝NPM相关资源
https://developer.aliyun.com/mirror/npm/


### 后台框架
https://panjiachen.github.io/vue-element-admin-site/zh/guide/#%E5%8A%9F%E8%83%BD


### Element Vue 2.0 的桌面端组件库
https://element.eleme.cn/


### 浏览器调试
Vue.js devtools

### 数据验证
https://github.com/logaretm/vee-validate

### VsCode代码格式化
待处理

### 状态码
https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502


### vue package.json介绍
https://www.cnblogs.com/wangyingblock/p/11958702.html

### npm command
```bash
npm run <command> [-- <args>
```


### NPM 钩子
https://www.ruanyifeng.com/blog/2016/10/npm_scripts.html


### vue CLI 介绍
https://cli.vuejs.org/zh/guide/cli-service.html#%E4%BD%BF%E7%94%A8%E5%91%BD%E4%BB%A4


### node npm 包shell-executor
```text
# execute the npm commands lint, test and watch in parallel
shell-exec 'npm run lint' 'npm run test' 'npm run watch'
```
https://developer.aliyun.com/mirror/npm/package/shell-executor

### Commonjs的模块规范
https://javascript.ruanyifeng.com/nodejs/module.html

### sass CSS扩展语法
 https://www.sass.hk/docs/


### vue route 语法

### vue component 语法


### 项目扩展
Babel/TypeScript 转译、ESLint 集成、单元测试和 end-to-end 测试


### webpack
https://webpack.js.org/


### npm 
https://www.jianshu.com/p/30ef3c874c1e

### npm install 区别
https://www.jianshu.com/p/920c1a6e999c



### story 组件管理


### vue compnent 介绍
https://segmentfault.com/a/1190000016409329



### yarn
https://yarn.bootcss.com/


### vscode 使用vue 安装插件
npm install -g jshint


### vue 风格指南
https://cn.vuejs.org/v2/style-guide/index.html#%E8%A7%84%E5%88%99%E5%BD%92

### element-vue 文档
https://panjiachen.github.io/vue-element-admin-site/zh/guide/advanced/sass.html#%E5%8D%87%E7%BA%A7%E6%96%B9%E6%A1%88



### webpack 入门指南
https://github.com/wallstreetcn/webpack-and-spa-guide


### vscode vue
https://github.com/varHarrie/varharrie.github.io/issues/10


### async wait promise 实例已跑
> https://www.cnblogs.com/williamjie/p/9789212.html


### vuedraggable 拖拽
http://www.ptbird.cn/vue-draggable-dragging.html


### vue3 生命周期
https://www.jianshu.com/p/fb8ba5b2b474

scm-api
```bash
docker build -t hub.ark.jd.com/scm/scm-api-test:v0.63.49 .
docker push hub.ark.jd.com/scm/scm-api-test:v0.63.49
```

scm-console
```bash
docker build -t hub.ark.jd.com/scm-web/console:1.0.1217c .
docker push hub.ark.jd.com/scm-web/console:1.0.1208a

常用命令
npm cache clean -f
rm -rf ./node_modules
npm i

win10 编译
npm run servelocal --ip=scmtest-console.jdcloud.com --rewrite=/api/v1
```


windows
网络失败重复安装
```txt
- python2.7
- cmd
node-sass
npm insatll
```


linux 
```text
yarn 安装
wget https://dl.yarnpkg.com/rpm/yarn.repo -O /etc/yum.repos.d/yarn.repo
yum -y install yarn
```



docker 

```text
# stage 0
FROM node:12-alpine as build-stage
WORKDIR /app
COPY package*.json ./
ENV SASS_BINARY_SITE https://npm.taobao.org/mirrors/node-sass
RUN npm i node-sass --sass_binary_site=https://npm.taobao.org/mirrors/node-sass
#ARG NPMREGISTRY='https://registry.npm.taobao.org'
#RUN npm install --registry=$NPMREGISTRY
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    # && npm config set @jd:registry http://registry.m.jd.com \
    && npm config set registry https://registry.npm.taobao.org \
    && npm install 
COPY . .
RUN npm run build-prod --silent

# stage 1 (nginx)
#FROM nginx:1.18-alpine
#COPY config/nginx.conf /etc/nginx/conf.d/default.conf
#COPY --from=build-stage /app/dist /usr/share/nginx/html
FROM hub.ark.jd.com/skywing/nginx:latest
ENV PORT 8081
COPY config/nginx-skywing.conf /export/servers/nginx/conf/nginx.conf
COPY --from=build-stage /app/dist /export/servers/nginx/html 
EXPOSE 8081
```



vue run env 
```text
        // "servelocal": "shell-exec --colored-output \"npm run serve --local\" \"npm run express\"",
```
#### vue elment table 拖动排序
https://blog.csdn.net/weixin_46527264/article/details/111693851