// Package sublist tells whether relation b/w 2 list
package sublist

import "reflect"

//Relation is return type   which can be superlist, sublist, equal, unequl
type Relation string

//Sublist finds relation with te lsit2 with respect to list1
func Sublist(list1, list2 []int) Relation {

	// if len(list1) == 0 && len(list2) == 0 {
	// 	return Relation("equal")
	// }
	if reflect.DeepEqual(list1, list2) {
		return Relation("equal")
		// if isEqual(list1, list2) {
		// 	return Relation("equal")

	} else if isSublist(list2, list1) {
		return Relation("superlist")
	} else if isSublist(list1, list2) {
		return Relation("sublist")
	}
	return Relation("unequal")

}

//isSublist check whether list1 is sublist of list2
func isSublist(listOne []int, listTwo []int) bool {
	if len(listOne) >= len(listTwo) {
		return false
	}
	if len(listOne) == 0 {
		return true
	}
	for i := 0; i < len(listTwo); i++ {
		if listTwo[i] == listOne[0] {
			j := 1
			for j < len(listOne) {
				if i+j >= len(listTwo) {
					return false
				}
				if listTwo[i+j] != listOne[j] {
					break
				}
				j++
			}
			if j == len(listOne) {
				return true
			}
		}
	}
	return false
}

// isEqual checks whethe 2 lists are equal
func isEqual(listOne []int, listTwo []int) bool {
	if len(listOne) != len(listTwo) {
		return false
	}
	for i, v := range listOne {
		if listTwo[i] != v {
			return false
		}
	}
	return true
}
