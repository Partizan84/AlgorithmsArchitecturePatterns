package interfaces

type Repository interface {
	Save(item *Models.Item) error

	Get(id string) (*Models.Item, error)

	GetAll() (Models.Items, error)

	Delete(id string) error
}
