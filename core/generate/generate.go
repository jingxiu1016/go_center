package generate

// 这个包用作AppKey与加密密钥之间的解析

type Generate interface {
	//	加密
	Encrypt(string)
	//	解密
	Verify(string) bool
}

//func GenerateAppKey()
