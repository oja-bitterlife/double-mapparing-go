package main

import (
	"encoding/json"
	"fmt"

	doublemapparing "github.com/oja-bitterlife/double-mapparing-go"
)

// テスト用のデータ構造（ネストしたもの）
type Config struct {
	AppID   string   `json:"app_id"`
	Version int      `json:"version"`
	Meta    Metadata `json:"meta"`
}

type Metadata struct {
	Owner string `json:"owner"`
}

func main() {
	// New を使って一撃で初期化
	db := doublemapparing.New[Config](
		func(v any) ([]byte, error) { return json.Marshal(v) },
		func(b []byte, v any) error { return json.Unmarshal(b, v) },
	)

	// あとは Update や Raw を呼ぶだけ
	db.Update(func(cfg *Config) error {
		cfg.Version = 2
		return nil
	})

	fmt.Printf("Current Version: %d\n", db.Raw().Version)
}
