package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Email    string 
	Password string 
	Name     string 
	Age      int    
}

type Database interface {
	RegisterUser(user User) error
	GetAllUsers() ([]User, error)
}

type SQLDatabase struct {
	db *sql.DB
}

func NewSQLDatabase() (*SQLDatabase, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			email TEXT PRIMARY KEY,
			password TEXT,
			name TEXT,
			age INTEGER
		);
	`)
	if err != nil {
		return nil, err
	}

	return &SQLDatabase{db: db}, nil
}

func (h *SQLDatabase) RegisterUser(user User) error {
	//проверяю email и возраст
	var count int
	err := h.db.QueryRow(`
		SELECT COUNT(*) FROM users
		WHERE email = ? OR (age < 18 AND age = ?)
	`, user.Email, user.Age).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("User with the same email or age below 18 already exists")
	}

	//регистрирую
	_, err = h.db.Exec(`
		INSERT INTO users (email, password, name, age)
		VALUES (?, ?, ?, ?)
	`, user.Email, user.Password, user.Name, user.Age)

	return err
}

func (h *SQLDatabase) GetAllUsers() ([]User, error) {
	rows, err := h.db.Query(`SELECT email, password, name, age FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Email, &user.Password, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

// реализует интерфейс
type InMemoryCache struct {
	cache map[string]interface{}
	mu    sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		cache: make(map[string]interface{}),
	}
}

func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.cache[key]
	return value, ok
}

func (c *InMemoryCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

// реализует Database с использованием кэша
type CachedDatabase struct {
	Database
	Cache
}

func NewCachedDatabase(db Database, cache Cache) *CachedDatabase {
	return &CachedDatabase{
		Database: db,
		Cache:    cache,
	}
}

func (h *CachedDatabase) GetAllUsers() ([]User, error) {
	if value, ok := h.Get("all_users"); ok {
		if users, ok := value.([]User); ok {
			return users, nil
		}
	}

	users, err := h.Database.GetAllUsers()
	if err != nil {
		return nil, err
	}

	h.Set("all_users", users)

	return users, nil
}

func main() {
	r := chi.NewRouter()

	db, err := NewSQLDatabase()
	if err != nil {
		log.Fatal(err)
	}

	cache := NewInMemoryCache()
	cachedDB := NewCachedDatabase(db, cache)

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := cachedDB.RegisterUser(user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := cachedDB.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
