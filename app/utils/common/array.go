package common

//
// GetStringKeysArrayByMap
// @Description: 获取键值数组
// @param data
// @return []string
//
func GetStringKeysArrayByMap(data map[string]interface{}) []string {
	var keys []string
	for key, _ := range data {
		keys = append(keys, key)
	}
	return keys
}
