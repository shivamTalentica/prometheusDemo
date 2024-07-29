package database

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

type sqlOpenFunc func(drivername, datasource string) (*sql.DB, error)

var sqlOpen sqlOpenFunc = sql.Open

func TestInitDB(t *testing.T) {

	mockdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer mockdb.Close()

	tests := []struct {
		mock    func()
		name    string
		want    *DBRepository
		wantErr bool
	}{
		{
			name: "Not Successful Connection",
			mock: func() {
				sqlOpen = func(string, string) (*sql.DB, error) {
					return mockdb, nil
				}
				mock.ExpectPing().WillReturnError(nil)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("InitDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
