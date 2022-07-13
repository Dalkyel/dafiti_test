package main

import (
	"reflect"
	"testing"
)

// TestMain_IsStraight Tests Different hand results
func TestMain_IsStraight(t *testing.T) {
	type args struct {
		hand []int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "results1",
			args: args{
				hand: []int{2, 3, 4, 5, 6},
			},
			want:    "Congrats!! you have a valid hand",
			wantErr: false,
		},
		{
			name: "results2",
			args: args{
				hand: []int{2, 7, 8, 5, 10, 9, 11},
			},
			want:    "Congrats!! you have a valid hand",
			wantErr: false,
		},
		{
			name: "results3",
			args: args{
				hand: []int{11, 12, 13, 14, 2},
			},
			want:    "Congrats!! you have a valid hand",
			wantErr: false,
		},
		{
			name: "results4",
			args: args{
				hand: []int{14, 5, 4, 2, 3},
			},
			want:    "Congrats!! you have a valid hand",
			wantErr: false,
		},
		{
			name: "results5",
			args: args{
				hand: []int{7, 3, 2},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "results6",
			args: args{
				hand: []int{7, 11, 2, 4, 1, 5, 7, 15},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "results7",
			args: args{
				hand: []int{7, 7, 12, 11, 3, 4, 14},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsStraight(tt.args.hand)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsStraight() = \n got  %v \n want %v", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("IsStraight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
