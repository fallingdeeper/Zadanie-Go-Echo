package singleton

import (
    "database/sql"
    "os"
    "fmt"
    "sync"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

var lock = &sync.Mutex{}

type singleton struct {
	single	*sql.DB
}

var singleInstance *singleton

func GetInstance() *singleton {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
	 if singleInstance == nil {
             fmt.Println("Creating single instance now.")
	     singleInstance = &singleton{}
	 } else {
	     fmt.Println("Single instance already created.")
	 }
    } else {
	fmt.Println("Single instance already created.")
    }


    return singleInstance
}
 
func (st *singleton)  OpenS() {
    single, er :=sql.Open("sqlite3", "./table.db")
    if er !=nil {
	    log.Fatal("Connection failed")
    }
    st.single = single
    fmt.Println("Connection opened")
}

func (st *singleton) CloseS(){
    if st.single ==nil{
	log.Fatal("DB was not opened before")
     }
	st.single.Close()
	fmt.Println("Connection closed")
}

func (st *singleton) QueryS(ql string) (*sql.Rows, error) {
    if st.single == nil {
	log.Fatal ("Open connection before start")
	}
	a , er := st.single.Query(ql)
	if er !=nil {
		return nil, er
	}

	return a, er
}

func (st *singleton) Create () {
	file, er := os.Create("sqlite-database.db")
	if er != nil {
		log.Fatal("Was created before")
	}
	file.Close()
	fmt.Println("Db was created now")
}
