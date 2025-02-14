package e2e

import (
	"bytes"
	"encoding/json"
	"io"
	"merch_store/internal/api"
	"merch_store/internal/api/dto"
	"merch_store/internal/config"
	"merch_store/pkg/pb"
	"merch_store/tests/env"
	"net/http"
	"net/http/httptest"
	"testing"
)

const MODE = "test"

func TestAuthSuccess(t *testing.T) {
	conf := config.LoadConfig("../../configs", MODE)
	db, err := env.InitTestDb(conf.Database.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	db.Up()
	defer db.Close()
	defer db.Down()

	app := api.NewApp(&api.AppDeps{
		Config: conf,
	})
	router, err := app.Build()
	if err != nil {
		t.Fatal(err)
	}
	ts := httptest.NewServer(router)
	defer ts.Close()

	data, _ := json.Marshal(&dto.AuthReq{
		Username: "test",
		Password: "123456",
	})
	res, err := http.Post(ts.URL+"/auth", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("expected %d got %d", 200, res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData pb.AuthRes
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("token empty")
	}
}
func TestAuthFail(t *testing.T) {
	conf := config.LoadConfig("../../configs", MODE)
	db, err := env.InitTestDb(conf.Database.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	db.Up()
	defer db.Close()
	defer db.Down()
	_, err = db.Exec(`INSERT INTO users (name, password) VALUES ($1, $2)`, "test", "123456")
	if err != nil {
		t.Fatal(err)
	}
	app := api.NewApp(&api.AppDeps{
		Config: conf,
	})
	router, err := app.Build()
	if err != nil {
		t.Fatal(err)
	}
	ts := httptest.NewServer(router)
	defer ts.Close()
	data, _ := json.Marshal(&dto.AuthReq{
		Username: "test",
		Password: "654321",
	})
	res, err := http.Post(ts.URL+"/auth", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 400 {
		t.Fatalf("expected %d got %d", 400, res.StatusCode)
	}
}
