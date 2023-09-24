package skill_assessment

import (
	"testing"
)

func TestCreatePersons(t *testing.T) {
	person := CreatePerson("test", 7, "test@bob.de")
	if person.Name != "test" ||
		person.Age != 7 ||
		person.Email != "test@bob.de" {
		t.Errorf("Unable to create person: %v", person)
	}
}

func TestPerson_GetInfo(t *testing.T) {
	person := CreatePerson("test1", 7, "test@bob.de")
	expectedInfo := "Name: test1, Age: 7, Email: test@bob.de"
	personInfo := person.GetInfo()
	if personInfo != expectedInfo {
		t.Errorf(
			"PersonInfo %s info does not match the expected string: %s",
			personInfo, expectedInfo,
		)
	}
}
