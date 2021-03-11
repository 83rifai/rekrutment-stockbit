package usecase

import (
	"reflect"
	"stockbit/module/anagram/model"
	"testing"
)

func TestUseCaseAnagram(t *testing.T) {
	type args struct {
		b *model.Anagram
	}
	tests := []struct {
		name    string
		args    args
		wantRs  interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRs, err := UseCaseAnagram(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCaseAnagram() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRs, tt.wantRs) {
				t.Errorf("UseCaseAnagram() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}

func Test_sortStr(t *testing.T) {
	type args struct {
		k string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortStr(tt.args.k); got != tt.want {
				t.Errorf("sortStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
