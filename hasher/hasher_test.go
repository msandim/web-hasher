package hasher

import "testing"

func TestHTTPHasher(t *testing.T) {

	mockClient := newMockClient(func(string) string {
		return "example body from request to google"
	})
	defer mockClient.Close()

	hasher := NewHTTPHasher(mockClient)
	hash, err := hasher.Hash("google.pt")

	if err != nil {
		t.Fatal("Unexpected error hashing")
	}

	if hash != "419a2a2a5fb0adf7ebdcf48072beeacd" {
		t.Fatalf("Wrong hash returned. Got %s instead", hash)
	}
}

func TestMockHasher(t *testing.T) {

	mockHasher := NewMockHasher(map[string]string{"abc.com": "hash_xpto"})
	hash, err := mockHasher.Hash("abc.com")

	if err != nil {
		t.Fatalf("Not expected error received: %s", err)
	}

	if hash != "hash_xpto" {
		t.Fatalf("Invalid hash. Got %s", hash)
	}
}
