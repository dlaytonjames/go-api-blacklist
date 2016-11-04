package util

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"time"
)

func CheckRateLimit(ip, request, action string) bool {
	current := int(time.Now().Unix())
	currentStr := strconv.Itoa(current)
	//limit  100次
	//timeset 600秒
	//限制600秒最多访问100次
	limit, timeset := GetRateLimitConfig()
	allowanceStr, timestampStr := LoadAllowance(ip, request, action)
	allowance, _ := strconv.Atoi(allowanceStr)
	timestamp, _ := strconv.Atoi(timestampStr)
	allowance += int(current-timestamp) * limit / timeset
	if allowance > limit {
		allowance = limit
	}

	if allowance < 1 {
		SaveAllowance(ip, request, action, "0", currentStr)
		//返回true 代表速率超过,进行错误输出
		return true
	} else {
		allowanceStr = strconv.Itoa(allowance - 1)
		SaveAllowance(ip, request, action, allowanceStr, currentStr)
		//返回false 代表速率未超过
		return false
	}
}

func LoadAllowance(ip, request, action string) (allowance, timestamp string) {
	rs, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "key":"YK_OAUTH_APP"}`)
	res, _ := (redis.String(rs.Get(ip+"_"+request), err))
	if len(res) == 0 {
		currentStr := string(time.Now().Unix())
		defaultLimitInt, _ := GetRateLimitConfig()
		defaultLimitStr := strconv.Itoa(defaultLimitInt)
		allowance, timestamp = defaultLimitStr, currentStr
	} else {
		kv := strings.Split(res, "-")
		allowance, timestamp = kv[0], kv[1]
	}
	return

}

func GetRateLimitConfig() (limit, timeset int) {
	limit = 100
	timeset = 600
	return
}

func SaveAllowance(ip, request, action, allowance, current string) {
	rs, _ := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "key":"YK_OAUTH_APP"}`)
	rs.Put(ip+"_"+request, allowance+"-"+current, 600*time.Second)
}
