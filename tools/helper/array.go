package helper

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// InArray 元素是否在数组(切片/字典)内.
func (ta *TsArr) InArray(needle interface{}, arr interface{}) (r bool) {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				r = true
				return
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				r = true
				return
			}
		}
	default:
		panic("[InArray]arr type must be array, slice or map")
	}
	return
}

// ArrayFill 用给定的值value填充数组,num为插入元素的数量.
func (ta *TsArr) ArrayFill(value interface{}, num int) []interface{} {
	if num <= 0 {
		return nil
	}

	var fillArr = make([]interface{}, num)
	for i := 0; i < num; i++ {
		fillArr[i] = value
	}

	return fillArr
}

// ArrayFlip 交换数组中的键和值.
func (ta *TsArr) ArrayFlip(arr interface{}) map[interface{}]interface{} {
	flipArr := make(map[interface{}]interface{})
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			if val.Index(i).Interface() != nil && fmt.Sprintf("%v", val.Index(i).Interface()) != "" {
				flipArr[val.Index(i).Interface()] = i
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if val.MapIndex(k).Interface() != nil && fmt.Sprintf("%v", val.MapIndex(k).Interface()) != "" {
				flipArr[val.MapIndex(k).Interface()] = k
			}
		}
	default:
		panic("[ArrayFlip]arr type must be array, slice or map")
	}

	return flipArr
}

// ArrayKeys 返回数组中所有的键名.
func (ta *TsArr) ArrayKeys(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	keys := make([]interface{}, val.Len())
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			keys[i] = i
		}
	case reflect.Map:
		for i, k := range val.MapKeys() {
			keys[i] = k
		}
	default:
		panic("[arrayValuesHelper]arr type must be array, slice or map")
	}

	return keys
}

// ArrayValues 返回数组(切片/字典)中所有的值.
// filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素.
func (ta *TsArr) ArrayValues(arr interface{}, filterNil bool) []interface{} {
	return ta.arrayValuesHelper(arr, filterNil)
}

// MergeSlice 合并一个或多个数组/切片.
// filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素;ss是元素为数组/切片的数组.
func (ta *TsArr) MergeSlice(filterNil bool, arr ...interface{}) []interface{} {
	var merge []interface{}
	switch len(arr) {
	case 0:
		break
	default:
		n := 0
		for i, v := range arr {
			chkLen := isArrayOrSliceHelper(v, 3)
			if chkLen == -1 {
				panic(fmt.Sprintf("[MergeSlice]ss type must be array or slice, but [%d]th item not is.", i))
			} else {
				n += chkLen
			}
		}
		merge = make([]interface{}, 0, n)
		var item interface{}
		for _, v := range arr {
			val := reflect.ValueOf(v)
			switch val.Kind() {
			case reflect.Array, reflect.Slice:
				for i := 0; i < val.Len(); i++ {
					item = val.Index(i).Interface()
					if !filterNil || (filterNil && item != nil && fmt.Sprintf("%v", item) != "") {
						merge = append(merge, item)
					}
				}
			}
		}
	}
	return merge
}

// MergeMap 合并字典.
// 相同的键名时,后面的值将覆盖前一个值;key2Str是否将键转换为字符串;ss是元素为字典的数组.
func (ta *TsArr) MergeMap(key2Str bool, arr ...interface{}) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	switch len(arr) {
	case 0:
		break
	default:
		for i, v := range arr {
			val := reflect.ValueOf(v)
			switch val.Kind() {
			case reflect.Map:
				for _, k := range val.MapKeys() {
					if key2Str {
						m[k.String()] = val.MapIndex(k).Interface()
					} else {
						m[k] = val.MapIndex(k).Interface()
					}
				}
			default:
				panic(fmt.Sprintf("[MergeMap]ss type must be map, but [%d]th item not is.", i))
			}
		}
	}
	return m
}

