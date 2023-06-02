package generics

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAddStudentString(t *testing.T) {
	var students []String
	result := addStudent[String](students, "Udin")
	result = addStudent[String](result, "Khoirudin")
	result = addStudent[String](result, "Rosyadi")
	assert.Equal(t, 3, len(result))
	assert.Equal(t, "String", reflect.TypeOf(result[0]).Name())
	assert.Contains(t, "Udin", result[0])
	assert.Contains(t, "Khoirudin", result[1])
	assert.Contains(t, "Rosyadi", result[2])
}

func TestAddStudentInt(t *testing.T) {
	students1 := []Integer{}
	result1 := addStudent[Integer](students1, 45)
	result1 = addStudent[Integer](result1, 64)
	result1 = addStudent[Integer](result1, 78)
	assert.Equal(t, 3, len(result1))
	assert.Equal(t, "Integer", reflect.TypeOf(result1[0]).Name())
	assert.EqualValues(t, 45, result1[0])
	assert.EqualValues(t, 64, result1[1])
	assert.EqualValues(t, 78, result1[2])
}

func TestAddStudent(t *testing.T) {
	students := []Student{}
	result := addStudent[Student](students, Student{"John", 213, 17.5})
	result = addStudent[Student](result, Student{"John", 214, 17.5})
	result = addStudent[Student](result, Student{"John", 215, 17.5})
	assert.Equal(t, 3, len(result))
	assert.Equal(t, "Student", reflect.TypeOf(result[0]).Name())
	assert.EqualValues(t, 213, result[0].ID)
	assert.EqualValues(t, 214, result[1].ID)
	assert.EqualValues(t, 215, result[2].ID)
}
