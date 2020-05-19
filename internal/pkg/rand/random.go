package rand

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

//获取几位随机数字不是0开头
func GetRandomInt(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	time.Sleep(100 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
again:
	for i := 0; i < l; i++ {
		b := bytes[r.Intn(len(bytes))]
		if i == 0 && string(b) == "0" {
			goto again
		}
		result = append(result, b)
	}
	return string(result)
}

// 获得指定位数随机码
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	time.Sleep(100 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 根据随机Hash 生成随机数
func GetNumByHash(hash string) int {
	var valid = regexp.MustCompile("[0-9]")
	validList := valid.FindAllStringSubmatch(hash, -1)
	if len(validList) == 0 {
		return 0
	}
	str := ""
	index := RandInt64(1, 5)
	for i := 1; i <= int(index); i++ {
		num := RandInt64(1, int64(len(validList)))
		str = str + validList[num][0]
	}

	num, _ := strconv.Atoi(str)
	return num
}

// 区间随机数 几到几的数字
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	time.Sleep(100 * time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

// 指定精度 生成 某个大小区间随机数
func RandFloat(min, max float64, precision int) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(precision)+"f", min+rand.Float64()*(max-min)), 64)
	return value
}

// 根据指定两个数据， 返回两个数据区间的随机两个数据
// return min, max
func RandFloatMinMax(min, max float64, precision int) (float64, float64) {
again:
	m1, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(precision)+"f", min+rand.Float64()*(max-min)*RandFloat(0, 1, 2)), 64)
	time.Sleep(10 * time.Millisecond)
	m2, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(precision)+"f", min+rand.Float64()*(max-min)*RandFloat(0, 1, 2)), 64)
	if m1 < m2 {
		return m1, m2
	} else if m1 == m2 {
		goto again
	}
	return m2, m1
}
