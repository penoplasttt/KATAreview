package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	State string `db:"state"`
}

type CityRepository struct {
	db *sqlx.DB
}

func NewCityRepository(db *sqlx.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) Create(city *City) error {
	query := "INSERT INTO cities (name, state) VALUES (?, ?)"
	_, err := r.db.Exec(query, city.Name, city.State)
	return err
}

func (r *CityRepository) Delete(id int) error {
	query := "DELETE FROM cities WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *CityRepository) Update(city *City) error {
	query := "UPDATE cities SET name = ?, state = ? WHERE id = ?"
	_, err := r.db.Exec(query, city.Name, city.State, city.Id)
	return err
}

func (r *CityRepository) List() ([]City, error) {
	var cities []City
	query := "SELECT * FROM cities"
	err := r.db.Select(&cities, query)
	return cities, err
}

func main() {
	db, err := sqlx.Connect("mysql", "")
	if err != nil {
		log.Println(err)
	}

	repo := NewCityRepository(db)

	city := City{
		Id: 1,
		Name: "Chipi",
		State: "Chapa",
	}

	err = repo.Create(&city)
	if err != nil {
		log.Println(err)
	}
}
