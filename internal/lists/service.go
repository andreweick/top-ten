package lists

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"

	"filippo.io/age"
	"filippo.io/age/armor"
)

type Service struct {
	collection TopTenCollection
}

func NewService(ctx context.Context) (*Service, error) {
	encryptedData := GetEmbeddedData()

	// Get the password from environment variable
	password := os.Getenv("AGE_ENCRYPTION_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("AGE_ENCRYPTION_PASSWORD environment variable is not set")
	}

	// Create age identity for decryption
	identity, err := age.NewScryptIdentity(password)
	if err != nil {
		return nil, fmt.Errorf("failed to create age identity: %w", err)
	}

	// Check if the data is armored (ASCII format)
	var ageReader io.Reader
	if bytes.HasPrefix(encryptedData, []byte("-----BEGIN AGE ENCRYPTED FILE-----")) {
		// It's armored, decode it first
		armorReader := armor.NewReader(bytes.NewReader(encryptedData))
		ageReader = armorReader
	} else {
		// It's binary format
		ageReader = bytes.NewReader(encryptedData)
	}

	// Decrypt the age data
	reader, err := age.Decrypt(ageReader, identity)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt age data: %w", err)
	}

	// Read the decrypted JSON data
	decryptedData, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read decrypted data: %w", err)
	}

	var collection TopTenCollection
	if err := json.Unmarshal(decryptedData, &collection); err != nil {
		return nil, fmt.Errorf("failed to unmarshal top ten data: %w", err)
	}

	if len(collection.Lists) == 0 {
		return nil, fmt.Errorf("no top ten lists found in data")
	}

	return &Service{
		collection: collection,
	}, nil
}

func (s *Service) GetRandomList() (TopTenList, error) {
	if len(s.collection.Lists) == 0 {
		return TopTenList{}, fmt.Errorf("no lists available")
	}

	max := big.NewInt(int64(len(s.collection.Lists)))
	randomIndex, err := rand.Int(rand.Reader, max)
	if err != nil {
		return TopTenList{}, fmt.Errorf("failed to generate random number: %w", err)
	}

	return s.collection.Lists[randomIndex.Int64()], nil
}

func (s *Service) GetListCount() int {
	return len(s.collection.Lists)
}
