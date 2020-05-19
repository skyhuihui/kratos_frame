package GoogleVerify

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type GoogleAuth struct {
}

func NewGoogleAuth() *GoogleAuth {
	return &GoogleAuth{}
}
func (this *GoogleAuth) un() int64 {
	return time.Now().UnixNano() / 1000 / 30
}

func (this *GoogleAuth) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func (this *GoogleAuth) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (this *GoogleAuth) base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func (this *GoogleAuth) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (this *GoogleAuth) toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func (this *GoogleAuth) oneTimePassword(key []byte, data []byte) uint32 {
	hash := this.hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := this.toUint32(hashParts)
	return number % 1000000
}

// 获取秘钥
func (this *GoogleAuth) GetSecret() string {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, this.un())
	return strings.ToUpper(this.base32encode(this.hmacSha1(buf.Bytes(), nil)))
}

// 获取动态码
func (this *GoogleAuth) GetCode(secret string) (string, error) {
	secretUpper := strings.ToUpper(secret)
	secretKey, err := this.base32decode(secretUpper)
	if err != nil {
		return "", err
	}
	number := this.oneTimePassword(secretKey, this.toBytes(time.Now().Unix()/30))
	return fmt.Sprintf("%06d", number), nil
}

// 获取动态码二维码内容
func (this *GoogleAuth) GetQrcode(user, secret string) string {
	return fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
}

// 获取动态码二维码图片地址,这里是第三方二维码api
func (this *GoogleAuth) GetQrcodeUrl(user, secret string) string {
	qrcode := this.GetQrcode(user, secret)
	width := "200"
	height := "200"
	data := url.Values{}
	data.Set("data", qrcode)
	return "https://api.qrserver.com/v1/create-qr-code/?" + data.Encode() + "&size=" + width + "x" + height + "&ecc=M"
}

// 验证动态码
func (this *GoogleAuth) VerifyCode(secret, code string) (bool, error) {
	_code, err := this.GetCode(secret)
	if err != nil {
		return false, err
	}
	return _code == code, nil
}

var err error

func main() {
	//
	fmt.Println("-----------------开启二次认证----------------------")
	user := "hz_ex"
	goler := InitAuth(user)
	fmt.Println(goler)

	fmt.Println("-----------------信息校验----------------------")

	// secret最好持久化保存在
	// 验证,动态码(从谷歌验证器获取或者freeotp获取)
	//time.Sleep(30 * time.Second)
	bool, err := NewGoogleAuth().VerifyCode("LHDORZFO6PKSVA4BOKUNXS7XEVNXFR3P", "663821")
	if bool {
		fmt.Println("√")
	} else {
		fmt.Println("X", err)
	}
}

// 开启二次认证
func InitAuth(user string) GoogleVerify {
	// 秘钥
	secret := NewGoogleAuth().GetSecret()
	//secret := "YE2BQMNZ557KZY6GC4CCLR6ZSDDJVLBC"

	// 动态码(每隔30s会动态生成一个6位数的数字)
	code, _ := NewGoogleAuth().GetCode(secret)

	// 用户名
	_ = NewGoogleAuth().GetQrcode(user, code)

	// 打印二维码地址
	qrCodeUrl := NewGoogleAuth().GetQrcodeUrl(user, secret)
	gogle := GoogleVerify{
		Secret:    secret,
		Code:      code,
		QrCodeUrl: qrCodeUrl,
		User:      user,
	}
	return gogle
}

//生成谷歌验证
type GoogleVerify struct {
	User      string //用户名
	Secret    string //用户密钥
	QrCodeUrl string // 打印二维码地址
	Code      string //动态码
}
