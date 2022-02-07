package manifest

import (
	"encoding/json"
	"fmt"
	"os"
)

type Manifest map[string]string

func Read(file string) (Manifest, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("reading asset manifest: %w", err)
	}

	var mfest Manifest
	err = json.Unmarshal(bytes, &mfest)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling asset manifest: %w", err)
	}

	return mfest, nil
}

func Write(mfest Manifest, file string) error {
	contents, err := json.MarshalIndent(mfest, "", "\t")
	if err != nil {
		return fmt.Errorf("marshalling asset manifest: %w", err)
	}

	err = os.WriteFile(file, contents, 0644)
	if err != nil {
		return fmt.Errorf("writing asset manifest: %w", err)
	}

	return nil
}
