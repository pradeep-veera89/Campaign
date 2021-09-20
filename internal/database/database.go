package database

import (
	"log"
	"time"

	"github.com/pradeep-veera89/campaign/internal/models"
)

func (db *DB) InsertLead(l models.Lead) {
	stmt := `insert into leads(email, first_name, last_name,  salutation, created_at, updated_at)
			  values($1,$2,$3,$4,$5,$6)`
	_, err := db.SQL.Query(stmt,
		l.EMail,
		l.FirstName,
		l.LastName,
		l.Salutation,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		log.Println("Insert Error: ", err)
	}
}

func (db *DB) GetLeadByEmail() {

}

func (db *DB) TestPrint() {
	log.Println("inside testPrint")
}

type DatabaseRepo interface {
	InsertLead(models.Lead)
	GetLeadByEmail()
	TestPrint()
}

var DBRepo DatabaseRepo

var Conn DB

func InitiateDB(dbConn *DB) {
	Conn = *dbConn
}

func GetDB() *DB {
	return &Conn
}
