package e2e

import (
	"golang.org/x/crypto/bcrypt"
	"merch_store/internal/models"
	"merch_store/pkg/jwt"
	"merch_store/tests/env"
	"net/http"
	"testing"
	"time"
)

func TestBuy_Success(t *testing.T) {
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

	user := models.User{
		Name:     "test",
		Password: string(password),
	}
	var balance int64
	err = db.QueryRow(`INSERT INTO users (name, password) VALUES ($1, $2) RETURNING balance`, user.Name, user.Password).Scan(&balance)
	if err != nil {
		t.Fatal(err)
	}

	token, err := jwt.NewJWT(e.Jwt).Create(jwt.Data{
		Name: user.Name,
	}, time.Now().Add(time.Minute*2))
	if err != nil {
		t.Fatal()
	}

	product := models.Product{
		Type:  "cup",
		Price: 100,
	}

	_, err = db.Exec(`INSERT INTO products (type, price) VALUES ($1, $2)`, product.Type, product.Price)
	if err != nil {
		t.Fatal()
	}

	req, err := http.NewRequest("GET", e.ApiAddress+"/buy/cup", nil)
	if err != nil {
		t.Fatal()
	}
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("expected %d got %d", 200, res.StatusCode)
	}
}

func TestBuy_InvalidToken(t *testing.T) {
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

	user := models.User{
		Name:     "test",
		Password: string(password),
	}

	_, err = db.Exec(`INSERT INTO users (name, password) VALUES ($1, $2) RETURNING balance`, user.Name, user.Password)
	if err != nil {
		t.Fatal()
	}

	req, err := http.NewRequest("GET", e.ApiAddress+"/buy/test", nil)
	if err != nil {
		t.Fatal()
	}

	req.Header.Add("Authorization", "Bearer "+"bad token")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 401 {
		t.Fatalf("expected %d got %d", 401, res.StatusCode)
	}
}

func TestBuy_ProductNotFound(t *testing.T) {
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

	user := models.User{
		Name:     "test",
		Password: string(password),
	}

	_, err = db.Exec(`INSERT INTO users (name, password) VALUES ($1, $2)`, user.Name, user.Password)
	if err != nil {
		t.Fatal()
	}

	token, err := jwt.NewJWT(e.Jwt).Create(jwt.Data{
		Name: user.Name,
	}, time.Now().Add(time.Minute*2))
	if err != nil {
		t.Fatal()
	}

	req, err := http.NewRequest("GET", e.ApiAddress+"/buy/bad_product", nil)
	if err != nil {
		t.Fatal()
	}

	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 400 {
		t.Fatalf("expected %d got %d", 400, res.StatusCode)
	}
}
