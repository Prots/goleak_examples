package main

import (
	"testing"
	"github.com/Prots/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestLoadCache(t *testing.T){
	err := loadCache()
	if err != nil {
		t.Fatalf("Test failed, err: %v", err)
	}
}

func TestLoadCacheProperly(t *testing.T){
	err := loadCacheProperly()
	if err != nil {
		t.Fatalf("Test failed, err: %v", err)
	}
}