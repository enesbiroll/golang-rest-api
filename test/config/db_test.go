package config

import (
	"rest-api/config" // config paketini import et
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	// Veritabanı bağlantısını kontrol et
	config.Connect() // Connect fonksiyonunu config paketinden çağır

	// Veritabanı bağlantısının null olmadığını doğrula
	assert.NotNil(t, config.DB) // DB'yi config paketinden çağır
}
