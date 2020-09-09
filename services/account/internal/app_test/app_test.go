package app

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/wafuwafu/Hatena-Intern-2020/services/account/app"
	"github.com/wafuwafu/Hatena-Intern-2020/services/account/internal/testutil"
	"github.com/wafuwafu/Hatena-Intern-2020/services/account/repository"
	"github.com/jmoiron/sqlx"
)

func TestMain(m *testing.M) {
	os.Exit(run(m))
}

var testDB *sqlx.DB

func run(m *testing.M) int {
	rand.Seed(time.Now().UnixNano())

	db, err := testutil.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	testDB = db

	return m.Run()
}

func setup() (*app.App, *repository.Repository) {
	a := app.NewApp(testDB)
	repo := repository.NewRepository(testDB)
	return a, repo
}
