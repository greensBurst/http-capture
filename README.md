# http-capture
HTTP抓包工具，运行在命令行

1. cap -g -u https://www.baidu.com

   -g：get请求
   -u：url，连接必须加上http/https

     ![get](C:\Users\greens\Desktop\get.png)

   这里添加的参数等同于 https://www.baidu.com?a=1&b2

2. cap -p -u https://www.baidu.com

   -p：post请求

    ![post](C:\Users\greens\Desktop\post.png)

​	post的添加参数和消息头与get请求一样

​	 发起get/post请求后会显示request和response的请求行和请求头

![result](C:\Users\greens\Desktop\result.png)

​	支持通过html标签筛选响应正文

 ![label](C:\Users\greens\Desktop\label.png)