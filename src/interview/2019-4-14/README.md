## 1，为什么 TCP 连接叫 3 次握手，而断开却叫 4 次挥手？  

## 2，在浏览器中输入：https://baidu.com 到页面显示出来，你能描述下中间都经历了什么，涉及到什么技术，越详细越好。  

A1：首先在浏览器缓存查找baidu.com的IP地址（如果有的话），再在系统缓存查找IP地址→路由器dns缓存→isp的dns缓存→isp的dns服务器从根服务器递归搜索（一般不用查这么长）。查到IP地址后，浏览器作为客户端添加443端口后发起tcp连接（三次握手），向百度服务器发起一条HTTP get请求，中间可能经历很多代理，百度服务器收到请求后，没有问题的话就返回一条响应，客户端读取HTTP响应报文渲染在屏幕上，最后浏览器关闭连接  

A2：客户端-dns服务器（返回ip），客户端（发送请求包数据）-服务端，服务端（发送响应包数据）-客户端，涉及技术：http协议