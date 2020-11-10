package models

const CarsTableName = "cars"

type Car struct {
	ID       string `db:"id"`
	Model    string `db:"model"`
	Number   string `db:"number"`
	Status   string `db:"status"`
	ImageUrl string `db:"image_url"`
}

func (c Car) TableName() string {
	return CarsTableName
}
