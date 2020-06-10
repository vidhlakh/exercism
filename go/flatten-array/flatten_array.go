//Package flatten flattens the deep nested array into single list while also removing nil,null
package flatten

// Flatten flattens the array
//for slice of interfaces , refer https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d
func Flatten(array interface{}) []interface{} {
	slc := make([]interface{}, 0)
	for _, v := range array.([]interface{}) {
		switch v.(type) {
		case []interface{}:
			slc = append(slc, Flatten(v)...)
		case int:
			slc = append(slc, v)
		}
	}
	return slc
}
