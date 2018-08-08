# Client的说明

client包实现了基本的mysql客户端协议,包括
- connection
- charset
- capability
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
服务端需要至少提供如下capability:
- CLIENT_PLUGIN_AUTH
- CLIENT_PROTOCOL_41
- CLIENT_SECURE_CONNECTION


## 验证方法

客户端支持如下几种连接方法

- mysql_old_password(不推荐)
- mysql_native_password
- mysql_clear_password
- sha256_password


