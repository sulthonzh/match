package match

import "testing"

// func TestCombination(t *testing.T) {
// 	stringList := [][]interface{}{}
// 	stringList = append(stringList, []interface{}{"Kilometer_Dekat", "Kilometer_Sedang", "Kilometer_Jauh"})
// 	stringList = append(stringList, []interface{}{"And Umur_Baru", "And Umur_Sedang", "And Umur_Lama"})

// 	f := GetCombination(0, stringList)
// 	fmt.Printf("Result: %v\n", f)

// 	float64List := [][]interface{}{}
// 	float64List = append(float64List, []interface{}{7, 8, 9})
// 	float64List = append(float64List, []interface{}{4, 6, 6})

// 	f = GetCombination(0, float64List)
// 	fmt.Printf("Result: %v\n", f)
// }

// func TestCombo(t *testing.T) {
// 	b := comboChain([][]anyType{
// 		[]anyType{"Abu", "Ali", "Bakar"},
// 		[]anyType{"Affendi", "Azhar", "Hamin"},
// 		[]anyType{"living in", "born in"},
// 		[]anyType{"Japan", "Malaysia", "Singapore"},
// 		[]anyType{"since"},
// 		[]anyType{1990, 1980, 2000},
// 	})

// 	for o := range b {
// 		fmt.Println(o)
// 	}
// }
var inputs = [][]anyType{
	{"Abu", "Ali", "Bakar"},
	{"Affendi", "Azhar", "Hamin"},
	{"living in"},
	{"Japan", "Malaysia", "Singapore"},
}

func BenchmarkChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b := comboChain(inputs)
		for _ = range b {
		}
	}
}

func BenchmarkOrig(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getCombination(0, inputs)
	}
}

func BenchmarkCombine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		combine(func(row []anyType) {
			//fmt.Println(row)
		}, inputs...)
	}
}
