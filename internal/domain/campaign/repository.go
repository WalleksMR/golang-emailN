package campaign

type Repository interface {
	Save(campaign *Campaign) error
	ListAll() ([]Campaign, error)
	GetById(id string) (*Campaign, error)
}
