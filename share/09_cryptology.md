
密码学家工具箱中的 6 大工具之一：对称密码、公钥密码、单向散列函数、消息认证码、数字签名和伪随机数生成器。

# Hash 算法

Hash 算法可以讲任意长度的二进制明文串映射为较短（通常是固定长度）的二进制串（Hash 值）。

```shell
echo "明文串" | shasum -a 256
```

详见 [Hash 算法](https://github.com/fms5cmS/Note/blob/master/Base/06-DataStructure%26Algorithm/07A-Hash.md)

# 对称加密

![](https://github.com/fms5cmS/Note/blob/master/images/tsl1.2_crypto.jpg)

对称加密(Symmetric encryption)，指在加密和解密时都是用同一密钥，目前广泛使用的对称加密算法为 AES。

对称密码需要解决将解密密钥配送给接收者的密钥配送问题！

## XOR 异或运算

对称加密是基于异或运算实现的！示例：

```shell
# 其中，明文 0110，密钥 1010
0110 XOR 1010 = 1100  # 加密得到密文
1100 XOR 1010 = 0110  # 密文和密钥异或，完成解密得到明文
```

可以看出，密钥和明文的长度是一致的。如果明文的长度不同就需要不同的密钥，且相同长度的密钥也要尽量不同，防止一次密码泄漏后所有相同长度的密文都会被解密，也就是说每次通信的密钥都是一次性，而如果每次都有办法把不同长度的密钥安全的送给接收方，为什么不直接把每次的明文送到接收方？

## padding 填充

密钥的长度通常不会太长，而明文则不一定，所以加密时通常会进行分组加密(Block cipher)，将明文分成多个等长的 Block
模块，对每个模块分别加解密，且每个模块的长度都和密钥的长度一致。

而分组加密时不会每个模块都和密钥长度一致，当模块长度比密钥短时，就需要填充(padding)了。常见的填充方式有：

- 位填充，以 bit 位为单位来填充
- 字节填充，以字节为单位为填充，常用的有以下四种
    - 补零
    - 根据 ANSI X9.23 填充，**前面都填 0，最后一个字节记录前面填充了几个 0**，解密的时候需要根据最后一个字节来判断忽略多少填充的字节
    - 根据 ISO 10126 填充，**前面填随机字符，最后一个字节记录前面填充了几个字节**
    - 根据 PKCS7（RFC5652）填充，在所有需要填充的位置填充完全相同的内容，内容为"需要填充几个字节"，如需要填充 3 个字节，则填充 3 个 3
        - PKCS7 是冗余填充，即使明文可以完整的分组，也需要填充内容，此时填充的长度最大，需要填充一个分组长度

## 工作模式

分组密码的工作模式，即如何对分组后的每个模块进行加密，通常有：

- ECB(Electronic codebook，电子密码本) 模式，直接将明文分解为多个块，对每个块独立加密，加密后的每个块没有联系
    - 无法隐藏数据特征
- CBC(Cipher-block chaining，密码分组链接) 模式，每个明文块先与前一个密文块进行 XOR 运算后，再进行加密
    - 第一个明文块需要和一个提前为其生成的初始向量(Initialization Vector，IV) 进行 XOR 运算
    - 整个解密过程是串行化的
- CFB(Cipher FeedBack，密文反馈) 模式，先将前一个密文分组进行加密，再与当前明文分组进行 XOR 运算，来生成密文分组
    - 也需要一个初始向量
- OFB(Output FeedBack，输出反馈) 模式，将密码算法的前一个输出值做为当前密码算法的输入值，该输入值再与明文分组进行 XOR
  运算，得出密文分组。
- CTR(Counter，计数器) 模式，通过递增一个加密计数器以产生连续的密钥流。计数器的值先进行加密，再与明文块进行 XOR 运算，得到密文分组

## AES 了解

AES(Advanced Encrytion Standard) 加密算法，常用填充算法为 PKCS7，常用分组工作模式为 GCM(Galois/Counter Mode)。

> GCM 即 CTR + GMAC，在 CTR 加密的基础上增加 GMAC 的特性，解决了 CTR 不能对加密消息进行完整性校验的问题。
> 
> GMAC 利用伽罗华域（Galois Field，GF，有限域）乘法运算来计算消息的 MAC 值（消息认证码）。

AES 的分组长度固定为 128 位（16字节），AES 涉及 4 种操作：字节替代(SubBytes)、行移位(ShiftRows)、列混淆(MixColumns) 和轮密钥加(AddRoundKey)，加解密需要经过多轮操作，密钥长度不同，轮数不同：

- AES-128，密钥长度 16 字节，加密轮数 10
- AES-192，密钥长度 24 字节，加密轮数 12
- AES-256，密钥长度 32 字节，加密轮数 14

加密步骤：

1. 把明文按照 128bit(16 字节)拆分成若干个明文块，每个明文块是 4*4 矩阵
2. 按照选择的填充方式来填充最后一个明文块
3. 每一个明文块利用 AES 加密器和密钥，加密成密文块
4. 拼接所有的密文块，成为最终的密文结果

## 使用

[Golang 中的使用示例](./examples/digitalSignature/aesEncryption_test.go)

```shell
# -aes-128-cbc 采用 AES-128、使用 CBC 工作模式
# -in 指定输入文件 -out 指定输出文件
# -K 指定密钥
# -iv 指定初始向量
openssl enc -aes-128-cbc -in in.txt -out out.txt -K 12345678901234567890 -iv 12345678
# -d 解密
openssl enc -aes-128-cbc -in in.txt -out out.txt -K 12345678901234567890 -iv 12345678 -d

# 根据字符串 "helloworld" 和随机生成的 salt 生成密钥和初始向量，也可以用 -nosalt 不加盐
openssl enc -aes-128-cbc -in in.txt -out out.txt -pass pass:helloworld
```

# 非对称加密

非对称加密(Asymmetric encryption) 的参与方都有一个密钥对(key pair)，包含公钥、私钥，其中私钥需要保密，而公钥则是可以公开的，加密和解密时不会使用同一个密钥。两者在数学上相关，但知道一个无法推导出另一个。

一般用法：

- 私钥加密公钥解密，可以证明"私钥拥有者"的唯一身份，用于签名
  - 注：使用私钥加密过的密文，也只有公钥才能解密
- 公钥加密私钥解密，确保发送的信息只有"私钥拥有者"可以解密，防止泄漏

假设 A 要和 B 通话，两个人各有自己的密钥对，当 A 要向 B 发送消息时，先使用 B 的公钥对消息加密，然后再将密文发送给 B，而 B 收到密文后，会使用自己的私钥对密文解密。

在对称加密中存在密钥配送的问题，由于加密和解密的密钥是相同的，所以必须向接收者配送密钥。而如果使用非对称加密，则无需向接收者配送用于解密的密钥，这样就解决了密钥配送的问题。

RSA(三位开发者首字母组成)是目前广泛使用的非对称加密算法。

```shell
openssl genrsa -out private.pem # 生成私钥
openssl rsa -in private.pem -pubout -out public.pem # 从私钥中提取出公钥

openssl asn1parse -i -in private.pem # 查看 ASN.1 格式的私钥
openssl asn1parse -i -in public.pem  # 查看 ASN.1 格式的公钥
openssl asn1parse -i -in public.pem -strparse 19 # 19 是上一步拿到的 BIT STRING 最前面的数字

openssl rsautl -encrypt -in hello.txt -inkey public.pem -pubin -out hello.en # 使用公钥对文件 hello.txt 加密得到 hello.en
openssl rsautl -decrypt -in hello.en -inkey private.pem -out hello.de # 使用私钥对加密文件 hello.en 解密的搭配 hello.de
```

公钥密码解决了密钥配送的问题，但是存在通过中间人攻击被伪装的风险，因此需要对带有数字签名的公钥进行认证。

# 混合加密

混合加密机制同时结合了对称加密和非对称加密的优点。

该机制的主要过程为：

1. 先用非对称加密（计算复杂度较高）协商出一个临时的对称加密密钥（或称会话密钥）
2. 然后双方再通过对称加密算法（计算复杂度较低）对所传递的大量数据进行快速的加密处理。

典型的应用案例是网站中使用越来越普遍的通信协议：安全超文本传输协议（Hyper Text Transfer Protocol Secure，HTTPS），详见 [HTTPS 加密](#https-加密)

# 单向散列函数

消息摘要(Message Digest) 算法可以利用哈希函数生成文件的"指纹"，得到指纹后就可以利用它来验证消息（文件、图片等）的完整性。也称为单向散列函数。

消息摘要算法的哈希函数具有以下特点：

- 根据任意长度的消息计算出固定长度的散列值
- 能够快速计算出散列值
- 消息不同，散列值也不同！即哈希碰撞的概率要很低
- 单向性，即无法根据哈希值反推出原始消息

常见消息摘要算法：

|           算法           |                                                    比特位                                                    |  现状   |
|:----------------------:|:---------------------------------------------------------------------------------------------------------:|:-----:|
| MD4 (Message Digest 4) |                                                   128 位                                                   | 已不再安全 |
| MD5 (Message Digest 5) |                                                   128 位                                                   | 已不再安全 |
|          SHA           |                                                   160 位                                                   | 已不再安全 |
|         SHA-1          |                                                   160 位                                                   | 已不再安全 |
|       RIPEMD-160       |                                                   160 位                                                   | 目前安全  |
|         SHA-2          |        SHA-224 224 位、SHA-256 256 位、SHA-384 384 位、SHA-512 512 位、SHA-512/224 224 位、SHA-512/256 256 位        | 目前安全  |
|         SHA-3          | SHA3-224 224 位、SHA3-256 256 位、SHA3-384 384 位、SHA3-512 512 位、SHAKE128 d(arbitrary)位、SHAKE256 d(arbitrary)位 | 目前安全  |

消息摘要算法能够辨别出“篡改”，但是无法辨别出“伪装”。对于一条消息不仅需要确认消息的完整性，还需要确认这个消息来自谁。而这仅靠完整性检查是不够的，还需要进行消息认证。认证的技术包括消息认证码和数字签名。

其中，数字签名不仅能够向通信对象保证消息没有篡改，还能够向所有第三方做出这样的保证。

# 消息认证

消息认证由于可以完成认证，所以不会存在中间人伪装。

## 消息认证码

MAC(Message Authentication Code，消息认证码) 是一种**确认完整性并进行认证**的技术，相比于单向散列函数直接根据消息生成哈希值，MAC 则是根据消息、共享密钥来生成 MAC 值的。

传输时将密文和 MAC 值同时传输给接收方，接收方收到后先通过同样的 MAC 算法计算收到后密文的 MAC 值，然后与接受到的 MAC 对比，如果一致再做解密。

> SSL/TLS 对通信内容的认证和完整性校验使用了消息认证码。

GCM(Galois/Counter Mode) 是一种认证加密方式的实现。GCM 中使用分组加密中的 CTR 工作模式，并使用一个反复进行加法和乘法运算的散列函数来计算 MAC 值。专门用于消息认证码的 GCM 成为 GMAC。

最常用的是 HMAC(Hash-based Message Authentication Code) 基于哈希的消息认证码算法，它使用单向散列函数来构造消息认证码，任何高强度的单向散列函数都可以用于 HMAC 中。

基本过程为对某个消息，利用**提前共享的对称密钥和 Hash 算法**进行处理，得到 HMAC 值。该 HMAC 值持有方可以向对方证明自己拥有某个对称密钥，并且确保所传输消息内容未被篡改。

消息认证码可以用于简单证明身份的场景，但无法"对第三方认证"、"防止抵赖"。解决这两个问题都需要用到数字签名！

> 因为消息认证码中用到的密钥是共享密钥，通信双方都有这个密钥，所以对第三方无法证明消息到底出自双方中的哪一方。
> 
> 因为消息认证码的共享密钥双方都有，无法判断消息是发自于哪一方。所以消息认证码无法防止抵赖(NonRepudiation)。

## 数字签名

消息认证码由于共享密钥的问题，导致无法"对第三方认证"、"防止抵赖"。而数字签名通过让通信双方的共享密钥不同来解决这些问题，因为从密钥上能区分出谁是谁。

场景：Alice 通过信道发给 Bob 一个文件（一份信息），Bob 如何获知所收到的文件即为 Alice 发出的原始版本？

1. Alice 可以先对文件内容进行摘要，然后用自己的私钥对摘要进行加密（签名），之后同时将文件和签名都发给 Bob；
2. Bob 收到文件和签名后，用 Alice 的公钥来解密签名，得到数字摘要，与对文件进行摘要后的结果进行比对。如果一致，说明该文件确实是 Alice 发过来的（因为别人无法拥有 Alice 的私钥），并且文件内容没有被修改过（摘要结果一致）。

数字签名可以识别篡改、伪装、防止抵赖

数字签名中有两种行为：

- 生成消息签名 sign，由消息发送者完成，根据消息内容计算数字签名的值
- 验证消息签名 verify，第三方进行验证，验证消息的来源是否属于发送者

在非对称加密中，发送者使用公钥进行加密，接受者使用私钥完成解密；而在数字签名中，**签名者使用私钥进行加密（生成签名），验签者使用公钥完成解密（验证签名）**！！

加签过程中，通常是**对消息的散列值**进行签名，而不是直接对消息签名！因为前者在加密时会快很多。

实践中常用算法包括 DSA(Digital Signature Algorithm，基于 ElGamal 算法) 和安全强度更高的 ECDSA(Elliptic Curve Digital Signature Algorithm，基于椭圆曲线算法) 等。

Q：为什么加密后的密文能够具备签名的意义？

利用私钥加密并非是为了保证机密性，而是用于表示，只有持有该私钥的人才能生成这样的密文。所以私钥产生的密文是为了认证。

所以，数字签名并不保证消息的机密性，如果需要保证，可以考虑加密和数字签名结合起来使用。

# 随机数

用随机数的目的是为了提高密文的不可预测性，让攻击者无法一眼看穿。

- 生成密钥，用于对称密码和消息认证码
- 生成公钥密码，用于生成公钥密码和数字签名
- 生成初始化向量 IV，用于分组密码中的 CBC、CFB、OFB 模式
- 生成 nonce，用于防御重放攻击和分组密码中的 CTR 模式 
- 生成盐，用于基于口令密码的 PBE 等

# HTTPS 加密

如果你在公司内网里做开发，并且写的代码也只对内网提供服务。那么大概率你的服务是用的 HTTP 协议。

但如果哪天想让外网的朋友们也体验下你的服务功能，就需要将服务暴露到外网，这时为了防止信息被抓包，就需要将传输的明文变成密文，所以就需要在 HTTP 层再加一层 TLS 层，来做加密。

TLS 分为 1.2 和 1.3 版本，目前主流使用 1.2。

- HTTPS 的握手过程

1. 建立 TCP 连接
2. HTTPS 加密流程
   1. TLS 四次握手，利用非对称加密的特性交换信息，最后得到一个"会话密钥"
   2. 在获得"会话密钥"的基础上，进行对称加密通信

TLS 四次握手：

1. Client 告诉 Server 它支持什么样的加密协议版本，如 TLS1.2；使用什么样的加密套件，如 RSA；给出一个客户端随机数
2. Server 告诉 Client，服务器随机数 + 服务器证书 + 确定的加密协议版本
3. 这里有三步：
   1. Client 又生成一个名为 pre_master_key 的随机数，并从上一步拿到的服务器证书中取出 Server 发出的服务器公钥，利用公钥加密 pre_master_key 后发送给 Server
   2. Client 会利用目前已有的三个随机数（客户端随机数、服务器随机数、pre_master_key）计算得到一个"会话密钥"，Client 通知 Server，后续使用这个"会话密钥"进行对称加密通信
   3. Client 把迄今为止的通信数据内容生成一个摘要，利用"会话密钥"加密，发送给 Server 做校验
4. 这里有两步
   1. Server 拿到 Client 的 pre_master_key（解密得到的）后 Server 也和 Client 一样有了三个一样随机数，然后利用这三个随机数使用同样的算法得到一个"会话密钥"，Server 告诉 Client 之后利用这个"会话密钥"进行对称加密通信
   2. Server 把迄今为止的通信数据内容生成一个摘要，利用"会话密钥"加密，发送给 Client 做校验

![](../images/HTTPS.png)

- 服务器证书是什么？

服务器证书本质上是将服务器公钥利用权威数字证书机构(CA) 的私钥加密过的密文。

所以在 TLS 四次握手过程中的第三步，Client 可以通过 CA 的公钥来解密服务器证书，从而拿到服务器公钥。

- 为什么需要拿 CA 的私钥加密一次再传过去，而不能直接传服务器公钥？

因为需要一个方法来证明客户端拿到的公钥是真正的服务器公钥。

如果只传公钥，公钥就有可能会在传输的过程中就被黑客替换掉；TLS 握手第三步时客户端会拿着假公钥来加密第三个随机数 pre_master_key，黑客解密后自然就知道了最为关键的 pre_master_key，又因为第一和第二个随机数是公开的，因此就可以计算出"会话秘钥"。

- 怎么获得 CA 的公钥？

CA 的公钥是作为配置放到操作系统或浏览器中的。

- 为什么需要三个随机数？

三个随机数中 pre_master_key 是关键，但是只有一个随机数的话随机性不够，多次随机的情况下有可能出来的密钥是一样的，所以其他两个随机数是为了增加"会话密钥"的随机程度，从而保证每次 HTTPS 通信使用的"会话密钥"不同。

- 为什么 TSL 握手的第三、四步还需要提供摘要？

第三步，客户端生成摘要，服务端验证，如果验证通过，说明客户端生成的数据没被篡改过，服务端后面才能放心跟客户端通信。

第四步，则是反过来，由服务端生成摘要，客户端来验证，验证通过了，说明服务端是可信任的。

> 直接使用原文对比的话，原文内容可能会很长，所以需要使用摘要，传输成本也会更小

- 整个 HTTPS 过程中涉及几对公私钥？

两对，服务器的公私钥、CA 的公私钥。

# 同态加密

同态加密（Homomorphic Encryption）是一种特殊的加密方法：对密文直接进行函数处理后的结果 = 对明文进行函数处理后再对处理结果加密的结果，即：

![同态加密](../images/HomomorphicEncryption.png)

同态加密可以保证实现处理者无法访问到数据自身的信息。

同态性来自代数领域，一般包括四种类型：加法同态、乘法同态、减法同态和除法同态。

仅满足加法同态的算法包括 Paillier 和 Benaloh 算法；仅满足乘法同态的算法包括 RSA 和 ElGamal 算法。

零知识证明（Zero Knowledge Proof）的数学基础就是同态加密。

> 零知识证明：证明者在不向验证者提供任何额外信息的前提下，使验证者相信某个论断（Statement）是正确的。