// ArrayChunk 将一个数组分割成多个,size为每个子数组的长度.
func (ta *TsArr) ArrayChunk(arr interface{}, size int) [][]interface{} {
	if size < 1 {
		panic("[ArrayChunk]size: cannot be less than 1")
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 {
			return nil
		}

		chunks := int(math.Ceil(float64(length) / float64(size)))
		var chunk [][]interface{}
		var item []interface{}
		var start int
		for i, end := 0, 0; chunks > 0; chunks-- {
			end = (i + 1) * size
			if end > length {
				end = length
			}

			item = nil
			start = i * size
			for ; start < end; start++ {
				item = append(item, val.Index(start).Interface())
			}
			if item != nil {
				chunk = append(chunk, item)
			}

			i++
		}

		return chunk
	default:
		panic("[ArrayChunk]arr type must be array or slice")
	}
}

// ArrayPad 以指定长度将一个值item填充进数组.
// 若 size 为正，则填补到数组的右侧，如果为负则从左侧开始填补;
// 若 size 的绝对值小于或等于 arr 数组的长度则没有任何填补.
func (ta *TsArr) ArrayPad(arr interface{}, size int, item interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 && size > 0 {
			return ta.ArrayFill(item, size)
		}

		orig := make([]interface{}, length)
		for i := 0; i < length; i++ {
			orig[i] = val.Index(i).Interface()
		}

		if size == 0 || (size > 0 && size < length) || (size < 0 && size > -length) {
			return orig
		}

		n := size
		if size < 0 {
			n = -size
		}
		n -= length
		items := make([]interface{}, n)
		for i := 0; i < n; i++ {
			items[i] = item
		}

		if size > 0 {
			return append(orig, items...)
		}
		return append(items, orig...)
	default:
		panic("[ArrayPad]arr type must be array, slice")
	}
}

// ArraySlice 返回根据 offset 和 size 参数所指定的 arr 数组中的一段切片.
func (ta *TsArr) ArraySlice(arr interface{}, offset, size int) []interface{} {
	if size < 1 {
		panic("[ArraySlice]size: cannot be less than 1")
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 || (offset > 0 && offset > length-1) {
			return nil
		}

		items := make([]interface{}, length)
		for i := 0; i < val.Len(); i++ {
			items[i] = val.Index(i).Interface()
		}

		if offset < 0 {
			offset = offset%length + length
		}
		end := offset + size
		if end < length {
			return items[offset:end]
		}
		return items[offset:]
	default:
		panic("[ArraySlice]arr type must be array or slice")
	}
}

// ArrayRand 从数组中随机取出num个单元.
func (ta *TsArr) ArrayRand(arr interface{}, num int) []interface{} {
	if num < 1 {
		panic("[ArrayRand]num: cannot be less than 1")
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 {
			return nil
		}
		if num > length {
			num = length
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randArr := make([]interface{}, num)
		for i, v := range r.Perm(length) {
			if i < num {
				randArr[i] = val.Index(v).Interface()
			} else {
				break
			}
		}
		return randArr
	default:
		panic("[ArrayRand]arr type must be array or slice")
	}
}

// ArrayColumn 返回数组中指定的一列.
// arr的元素必须是字典;该方法效率低,因为嵌套了两层反射和遍历.
func (ta *TsArr) ArrayColumn(arr interface{}, columnKey string) []interface{} {
	val := reflect.ValueOf(arr)
	var column []interface{}
	var item interface{}
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = val.Index(i).Interface()
			itemVal := reflect.ValueOf(item)
			switch itemVal.Kind() {
			case reflect.Map:
				for _, subKey := range itemVal.MapKeys() {
					if fmt.Sprintf("%s", subKey) == columnKey {
						column = append(column, itemVal.MapIndex(subKey).Interface())
						break
					}
				}
			default:
				panic("[ArrayColumn]arr`s slice item must be map")
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			itemVal := reflect.ValueOf(item)
			switch itemVal.Kind() {
			case reflect.Map:
				for _, subKey := range itemVal.MapKeys() {
					if fmt.Sprintf("%s", subKey) == columnKey {
						column = append(column, itemVal.MapIndex(subKey).Interface())
						break
					}
				}
			default:
				panic("[ArrayColumn]arr`s map item must be map")
			}
		}
	default:
		panic("[ArrayColumn]arr type must be array, slice or map")
	}

	return column
}

