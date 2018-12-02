package main

import (
	"errors"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const PersonCollection = "person"

var ErrDuplicatedPerson = errors.New("Duplicated person")

type Person struct {
	Id      string `bson:"_id"`
	Name    string `bson:"name"`
	Inative bool   `bson:inative`
}

type PersonRepository struct {
	session *mgo.Session
}

func (r *PersonRepository) Create(p *Person) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PersonCollection)
	err := collection.Insert(p)
	mongoErr, ok := err.(*mgo.LastError)
	if ok && mongoErr.Code == 11000 {
		return ErrDuplicatedPerson
	}
	return err
}

func (r *PersonRepository) Update(p *Person) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PersonCollection)
	return collection.Update(bson.M{"_id": p.Id}, p)
}

func (r *PersonRepository) Remove(id string) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PersonCollection)
	return collection.Remove(bson.M{"_id": id})
}

func (r *PersonRepository) FindAllActive() ([]*Person, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PersonCollection)
	query := bson.M{"inative": false}

	documents := make([]*Person, 0)

	err := collection.Find(query).All(&documents)
	return documents, err
}

func (r *PersonRepository) FindById(id string) (*Person, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PersonCollection)
	query := bson.M{"_id": id}

	person := &Person{}

	err := collection.Find(query).One(person)
	return person, err
}

func NewPersonRepository(session *mgo.Session) *PersonRepository {
	return &PersonRepository{session}
}

func main() {
	session, err := mgo.Dial("localhost:27017/go-course")

	if err != nil {
		log.Fatal(err)
	}

	repository := NewPersonRepository(session)

	// creating a person
	person := &Person{Id: "123", Name: "Juliana"}
	err = repository.Create(person)

	if err == ErrDuplicatedPerson {
		log.Printf("%s is already created\n", person.Name)
	} else if err != nil {
		log.Println("Failed to create a person: ", err)
	}

	// updating a person
	person.Name = "Juliana updated"
	err = repository.Update(person)

	if err != nil {
		log.Println("Failed to update a person: ", err)
	}

	repository.Create(&Person{Id: "124", Name: "Marcos"})
	repository.Create(&Person{Id: "125", Name: "Kaio", Inative: true})
	repository.Create(&Person{Id: "126", Name: "Gabriel"})
	repository.Create(&Person{Id: "127", Name: "Maisa"})

	// remove
	err = repository.Remove("126")

	if err != nil {
		log.Println("Failed to remove a person: ", err)
	}

	// findAll
	people, err := repository.FindAllActive()
	if err != nil {
		log.Println("Failed to fetch people: ", err)
	}

	for _, person := range people {
		log.Printf("Have in database: %#v\n", person)
	}

	// FindById
	person, err = repository.FindById("123")
	if err == nil {
		log.Printf("Result of findById: %v\n", person)
	} else {
		log.Println("Failed to findById ", err)
	}
}
