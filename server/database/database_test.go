package database

import (
	"os"
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		wantErr  bool
	}{
		{
			name: "valid credentials",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USER":     "postgres",
				"DB_PASSWORD": "password",
				"DB_NAME":     "my_database",
			},
			wantErr: false,
		},
		{
			name: "invalid credentials",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USER":     "wronguser",
				"DB_PASSWORD": "wrongpassword",
				"DB_NAME":     "my_database",
			},
			wantErr: true,
		},
		{
			name: "missing credentials",
			envVars: map[string]string{},
			wantErr: true,
		},
		{
			name: "invalid DSN format",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USER":     "postgres",
				"DB_PASSWORD": "password",
				"DB_NAME":     "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			db, err := ConnectDatabase()

			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && db == nil {
				t.Errorf("ConnectDatabase() returned nil db but no error")
			}
		})
	}
}