// ArrayPush 将一个或多个元素压入数组的末尾(入栈),返回处理之后数组的元素个数.
func (ta *TsArr) ArrayPush(s *[]interface{}, elements ...interface{}) (length int) {
	*s = append(*s, elements...)
	length = len(*s)
	return
}

// ArrayPop 弹出数组最后一个元素(出栈),并返回该元素.
func (ta *TsArr) ArrayPop(s *[]interface{}) (pop interface{}) {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	pop = (*s)[ep]
	*s = (*s)[:ep]
	return
}

// ArrayUnshift 在数组开头插入一个或多个元素,返回处理之后数组的元素个数.
func (ta *TsArr) ArrayUnshift(s *[]interface{}, elements ...interface{}) (length int) {
	*s = append(elements, *s...)
	length = len(*s)
	return
}

// ArrayShift 将数组开头的元素移出数组,并返回该元素.
func (ta *TsArr) ArrayShift(s *[]interface{}) (shift interface{}) {
	if len(*s) == 0 {
		return nil
	}
	shift = (*s)[0]
	*s = (*s)[1:]
	return
}

// ArrayKeyExists 检查数组里是否有指定的键名或索引.
func (ta *TsArr) ArrayKeyExists(key interface{}, arr interface{}) (r bool) {
	if key == nil {
		return
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		var keyInt int
		var keyIsInt, ok bool
		if keyInt, ok = key.(int); ok {
			keyIsInt = true
		}

		length := val.Len()
		if keyIsInt && length > 0 && keyInt >= 0 && keyInt < length {
			r = true
			return
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if fmt.Sprintf("%s", key) == fmt.Sprintf("%s", k) || reflect.DeepEqual(key, k) {
				r = true
				return
			}
		}
	default:
		panic("[ArrayKeyExists]arr type must be array, slice or map")
	}
	return
}

// ArrayReverse 返回单元顺序相反的数组(仅限数组和切片).
func (ta *TsArr) ArrayReverse(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		reverse := make([]interface{}, length)
		i, j := 0, length-1
		for ; i < j; i, j = i+1, j-1 {
			reverse[i], reverse[j] = val.Index(j).Interface(), val.Index(i).Interface()
		}
		if length > 0 && reverse[j] == nil {
			reverse[j] = val.Index(j).Interface()
		}

		return reverse
	default:
		panic("[ArrayReverse]arr type must be array, slice")
	}
}

// Implode 用delimiter将数组(数组/切片/字典)的值连接为一个字符串.
func (ta *TsArr) Implode(delimiter string, arr interface{}) (s string) {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 {
			return
		}
		var buf bytes.Buffer
		j := length
		for i := 0; i < length; i++ {
			buf.WriteString(fmt.Sprintf("%s", val.Index(i)))
			if j--; j > 0 {
				buf.WriteString(delimiter)
			}
		}
		s = buf.String()
		return
	case reflect.Map:
		length := len(val.MapKeys())
		if length == 0 {
			return ""
		}
		var buf bytes.Buffer
		for _, k := range val.MapKeys() {
			buf.WriteString(fmt.Sprintf("%s", val.MapIndex(k).Interface()))
			if length--; length > 0 {
				buf.WriteString(delimiter)
			}
		}
		s = buf.String()
		return
	default:
		panic("[Implode]arr type must be array, slice")
	}
}

// JoinStrings 使用分隔符delimiter连接字符串数组.效率比Implode高.
func (ta *TsArr) JoinStrings(strArr []string, delimiter string) (join string) {
	length := len(strArr)
	if length == 0 {
		return
	}

	var sb strings.Builder
	for _, str := range strArr {
		sb.WriteString(str)
		if length--; length > 0 {
			sb.WriteString(delimiter)
		}
	}
	join = sb.String()
	return
}

// JoinIntArr 使用分隔符delimiter连接整数数组.
func (ta *TsArr) JoinIntArr(intArr []int, delimiter string) (join string) {
	length := len(intArr)
	if length == 0 {
		return
	}

	var sb strings.Builder
	for _, num := range intArr {
		sb.WriteString(strconv.Itoa(num))
		if length--; length > 0 {
			sb.WriteString(delimiter)
		}
	}
	join = sb.String()

	return
}

