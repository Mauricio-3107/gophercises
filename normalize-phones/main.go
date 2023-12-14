package main

import (
	"fmt"
	"regexp"

	phonedb "github.com/Mauricio-3107/gophercises.git/normalize-phones/db"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "revelare"
	dbname   = "gophercises_phone"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	must(phonedb.Reset("postgres", psqlInfo, dbname))

	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(phonedb.Migrate("postgres", psqlInfo))

	db, err := phonedb.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	must(db.Seed())

	phones, err := db.AllPhones()
	must(err)
	for _, p := range phones {
		number := normalize(p.Number)
		if p.Number != number {
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				// Delete phone since was found in the db
				err := db.DeletePhone(p.Id)
				must(err)
			} else {
				// Update phone because its not present in the db with the normalized data
				p.Number = number
				err := db.UpdatePhone(&p)
				must(err)
			}
		} else {
			fmt.Println("No changes required")
		}
	}
}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

// func normalize(phone string) string {
// 	var buf bytes.Buffer
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}
// 	return buf.String()
// }
