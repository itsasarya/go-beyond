package main

type Contact struct {
	Name  string
	Phone string
}

var contacts []Contact

func AddContact(name, phone string) Contact {
	newContact := Contact{Name: name, Phone: phone}
	contacts = append(contacts, newContact)
	return newContact
}

func UpdateContact(name, phone string) *Contact {
	for i := range contacts {
		if contacts[i].Name == name {
			contacts[i].Phone = phone
			return &contacts[i]
		}
	}
	return nil
}

func FindContact(name string) *Contact {
	for i := range contacts {
		if contacts[i].Name == name {
			return &contacts[i]
		}
	}
	return nil
}

func ListContact() []Contact {
	return contacts
}
