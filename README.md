# http-capture
HTTP抓包工具，运行在命令行

1. cap -g -u https://www.baidu.com

   -g：get请求
   -u：url，连接必须加上http/https

     ![get](https://raw.githubusercontent.com/greensburst/http-capture/master/photo/get.png)

   这里添加的参数等同于 https://www.baidu.com?a=1&b2

2. cap -p -u https://www.baidu.com

   -p：post请求

    ![post](https://github.com/greensburst/http-capture/blob/master/photo/post.png?raw=true)

​post的添加参数和消息头与get请求一样

​发起get/post请求后会显示request和response的请求行和请求头

![result](https://github.com/greensburst/http-capture/blob/master/photo/result.png?raw=true)

​支持通过html标签筛选响应正文

![label](https://github.com/greensburst/http-capture/blob/master/photo/label.png?raw=true)