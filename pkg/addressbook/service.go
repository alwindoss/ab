package addressbook

type Contact struct {
}

type Service interface {
	AddContact(c *Contact) (*Contact, error)
	UpdateContact(c *Contact) (*Contact, error)
	DeleteContact(id string) error
	FetchContact(id string) (*Contact, error)
	FetchAllContacts() ([]*Contact, error)
}

func NewService(r Repository) Service {
	return &addressBookService{}
}

type addressBookService struct {
}

// AddContact implements Service.
func (a *addressBookService) AddContact(c *Contact) (*Contact, error) {
	panic("unimplemented")
}

// DeleteContact implements Service.
func (a *addressBookService) DeleteContact(id string) error {
	panic("unimplemented")
}

// FetchAllContacts implements Service.
func (a *addressBookService) FetchAllContacts() ([]*Contact, error) {
	panic("unimplemented")
}

// FetchContact implements Service.
func (a *addressBookService) FetchContact(id string) (*Contact, error) {
	panic("unimplemented")
}

// UpdateContact implements Service.
func (a *addressBookService) UpdateContact(c *Contact) (*Contact, error) {
	panic("unimplemented")
}
