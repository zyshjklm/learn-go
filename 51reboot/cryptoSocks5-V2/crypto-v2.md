### crypto V2

V2版本的实现，完全是在一个经脉逆乱的过程中进行的。产生的结果也不好。一开始就走错了方面。需要对V1的文档读三遍。才能找回方向感。

——因为对解密的理解，一开始就搞错了方向。应该对远端来的conn进行Reader封装，结果我却在本地封装了一个Writer，将conn直接Copy到Writer。一错到底。实现了一个怪胎。

错误开始于：

```go
	// 对client端请求的数据进行解密
	decryptBuf := new(bytes.Buffer)
	decryptWr := mycrypto.NewCryptoWriter(decryptBuf, key)
	// 对server端返回的数据进行加密
	encryptBuf := new(bytes.Buffer)
	encryptWr := mycrypto.NewCryptoWriter(encryptBuf, key)

```

其他部分参考代码。这版代码就不动了。另实现V3版本。

