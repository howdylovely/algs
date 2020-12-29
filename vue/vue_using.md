<!--
 * @Descripttion: 
 * @version: 
 * @Author: WangShuaibing
 * @Date: 2020-12-15 11:08:15
 * @LastEditors: WangShuaibing
 * @LastEditTime: 2020-12-16 10:59:59
-->
## vue 属性解释



### data用法
> 定义模板返回的数据
```text
Vue 在data中定义的数据，其在dom中访问可以只用数据名，但是在method中访问必须前面加this. 不然提示not defined；
但是在组件中，因为可能在多处调用同一组件，所以为了不让多处的组件共享同一data对象，只能返回函数。

created:html加载完成之前，执行。执行顺序:父组件-子组件
mounted:html加载完成后执行。执行顺序:子组件-父组件
methods:事件方法执行。
watch:去监听一个值的变化，然后执行相对应的函数。
computed:computed是计算属性，也就是依赖其它的属性计算所得出最后的值

```

- data在vue实例中
```javascript
new Vue( {
    data : {
        title : 'abc'
    }
});
```

- data在组件中
```javascript
Vue.component( 'component-name', Vue.extend( {
    data : function() {
        return {
            title : 'abc'
        };
    }
}));

es写法
<script type="text/javascript">
    var app=new Vue({
        el:'#app',
        data() {
            return {
                isLogin: false
            }
        }
    })
</script>
```

## methods 用法
参考地址
https://www.cnblogs.com/zccblog/p/7206611.html

native  给组件绑定构造器里的原生事件。
```text

```


### refs ref 的用法
> https://blog.csdn.net/weixin_41646716/article/details/80455506

### vue 混入的用法


### slot 的用法


### 泛型用法
https://liyang0207.github.io/2018/01/15/%E3%80%8ATypeScript%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0-%E6%B3%9B%E5%9E%8B%E3%80%8B/


### 组合是API
https://vue-composition-api-rfc.netlify.app/zh/#%E6%A6%82%E8%BF%B0


### vue 编码但不渲染的标签 vue-fragment
https://blog.csdn.net/sunyaqi_/article/details/108307167

### typescript 代码格式化
https://my.oschina.net/someok/blog/3054456
eslint 配置
https://segmentfault.com/a/1190000009077086

### 箭头用法
https://www.jianshu.com/p/5ddc93b108c9


### 前端总结
https://juejin.cn/post/6901466994478940168

### 符号总结
> https://segmentfault.com/a/1190000023943952