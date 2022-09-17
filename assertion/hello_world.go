package assertion

func HelloWorld(name string) string {
	return "Hello " + name
}

type Person struct {
	Name string
}

func HelloWorldObject(person Person) interface{} {
	if person.Name == "" {
		return nil
	} else {
		return person.Name
	}
}
