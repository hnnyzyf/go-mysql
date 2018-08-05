# Client 协议说明

说明了当前实现的客户端所支持的协议版本,连接验证方法和压缩属性


## 服务端需求

本客户端协议需要服务器端至少支持mysql_native_password的验证方法,即服务端的capabilities包含至少如下三个属性：
- CLIENT_PROTOCOL_41
- CLIENT_PLUGIN_AUTH
- CLIENT_SECURE_CONNECTION 

## 客户端说明

客户端可提供的capabilities属性包括如下内容:

- CLIENT_PROTOCOL_41
- CLIENT_PLUGIN_AUTH
- CLIENT_SECURE_CONNECTION 
