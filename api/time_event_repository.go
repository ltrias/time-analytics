package api

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

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

func loadSuggest(field string, db *sql.DB) []string {
	query := fmt.Sprintf("SELECT DISTINCT %s FROM time_event ORDER BY %s ASC", field, field)

	rows, err := db.Query(query)

	var result []string

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var r string
		rows.Scan(&r)

		result = append(result, r)
	}

	return result
}

func (t *TimeEventRepository) LoadTypeSuggest() []string {
	return loadSuggest("tipo", t.db)
}

func (t *TimeEventRepository) LoadWhoSuggest() []string {
	return loadSuggest("quem", t.db)
}

// func (t *TimeEventRepository) LoadTypeSuggest() []string {
// 	return loadSuggest("", t.db)
// }

// func (t *TimeEventRepository) LoadTypeSuggest() []string {
// 	return loadSuggest("", t.db)
// }

// func (t *TimeEventRepository) LoadTypeSuggest() []string {
// 	return loadSuggest("", t.db)
// }

func (t *TimeEventRepository) LoadDurationSuggest() []int {
	temp := loadSuggest("tempo_ocupado", t.db)

	log.Println(len(temp))

	var result []int

	for _, v := range temp {
		intV, _ := strconv.Atoi(v)
		result = append(result, intV)
	}

	return result
}

func (t *TimeEventRepository) LoadEvent(id int) TimeEvent {
	stmt, err := t.db.Prepare("SELECT id, dia, tipo, quem, tempo_ocupado, tema, departamento, recorrente FROM time_event WHERE id=?")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result TimeEvent
	err = stmt.QueryRow(id).Scan(&result.ID, &result.Day, &result.Type, &result.Who, &result.Duration, &result.Subject, &result.Department, &result.Recurrent)

	if err != nil {
		log.Fatal(err)
	}

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
