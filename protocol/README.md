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

提供了两种基本的验证方法:
- mysql_native_password
- ssl

## 连接参数

提供了如下连接参数的设置:

- Capability 
- Com
- Statu Flag