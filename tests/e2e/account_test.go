package e2e

import (
	"bytes"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"merch_store/internal/api/dto"
	"merch_store/internal/models"
	"merch_store/pkg/jwt"
	"merch_store/pkg/pb"
	"merch_store/tests/env"
	"net/http"
	"testing"
	"time"
)

func TestAuthSuccess(t *testing.T) {
	e := env.NewEnv()
	db, err := env.InitTestDb(e.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Up(); err != nil {
		t.Fatal(err)
	}

	defer db.Close()
	defer db.Down()

	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(&dto.AuthReq{
		Username: "test",
		Password: "123456",
	})
	res, err := http.Post(e.ApiAddress+"/auth", "application/json", bytes.NewReader(data))
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
	e := env.NewEnv()
	db, err := env.InitTestDb(e.Dsn)
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
	if err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(&dto.AuthReq{
		Username: "test",
		Password: "654321",
	})
	res, err := http.Post(e.ApiAddress+"/auth", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 400 {
		t.Fatalf("expected %d got %d", 400, res.StatusCode)
	}
}

func TestInfo_Success(t *testing.T) {
	e := env.NewEnv()
	db, err := env.InitTestDb(e.Dsn)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Up(); err != nil {
		t.Fatal(err)
	}

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

	_, err = db.Exec(`INSERT INTO users (name, password, balance) VALUES ($1, $2, 890)`, user1.Name, user1.Password)
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

	product := models.Product{
		Type:  "cup",
		Price: 100,
	}

	_, err = db.Exec(`INSERT INTO products (type, price) VALUES ($1, $2)`, product.Type, product.Price)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO products_users (user_name, product_type, quantity) VALUES ($1, $2, 3)`, user1.Name, product.Type)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO transfers (from_user_name, to_user_name, amount) VALUES ($1, $2, 10)`, user1.Name, user2.Name)
	req, err := http.NewRequest("GET", e.ApiAddress+"/info", nil)

	token, err := jwt.NewJWT(e.Jwt).Create(jwt.Data{
		Name: user1.Name,
	}, time.Now().Add(time.Minute*2))
	if err != nil {
		t.Fatal(err)
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var data dto.GetInfoRes
	err = json.Unmarshal(body, &data)
	if err != nil {
		t.Fatal(err)
	}
	expectedData := dto.GetInfoRes{
		Inventory: []dto.Inventory{
			{
				Type:     "cup",
				Quantity: 3,
			},
		},
		Coins: 890,
		CoinsHistory: dto.CoinsHistory{
			Sent: []dto.Sent{
				{
					ToUser: user2.Name,
					Amount: 10,
				},
			},
			Received: []dto.Received{},
		},
	}
	if expectedData.Coins != data.Coins {
		t.Fatalf("coins: expected %d got %d", expectedData.Coins, data.Coins)
	}

	if len(data.Inventory) != len(expectedData.Inventory) || expectedData.Inventory[0] != data.Inventory[0] {
		t.Fatalf("inventory: expected %+v got %+v", expectedData.Inventory, data.Inventory)
	}
	if len(data.CoinsHistory.Sent) != len(expectedData.CoinsHistory.Sent) || expectedData.CoinsHistory.Sent[0] != data.CoinsHistory.Sent[0] {
		t.Fatalf("sent: expected %+v got %+v", expectedData.CoinsHistory.Sent, data.CoinsHistory.Sent)
	}
	if len(data.CoinsHistory.Received) != len(expectedData.CoinsHistory.Received) {
		t.Fatalf("received: expected empty got %+v", data.CoinsHistory.Received)
	}
}
