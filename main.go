package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"runtime/pprof"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func BizLogic1(ctx context.Context, in SomeReq1) (out SomeRes1, err error) {
	out = SomeRes1{}
	out.A = in.A
	out.B = in.B
	out.C = in.C
	in.Src = "bl1"
	err = db.Create(in).Error
	return
}

func BizLogic2(ctx context.Context, rawIn interface{}) (rawOut interface{}, err error) {

	//Value added?
	in, ok := rawIn.(*SomeReq1)
	if !ok {
		err = errors.New("cant cast")
		return
	}
	out := &SomeRes1{}
	//
	out.A = in.A
	out.B = in.B
	out.C = in.C
	rawOut = out
	in.Src = "bl2"
	err = db.Create(in).Error
	return
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db?_journal=wal"), &gorm.Config{})
	// db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared&_journal=memory"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&SomeReq1{})

	mx := http.NewServeMux()

	s := http.Server{
		Addr:    ":8080",
		Handler: mx,
	}
	mx.HandleFunc("/pprof/start", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile("pprof", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
	})
	mx.HandleFunc("/pprof/stop", func(w http.ResponseWriter, r *http.Request) {
		pprof.StopCPUProfile()
	})
	mx.HandleFunc("/gen/bl1", AdaptGen(BizLogic1))
	//http.HandleFunc("/int/bl1",AdaptInt(&SomeReq1{},BizLogic1)) <= ERROR
	mx.HandleFunc("/int/bl2", AdaptInt(SomeReq1{}, BizLogic2))
	mx.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
		go s.Shutdown(context.Background())
	})

	log.Printf("Listening")
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

}
