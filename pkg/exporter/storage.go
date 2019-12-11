package exporter

type Storage struct {
	url      string
	username string
	password string
}

func NewStorage(storage *Storage) (*Storage, error) {
	return storage, nil
}

func (c *Storage) GetUsage() (avg int, err error) {
	avg = 1234
	return
}
