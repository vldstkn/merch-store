package unit

import (
	"bytes"
	"encoding/json"
	"merch_store/pkg/req"
	"net/http"
	"testing"
)

func TestReq_Success(t *testing.T) {
	type Data struct {
		Name   string `json:"name" validate:"required"`
		Number int    `json:"password" validate:"required,number"`
	}
	data := Data{
		Name:   "test",
		Number: 123456,
	}
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("POST", "test", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	newData, err := req.HandleBody[Data](request)
	if err != nil {
		t.Fatal(err)
	}
	if newData.Number != data.Number {
		t.Fatalf("number: expected %d got %d", data.Number, newData.Number)
	}
	if newData.Name != data.Name {
		t.Fatalf("name: expected %s got %s", data.Name, newData.Name)
	}
}
func TestReq_Fail(t *testing.T) {
	type Data struct {
		Name   string `json:"name" validate:"required"`
		Number int    `json:"password" validate:"required,number"`
	}
	data := struct {
		Number string
	}{
		Number: "test",
	}
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("POST", "test", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	_, err = req.HandleBody[Data](request)
	if err == nil {
		t.Fatal("expected error")
	}
}
