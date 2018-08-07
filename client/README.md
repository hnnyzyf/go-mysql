# Client的说明

client包实现了基本的mysql客户端协议,包括
- connection
- charset
- capacility
- compress
- authentication

## 连接建立

提供了两种连接建立的方法:
1. Plain Handshake
    
    使用mysql_native_password作为连接建立的验证方法

2. SSL Handshake

    使用SSL作为连接建立的验证方法


## 字符集

提供基本字符集的选择使用,具体内容参考如下链接

[Charset](/util/charset/READEME.md)

## 连接属性

### 服务端要求
- 发送V10版本的信息
- 支持 CLIENT_PROTOCOL_41


## 验证方法

客户端仅支持两种验证方法,不支持其他类型的验证方法,在未指明验证方法的前提下，默认使用mysql_native_password作为验证方法.
支持的验证方法如下:
- mysql_native_password
- SSL

### mysql_native_password

1. 服务端需要提供如下至少如下capability
    -  CLIENT_PROTOCOL_41 
    -  CLIENT_SECURE_CONNECTION

2. 客户端需要提供至少如下capability
     -  CLIENT_PROTOCOL_41 
    -  CLIENT_SECURE_CONNECTION

3. 客户端不能提供的capability
    - CLIENT_SSL

### SSL

1. 服务端需要提供如下至少如下capability
    -  CLIENT_PROTOCOL_41 
    -  CLIENT_SECURE_CONNECTION
    - CLIENT_SSL

2. 客户端需要提供至少如下capability
     -  CLIENT_PROTOCOL_41 
    -  CLIENT_SECURE_CONNECTION
    - CLIENT_SSL