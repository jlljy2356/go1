package util

import (
	"crypto/md5"   //这个包提供了对 MD5 算法的支持，用于进行数据的 MD5 加密。
	"encoding/hex" //这个包提供了将数据转换为十六进制编码的功能，用于将 MD5 加密结果转换为字符串形式。
)

func EncryMd5(s string) string {
	ctx := md5.New()                        //创建一个新的 MD5 上下文。
	ctx.Write([]byte(s))                    //将要加密的字符串转换成字节数组，并写入到 MD5 上下文中。
	return hex.EncodeToString(ctx.Sum(nil)) //计算 MD5 加密结果，并将结果转换成十六进制字符串返回。
}
