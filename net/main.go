package main

/**
Dial() 函数： func Dial(net, address string) (Conn, error) net: 网络协议的名字， address: IP地址或域名，而端口以：的形式跟随在地址或域名的后面
														 如果连接成功返回对象，否则返回error
TCP:conn,err := net.Dial("tcp","127.0.0.1:8080")
UDP:conn,err := net.Dial("udp","127.0.0.1:975")
ICMP：conn,err := net.Dial("ip4:icmp","www.baidu.com")
ICMP 链接：使用协议编号 conn,err := net.Dial("ip4:1","10.0.0.3") 协议编号的意义：http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xml。
目前Dial() 函数支持如下几种网络协议，tcp、tcp4（仅限IPv4）、tcp6（仅限IPv6）udp，udp4、udp6，ip、ip4、ip6。
成功建立连接后，就可以进行数据的发生和接收了。发送数据时，使用conn 的Write() 成员方法，接收数据时使用Read() 方法，
针对 不同的网络协议，Go net包定义了如下接口：
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err error)
func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err error)
func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error)
func DialUnix(net string, laddr, raddr *UnixAddr) (c *UnixConn, err error)
 */
func main() {


}
