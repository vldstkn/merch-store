package unit

import (
	"merch_store/pkg/jwt"
	"testing"
	"time"
)

func TestJwt_Success(t *testing.T) {
	name := "test"
	j := jwt.NewJWT("123456")
	token, err := j.Create(jwt.Data{Name: name}, time.Now().Add(time.Minute*2))
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := j.Parse(token)
	if !isValid {
		t.Fatal("token is invalid")
	}
	if data.Name != name {
		t.Fatalf("expected %s got %s", name, data.Name)
	}
}
func TestJwt_Fail(t *testing.T) {
	name := "test"
	j := jwt.NewJWT("123456")
	token, err := j.Create(jwt.Data{Name: name}, time.Now().Add(time.Minute*-2))
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := j.Parse(token)
	if isValid {
		t.Fatal("token is valid")
	}
	if data != nil {
		t.Fatal("data not empty")
	}
}
