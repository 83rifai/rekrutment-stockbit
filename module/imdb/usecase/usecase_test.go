package usecase

import (
	"reflect"
	"testing"
)

func TestGetList(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetList()
		})
	}
}

func TestGetDetail(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		args     args
		wantResp interface{}
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetDetail(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetDetail() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
