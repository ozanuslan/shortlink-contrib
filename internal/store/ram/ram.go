package ram

import (
	"fmt"
	"sync"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

// RAMLinkList implementation of store interface
type RAMLinkList struct { // nolint unused
	links map[string]link.Link
	mu    sync.RWMutex
}

// Init ...
func (ram *RAMLinkList) Init() error { // nolint unparam
	ram.mu.Lock()
	ram.links = make(map[string]link.Link)
	ram.mu.Unlock()

	return nil
}

// Close ...
func (ram *RAMLinkList) Close() error {
	return nil
}

// Migrate ...
func (ram *RAMLinkList) migrate() error { // nolint unused
	return nil
}

// Get ...
func (ram *RAMLinkList) Get(id string) (*link.Link, error) {
	ram.mu.Lock()
	response := ram.links[id]
	ram.mu.Unlock()

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (ram *RAMLinkList) List(filter *query.Filter) ([]*link.Link, error) { // nolint unused
	links := []*link.Link{}

	ram.mu.Lock()
	// copy map by assigning elements to new map
	for key := range ram.links {
		links = append(links, &link.Link{
			Url:       ram.links[key].Url,
			Hash:      ram.links[key].Hash,
			Describe:  ram.links[key].Describe,
			CreatedAt: ram.links[key].CreatedAt,
			UpdatedAt: ram.links[key].UpdatedAt,
		})
	}
	ram.mu.Unlock()

	return links, nil
}

// Add ...
func (ram *RAMLinkList) Add(source link.Link) (*link.Link, error) { // nolint unused
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	ram.mu.RLock()
	ram.links[data.Hash] = data
	ram.mu.RUnlock()

	return &data, nil
}

// Update ...
func (ram *RAMLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (ram *RAMLinkList) Delete(id string) error { // nolint unused
	ram.mu.Lock()
	delete(ram.links, id)
	ram.mu.Unlock()

	return nil
}
