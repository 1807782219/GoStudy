package main

import (
	"crypto/md5"
	"fmt"
	"crypto/sha1"
)

/**
	安全：
（1）数据加密
- 加密 ：
1、 对称加密： 采用单密钥的加密算法，我们称为对称加密。整个系统由如下几部分构成：需要加密的明文、加密算法和密钥。在加密和解密中，使用的密钥只有一个。
常见的单密钥加密算法有DES、AES、RC4等。
2、非对称加密：采用双密钥的加密算法，我们称为非对称加密。整个系统由如下几个部分构成：需要加密的明文、加密算法、私钥和公钥。
在该系统中，私钥和公钥都可以被用作加密或者解密，但是用私钥加密的明文，必须要用对应的公钥解密，用公钥加密的明文，必须用对应的私钥解密。
常见的双密钥加密算法有RSA等。
在对称加密中，私钥不能暴露，否则在算法公开的情况下，数据等同于明文，而在非对称加密中，公钥是公开的，私钥是保密的。这样任何人都可以把自己的信息
通过公钥和算法加密，然后发送给公钥的发布方，只有公钥发布方才能解开密文
3、 求不可解密： 哈希算法是一种从任意数据中创建固定长度摘要信息的办法。一般我们要求，对于不同的数据，要求产生的摘要信息也是唯一的。
常见的哈希算法包括MD5、SHA-1等。
（2）数字签名
数字签名，是指用于标记数字文件拥有者、创造者、分发者身份的字符串。数字签名拥有标记文件身份、分发的不可抵赖性等作用。目前，常用的数字签名采用了非对称加密。
（3）数字证书
数字证书相当于公钥
（4）PKI体系
PKI，全称公钥基础设施，是使用非对称加密理论，提供数字签名、加密、数字证书等服务的体系，一般包括权威认证机构（CA）、数字证书库、密钥备份及恢复系统、证书作废系统、应
用接口（API）等。
围绕PKI体系，建立了一些权威的、公益的机构。它们提供数字证书库、密钥备份及恢复系统、证书作废系统、应用接口等具体的服务。比如，企业的数字证书，需要向认证机构申请，以
确保数字证书的安全。
 */

func main() {
	testSHA1()
}

// MD5 加密测试
func testMD5() {
	str := "hello world"
	obj := md5.New()
	obj.Write([]byte(str))

	result := obj.Sum([]byte(" "))
	fmt.Printf("%x\n\n", result)
}
// SHA-1 加密测试
func testSHA1()  {
	str := "hello world"
	obj := sha1.New()
	obj.Write([]byte(str))

	result := obj.Sum([]byte(" "))
	fmt.Printf("%x\n\n", result)
}