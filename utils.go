/*************************************************************************
> File Name: utils.go
> Author: Kee
> Mail: chinboy2012@gmail.com
> Created Time: 2019.10.25
************************************************************************/
package kiris

func Ternary(cond bool, Tval, Fval interface{}) interface{} {
	if cond {
		return Tval
	}
	return Fval
}

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}
