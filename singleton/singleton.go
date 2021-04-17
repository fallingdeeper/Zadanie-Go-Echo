package singleton

import (
    "database/sql"
    "os"
    "fmt"
    "sync"
    "github.com/mattn/go-sqlite3"
    "./db/table.db"
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
    singleton, er :=sql.Open("sqlite3", "./db/table.db")
    if er !=nil {
	    fmt.Fatal("Connection failed")
    }
    st.singleton = singleton
    fmt.Println("Connection opened")
}

func Close_s(){
    if st.singleton ==nil{
	fmt.Fatal("DB was not opened before")
     }
	st.singleton.Close()
	fmt.Println("Connection closed")
}

func  Query_s(db *sql.DB, Id strinf, Name string, Surname string) {
    if st.singleton == nil {
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
	
