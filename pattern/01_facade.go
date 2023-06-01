package main

import (
	"fmt"
)

// Data данные которые берем из бд или кеша
type Data struct{}

// DB ...
type DB struct{}

func (d *DB) Connect() error {
	fmt.Println("Connect to database")
	// ...
	return nil
}

// GetAll возвращает все данные с базы данных
func (d *DB) GetAll() ([]*Data, error) {
	data := make([]*Data, 0)
	//...
	return data, nil
}

// Close ...
func (d *DB) Close() {
	//...
}

// Cache ...
type Cache struct{}

// Connect ...
func (c *Cache) Connect() error {
	fmt.Println("Connect to cache")
	//...
	return nil
}

// Add добавляет данные в кеш
func (c *Cache) Add(_ []*Data) error {
	//...
	return nil
}

// Close ...
func (c *Cache) Close() {
	//...
}

// Store фасад хранилищей: базы данных и кеша
type Store struct {
	db    *DB
	cache *Cache
}

// RestoreCache восстанавливает кеш из базы данных
func (s *Store) RestoreCache() error {
	fmt.Println("Restore cache")
	data, err := s.db.GetAll()
	if err != nil {
		return err
	}
	if err := s.cache.Add(data); err != nil {
		return err
	}
	return nil
}

func main() {
	db, cache := &DB{}, &Cache{}
	if err := db.Connect(); err != nil {
		panic(err)
	}
	if err := cache.Connect(); err != nil {
		panic(err)
	}
	store := &Store{db, cache}
	if err := store.RestoreCache(); err != nil {
		panic(err)
	}
}
