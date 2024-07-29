package database

import (
	"database/sql"
	"testing"
)

func TestDBRepository_SaveCPUMetrics(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		host    string
		appName string
		cpuUse  float32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DB := &DBRepository{
				db: tt.fields.db,
			}
			if err := DB.SaveCPUMetrics(tt.args.host, tt.args.appName, tt.args.cpuUse); (err != nil) != tt.wantErr {
				t.Errorf("DBRepository.SaveCPUMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBRepository_SaveMemMetrics(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		host    string
		appName string
		memUse  float32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DB := &DBRepository{
				db: tt.fields.db,
			}
			if err := DB.SaveMemMetrics(tt.args.host, tt.args.appName, tt.args.memUse); (err != nil) != tt.wantErr {
				t.Errorf("DBRepository.SaveMemMetrics() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
