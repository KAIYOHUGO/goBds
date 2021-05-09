package gc

import "gobds/src/database"

func GC() {
	database.GC()
}
