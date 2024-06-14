package logger

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"log"
	"testing"
)

type s struct {
	a string
	b struct{ c map[uuid.UUID]s }
}

func TestLogger(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected r got nil")
		} else {
			t.Log(r)
		}
	}()
	var obj s
	err := faker.FakeData(&obj)
	if err != nil {
		t.Fatal(err)
	}
	var objSlice []s
	err = faker.FakeData(&objSlice)
	if err != nil {
		t.Fatal(err)
	}
	var objMap map[string]s
	err = faker.FakeData(&objMap)
	if err != nil {
		t.Fatal(err)
	}
	Info("Info", obj, "obj")
	Debug("Debug", obj, objSlice, objMap, "obj")
	Warn("Warn", obj, "obj")
	Imp("Imp", obj, "obj")
	Err("Err", obj, "obj")
	Fatal("ok")
	log.Println("Info", obj, "ms")
}
