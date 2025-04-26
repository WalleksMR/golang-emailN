package campaign

type Repository interface {
	Save(campaign *Campaign) error
	ListAll() ([]Campaign, error)
}
