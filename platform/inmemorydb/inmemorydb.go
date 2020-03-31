package inmemorydb

// Item is a single inmemorydb item
type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Repo DS
type Repo struct {
	Items []Item
}

// New Constructor function
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// Add new item to memory
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// GetAll data from memory
func (r *Repo) GetAll() []Item {
	return r.Items
}
