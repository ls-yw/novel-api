package common

//
// MapMergeString
// @Description: 合并map
// @param maps
// @return map[string]interface{}
//
func MapMergeString(maps ...map[string]interface{}) map[string]interface{} {
	var data map[string]interface{}
	for i := 0; i < len(maps); i++ {
		for key, value := range maps[i] {
			data[key] = value
		}
	}
	return data
}
