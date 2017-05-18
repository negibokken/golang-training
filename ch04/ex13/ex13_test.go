package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestGetPoster(t *testing.T) {
	type args struct {
		posterURL string
	}
	tests := []struct {
		name    string
		args    args
		wantO   string
		wantErr bool
	}{
		{
			"correct poster",
			args{
				"https://images-na.ssl-images-amazon.com/images/M/MV5BMDdmZGU3NDQtY2E5My00ZTliLWIzOTUtMTY4ZGI1YjdiNjk3XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg",
			},
			"peko",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &bytes.Buffer{}
			if err := GetPoster(o, tt.args.posterURL); (err != nil) != tt.wantErr {
				t.Errorf("GetPoster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Cannot check jpg file
			// if gotO := o.String(); gotO != tt.wantO {
			// 	t.Errorf("GetPoster() = %v, want %v", gotO, tt.wantO)
			// }
		})
	}
}

func TestGetMovie(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		want    *Movie
		wantErr bool
	}{
		{
			"Get correct moview",
			args{"titanic"},
			&Movie{
				"Titanic",
				"https://images-na.ssl-images-amazon.com/images/M/MV5BMDdmZGU3NDQtY2E5My00ZTliLWIzOTUtMTY4ZGI1YjdiNjk3XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMovie(tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMovie() = %v, want %v", got, tt.want)
			}
		})
	}
}
