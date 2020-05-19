package sha256

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha(arg string) string {
	h := sha256.New()    //创建sha256算法
	h.Write([]byte(arg)) //用sha256算法对参数a进行加密，得到8个变量
	hash := h.Sum(nil)   //将8个变量相加得到最终hash
	return hex.EncodeToString(hash)
}
