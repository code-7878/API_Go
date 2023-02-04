package main

type Cat struct {
	tableName struct{} `pg:"cats"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	IsStrip   bool     `json:"is_strip" pg:"IS_STRIP"`
	Color     string   `json:"color"  pg:"COLOR"`
}

// Получить всех котиков из таблицы.

func FindAllCats() []Cat {
	var cats []Cat
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&cats).Select()

	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return cats

}

//добавляем котика в таблицу

func CreateCat(cat Cat) Cat {
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&cat).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()

	return cat
}
