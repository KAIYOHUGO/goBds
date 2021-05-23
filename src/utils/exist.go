package utils

// only check non init value
// if the value can be 0,"" , don't use this func
func IsExist(v ...interface{}) bool {
	for _, el := range v {
		switch t := el.(type) {
		case string:
			if t == "" {
				return false
			}
		case int:
		case int8:
		case int16:
		case int32:
		case int64:
		case uint:
		case uint8:
		case uint16:
		case uint32:
		case uint64:
		case float32:
		case float64:
			if t == 0 {
				return false
			}
		}
	}
	return true
}
