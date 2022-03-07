package main

import (
	"reflect"
	"testing"
)

func Test_testValidity(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid #1",
			args: args{"23-ab-48-caba-56-haha"},
			want: true,
		},
		{
			name: "invalid #2",
			args: args{"cd-ab-48-caba-56-haha"},
			want: false,
		},
		{
			name: "valid #3",
			args: args{"12-cd"},
			want: true,
		},
		{
			name: "invalid #4",
			args: args{"12-cd-14"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testValidity(tt.args.str); got != tt.want {
				t.Errorf("testValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageNumber(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "test #1",
			args:    args{"2-ab-2-ab"},
			want:    2,
			wantErr: false,
		},
		{
			name:    "test #2",
			args:    args{"2-ab-2-ab-4-ac"},
			want:    2,
			wantErr: false,
		},
		{
			name:    "test #3",
			args:    args{"2-ax-cd"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := averageNumber(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("averageNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("averageNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wholeStory(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test #1",
			args:    args{"1-hello-2-world"},
			want:    "hello world",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wholeStory(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("wholeStory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("wholeStory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storyStats(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name         string
		args         args
		wantShortest string
		wantLongest  string
		wantAverage  int
		wantList     []string
		wantErr      bool
	}{
		{
			name:         "test #1",
			args:         args{"1-hellotest-2-world"},
			wantShortest: "world",
			wantLongest:  "hellotest",
			wantAverage:  7,
			wantList:     nil,
			wantErr:      false,
		},
		{
			name:         "test #2",
			args:         args{"1-hellotest-2-world-3-he-4-hellolongest-5-abcdefg"},
			wantShortest: "he",
			wantLongest:  "hellolongest",
			wantAverage:  7,
			wantList:     []string{"abcdefg"},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShortest, gotLongest, gotAverage, gotList, err := storyStats(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("storyStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotShortest != tt.wantShortest {
				t.Errorf("storyStats() gotShortest = %v, want %v", gotShortest, tt.wantShortest)
			}
			if gotLongest != tt.wantLongest {
				t.Errorf("storyStats() gotLongest = %v, want %v", gotLongest, tt.wantLongest)
			}
			if gotAverage != tt.wantAverage {
				t.Errorf("storyStats() gotAverage = %v, want %v", gotAverage, tt.wantAverage)
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("storyStats() gotList = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}
