package skill_assessment

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name  string
	Age   uint8
	Email string
}

type PersonCallback func(string, uint8, string) *Person

func CreatePerson(name string, age uint8, email string) *Person {
	return &Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func (p *Person) GetInfo() string {
	return fmt.Sprintf(
		"Name: %s, Age: %d, Email: %s",
		p.Name, p.Age, p.Email,
	)
}

type PersonChannel chan *Person

func (pC PersonChannel) CreateAndSendPerson(
	callback PersonCallback, name string, age uint8, email string,
) {
	pC <- callback(name, age, email)
}

func SkillAssessmentMain() {
	channel := make(PersonChannel)

	go channel.CreateAndSendPerson(CreatePerson, "Guenter", 93, "person2@test.de")
	go channel.CreateAndSendPerson(CreatePerson, "Christine", 73, "person1@test.de")

	person1 := <-channel
	person2 := <-channel

	persons := fmt.Sprintf("%s\n%s", person1.GetInfo(), person2.GetInfo())
	personsByte, err := json.Marshal(persons)
	if err != nil {
		log.Panic("Unable to marshal persons!", err)
	}

	err = os.WriteFile("persons.txt", personsByte, 777)
	if err != nil {
		log.Panic("Unable to write file!", err)
	}
}
