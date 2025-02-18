package e2e

import (
	"bytes"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"merch_store/internal/api/dto"
	"merch_store/internal/models"
	"merch_store/pkg/jwt"
	"merch_store/tests/env"
	"net/http"
	"testing"
	"time"
)

func TestSendCoins_Success(t *testing.T) {
	e := env.NewEnv()
	db, err := env.InitTestDb(e.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	db.Up()
	defer db.Close()
	defer db.Down()

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	user1 := models.User{
		Name:     "test1",
		Password: string(password),
	}
	_, err = db.Exec(`INSERT INTO users (name, password) VALUES ($1, $2)`, user1.Name, user1.Password)
	if err != nil {
		t.Fatal(err)
	}

	password, err = bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	user2 := models.User{
		Name:     "test2",
		Password: string(password),
	}
	_, err = db.Exec(`INSERT INTO users (name, password) VALUES ($1, $2)`, user2.Name, user2.Password)
	if err != nil {
		t.Fatal(err)
	}

	dataDto := dto.SendCoinsReq{
		Amount: 100,
		ToUser: user2.Name,
	}

	body, err := json.Marshal(dataDto)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", e.ApiAddress+"/sendCoins", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	token, err := jwt.NewJWT(e.Jwt).Create(jwt.Data{
		Name: user1.Name,
	}, time.Now().Add(time.Minute*2))
	if err != nil {
		t.Fatal()
	}

	req.Header.Add("Authorization", "Bearer "+token)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("expected %d got %d", 200, res.StatusCode)
	}
}

func TestSendCoins_BadUserName(t *testing.T) {
	e := env.NewEnv()
	db, err := env.InitTestDb(e.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	db.Up()
	defer db.Close()
	defer db.Down()

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	user1 := models.User{
		Name:     "test1",
		Password: string(password),
	}
	var balance1 int64
	err = db.QueryRow(`INSERT INTO users (name, password) VALUES ($1, $2) RETURNING balance`, user1.Name, user1.Password).Scan(&balance1)
	if err != nil {
		t.Fatal(err)
	}

	dataDto := dto.SendCoinsReq{
		Amount: 100,
		ToUser: "bad name",
	}

	body, err := json.Marshal(dataDto)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", e.ApiAddress+"/sendCoins", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	token, err := jwt.NewJWT(e.Jwt).Create(jwt.Data{
		Name: user1.Name,
	}, time.Now().Add(time.Minute*2))
	if err != nil {
		t.Fatal()
	}
	req.Header.Add("Authorization", "Bearer "+token)

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 400 {
		t.Fatalf("expected %d got %d", 400, res.StatusCode)
	}

}

func TestSendCoins_BadToken(t *testing.T) {
	e := env.NewEnv()

	dataDto := dto.SendCoinsReq{
		Amount: 100,
		ToUser: "test",
	}

	body, err := json.Marshal(dataDto)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", e.ApiAddress+"/sendCoins", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+"bad token")

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("expected %d got %d", 401, res.StatusCode)
	}

}

func TestSendCoins_BadBalance(t *testing.T) {
	e := env.NewEnv()
	db, err := env.InitTestDb(e.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	db.Up()
	defer db.Close()
	defer db.Down()

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	user1 := models.User{
		Name:     "test1",
		Password: string(password),
	}
	_, err = db.Exec(`INSERT INTO users (name, password, balance) VALUES ($1, $2, $3) RETURNING balance`, user1.Name, user1.Password, 10)
	if err != nil {
		t.Fatal(err)
	}

	password, err = bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	user2 := models.User{
		Name:     "test2",
		Password: string(password),
	}
	_, err = db.Exec(`INSERT INTO users (name, password) VALUES ($1, $2) RETURNING balance`, user2.Name, user2.Password)
	if err != nil {
		t.Fatal(err)
	}

	dataDto := dto.SendCoinsReq{
		Amount: 100,
		ToUser: user2.Name,
	}

	body, err := json.Marshal(dataDto)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", e.ApiAddress+"/sendCoins", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	token, err := jwt.NewJWT(e.Jwt).Create(jwt.Data{
		Name: user1.Name,
	}, time.Now().Add(time.Minute*2))
	if err != nil {
		t.Fatal()
	}

	req.Header.Add("Authorization", "Bearer "+token)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 400 {
		t.Fatalf("expected %d got %d", 400, res.StatusCode)
	}
}
