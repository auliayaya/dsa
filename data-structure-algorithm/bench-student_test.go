package data_structure_algorithm

import "testing"

func BenchmarkStudentJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StudentJSON()
	}
}
func BenchmarkStudentsJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StudentsJSON()
	}
}
