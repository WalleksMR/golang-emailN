package contract

type CampaingUpateInput struct {
	ID       string
	Name     *string
	Status   *string
	Content  *string
	Contacts *[]string
}