// UniqueIntArr 移除整数数组中的重复值.
func (ta *TsArr) UniqueIntArr(intArr []int) (unique []int) {
	sort.Ints(intArr)
	var last int
	for i, num := range intArr {
		if i == 0 || num != last {
			unique = append(unique, num)
		}
		last = num
	}
	return
}

// Unique64IntArr 移除64位整数数组中的重复值.
func (ta *TsArr) Unique64IntArr(intArr []int64) (unique []int64) {
	seen := make(map[int64]bool)
	for _, num := range intArr {
		if _, ok := seen[num]; !ok {
			seen[num] = true
			unique = append(unique, num)
		}
	}
	return
}

// UniqueStringsArr 移除字符串数组中的重复值.
func (ta *TsArr) UniqueStringsArr(strArr []string) (unique []string) {
	sort.Strings(strArr)
	var last string
	for _, str := range strArr {
		if str != last {
			unique = append(unique, str)
		}
		last = str
	}

	return
}

// ArrayDiff 计算数组(数组/切片/字典)的差集,返回在 arr1 中但是不在 arr2 里,且非空元素(nil,'')的值.
func (ta *TsArr) ArrayDiff(arrCompare, arrToBeCompare interface{}) []interface{} {
	valA := reflect.ValueOf(arrCompare)
	valB := reflect.ValueOf(arrToBeCompare)
	var diffArr []interface{}
	var item interface{}
	var notInB bool

	if (valA.Kind() == reflect.Array || valA.Kind() == reflect.Slice) && (valB.Kind() == reflect.Array || valB.Kind() == reflect.Slice) {
		//两者都是数组/切片
		if valA.Len() == 0 {
			return nil
		} else if valB.Len() == 0 {
			return ta.arrayValuesHelper(arrCompare, true)
		}

		for i := 0; i < valA.Len(); i++ {
			item = valA.Index(i).Interface()
			notInB = true
			for j := 0; j < valB.Len(); j++ {
				if reflect.DeepEqual(item, valB.Index(j).Interface()) {
					notInB = false
					break
				}
			}

			if notInB {
				diffArr = append(diffArr, item)
			}
		}
	} else if (valA.Kind() == reflect.Array || valA.Kind() == reflect.Slice) && (valB.Kind() == reflect.Map) {
		//A是数组/切片,B是字典
		if valA.Len() == 0 {
			return nil
		} else if len(valB.MapKeys()) == 0 {
			return ta.arrayValuesHelper(arrCompare, true)
		}

		for i := 0; i < valA.Len(); i++ {
			item = valA.Index(i).Interface()
			notInB = true
			for _, k := range valB.MapKeys() {
				if reflect.DeepEqual(item, valB.MapIndex(k).Interface()) {
					notInB = false
					break
				}
			}

			if notInB {
				diffArr = append(diffArr, item)
			}
		}
	} else if (valA.Kind() == reflect.Map) && (valB.Kind() == reflect.Array || valB.Kind() == reflect.Slice) {
		//A是字典,B是数组/切片
		if len(valA.MapKeys()) == 0 {
			return nil
		} else if valB.Len() == 0 {
			return ta.arrayValuesHelper(arrCompare, true)
		}

		for _, k := range valA.MapKeys() {
			item = valA.MapIndex(k).Interface()
			notInB = true
			for i := 0; i < valB.Len(); i++ {
				if reflect.DeepEqual(item, valB.Index(i).Interface()) {
					notInB = false
					break
				}
			}

			if notInB {
				diffArr = append(diffArr, item)
			}
		}
	} else if (valA.Kind() == reflect.Map) && (valB.Kind() == reflect.Map) {
		//两者都是字典
		if len(valA.MapKeys()) == 0 {
			return nil
		} else if len(valB.MapKeys()) == 0 {
			return ta.arrayValuesHelper(arrCompare, true)
		}

		for _, k := range valA.MapKeys() {
			item = valA.MapIndex(k).Interface()
			notInB = true
			for _, k2 := range valB.MapKeys() {
				if reflect.DeepEqual(item, valB.MapIndex(k2).Interface()) {
					notInB = false
					break
				}
			}

			if notInB {
				diffArr = append(diffArr, item)
			}
		}
	} else {
		panic("[ArrayDiff]arr1, arr2 type must be array, slice or map")
	}

	return diffArr
}

