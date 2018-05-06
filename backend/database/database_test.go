package database

import "testing"

func TestConnect(t *testing.T) {
	_, err := Connect()
	if err != nil {
		t.Errorf("DB connection error: %#v", err)
	}
}
