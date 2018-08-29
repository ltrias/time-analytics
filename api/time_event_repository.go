package api

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//TimeEventRepository repository for time events
type TimeEventRepository struct {
	db *sql.DB
}

//NewTimeEventRepository constructs a new instance
func NewTimeEventRepository() *TimeEventRepository {
	result := new(TimeEventRepository)

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/time_analytics?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	result.db = db

	return result
}

//LoadAllEvents loads all events
func (t *TimeEventRepository) LoadAllEvents() []TimeEvent {
	stmt, err := t.db.Prepare("SELECT id, dia, tipo, quem, tempo_ocupado, tema, departamento, recorrente FROM time_event")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	defer rows.Close()
	var result []TimeEvent

	for rows.Next() {
		r := TimeEvent{}
		err := rows.Scan(&r.ID, &r.Day, &r.Type, &r.Who, &r.Duration, &r.Subject, &r.Department, &r.Recurrent)

		if err != nil {
			log.Fatal(err)
		}
		result = append(result, r)
	}

	return result
}