// ArrayUnique 移除数组中重复的值.
func (ta *TsArr) ArrayUnique(arr interface{}) (unique []interface{}) {
	val := reflect.ValueOf(arr)
	var item interface{}
	var str, key string
	mp := make(map[string]interface{})
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = val.Index(i).Interface()
			str = fmt.Sprintf("%+v", item)
			key = string(TStr.Md5Hex([]byte(str), 32))
			if _, ok := mp[key]; !ok {
				mp[key] = true
				unique = append(unique, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			str = fmt.Sprintf("%+v", item)
			key = string(TStr.Md5Hex([]byte(str), 32))
			if _, ok := mp[key]; !ok {
				mp[key] = true
				unique = append(unique, item)
			}
		}
	default:
		panic("[ArrayUnique]arr type must be array, slice or map")
	}

	return
}

// ArraySearchItem 从数组中搜索对应元素(单个).
// arr为要查找的数组,condition为条件字典.
func (ta *TsArr) ArraySearchItem(arr interface{}, condition map[string]interface{}) (search interface{}) {
	// 条件为空
	if len(condition) == 0 {
		return
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			search = ta.compareConditionHelper(condition, val.Index(i).Interface())
			if search != nil {
				return
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			search = ta.compareConditionHelper(condition, val.MapIndex(k).Interface())
			if search != nil {
				return
			}
		}
	default:
		panic("[ArraySearchItem]arr type must be array, slice or map")
	}

	return
}

// ArraySearchMulti 从数组中搜索对应元素(多个).
// arr为要查找的数组,condition为条件字典.
func (ta *TsArr) ArraySearchMulti(arr interface{}, condition map[string]interface{}) (search []interface{}) {
	// 条件为空
	if len(condition) == 0 {
		return
	}

	val := reflect.ValueOf(arr)
	var item interface{}
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = ta.compareConditionHelper(condition, val.Index(i).Interface())
			if item != nil {
				search = append(search, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = ta.compareConditionHelper(condition, val.MapIndex(k).Interface())
			if item != nil {
				search = append(search, item)
			}
		}
	default:
		panic("[ArraySearchMulti]arr type must be array, slice or map")
	}

	return
}

// ArrayCombine 合并两个数组来创建一个新数组，其中的一个数组元素为键名，另一个数组元素为键值
func (ta *TsArr) ArrayCombine(keys, values []interface{}) map[interface{}]interface{} {
	m := make(map[interface{}]interface{}, len(values))
	if len(keys) != len(values) {
		return m
	}
	for i, v := range keys {
		m[v] = values[i]
	}
	return m
}

// arrayValuesHelper 返回数组/切片/字典中所有的值.
// filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素.
func (ta *TsArr) arrayValuesHelper(arr interface{}, filterNil bool) (values []interface{}) {
	var item interface{}
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = val.Index(i).Interface()
			if !filterNil || (filterNil && item != nil && fmt.Sprintf("%v", item) != "") {
				values = append(values, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			if !filterNil || (filterNil && item != nil && fmt.Sprintf("%v", item) != "") {
				values = append(values, item)
			}
		}
	default:
		panic("[arrayValuesHelper]arr type must be array, slice or map")
	}

	return
}

// compareConditionHelper 比对数组是否匹配条件.condition为条件字典,arr为要比对的数据数组.
func (ta *TsArr) compareConditionHelper(condition map[string]interface{}, arr interface{}) (compare interface{}) {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Map:
		condLen := len(condition)
		chkNum := 0
		if condLen > 0 {
			for _, k := range val.MapKeys() {
				if condVal, ok := condition[k.String()]; ok && reflect.DeepEqual(val.MapIndex(k).Interface(), condVal) {
					chkNum++
				}
			}
		}
		if chkNum == condLen {
			compare = arr
		}
	default:
		return
	}
	return
}
