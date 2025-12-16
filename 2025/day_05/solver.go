package main

func CountFreshIngredients(db *Database) int {
	count := 0
	for _, id := range db.Available {
		if db.IsFresh(id) {
			count++
		}
	}
	return count
}
