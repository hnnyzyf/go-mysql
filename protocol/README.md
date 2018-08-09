# Mysql 协议说明

Mysql 协议主要实现了通信的基本操作和信息，包括如下内容:
- 读取packet
- 写入packet
- 验证方法
- 连接参数

## Packet的处理

### ReadPacket方法

ReadPacket主要实现了读取packet的方法，函数原型为:

```
func ReadPacket(r io.Reader, id uint8) (*Packet, error)

```

### WritePacket方法

WirtePacket主要实现了写入packet的方法，会自动根据pyaload的长度确定是否要拆分成更小的packet,函数原型为:

```
func WritePacket(w io.Writer, packet *Packet) (uint8, error)

```

## 验证方法

提供了多种验证方法:
- mysql_old_password(不推荐)
- mysql_native_password
- mysql_clear_password
- sha256_password
