package main

type Dog struct {
	tableName struct{} `pg:"dogs"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	Proda     bool     `json:"Proda" pg:"Proda"`
	Color     string   `json:"color"  pg:"color"`
}

// Получить всех котиков из таблицы.

func FindAllDogs() []Dog {
	var dogs []Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dogs).Select()

	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dogs

}

//добавляем котика в таблицу

func CreateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()

	return dog
}

func FindDogById(dogId string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dog).Where("id = ?", dogId).First()

	if err != nil {
		panic(err)

	}
	pgConnect.Close()
	return dog
}

func DropDogById(dogId string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	dog = FindDogById(dogId)

	_, err := pgConnect.Model(&dog).Where("id = ?", dogId).Delete()

	if err != nil {
		panic(err)

	}
	pgConnect.Close()
	return dog
}

func UpdateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()
	oldDog := FindDogById(dog.ID)

	oldDog.Name = dog.Name
	oldDog.Proda = dog.Proda
	oldDog.Color = dog.Color

	_, err := pgConnect.Model(&oldDog).
		Set("name = ?", oldDog.Name).
		Set("Proda = ?", oldDog.Proda).
		Set("Color = ?", oldDog.Color).
		Where("id = ?", oldDog.ID).
		Update()
	if err != nil {
		panic(err)

	}
	pgConnect.Close()
	return oldDog

}
