package singleton

import (
    "database/sql"
    "os"
    "fmt"
    "sync"
    "github.com/mattn/go-sqlite3"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
	 if singleInstance == nil {
             fmt.Println("Creating single instance now.")
	     singleInstance = &single{}
    } else {
	fmt.Println("Single instance already created.")
    }


    return singleInstance
}
 
func  Open_s() {
    db, er :=sql.Open("sqlite3", "table.db")
    if er !=nil {
	    fmt.Fatal("Connection failed")
    }
    st.db = db
    fmt.Println("Connection opened")
}

func Close_s(){
    if st.db ==nil{
	fmt.Fatal("DB was not opened before")
     }
	st.db.Close()
	fmt.Println("Connection closed")
}

func  Query_s(db *sql.DB, Id strinf, Name string, Surname string) {
    if st.db == nil {
	log.Fatal ("Open connection before start")
	}
	a , er := st.db.Query(query)
	if er !=nil {
		return er
	}

	return a, er
}

func Create () {
	file, er := os.Create("sqlite-database.db")
	if er != nil {
		log.Fatal("Was created before")
	}
	file.Close()
	fmt.Println("Db was created now")
	
