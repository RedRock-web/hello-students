// @program: hello-students
// @author: edte
// @create: 2020-08-11 21:22
package service

import (
	"reflect"
	"testing"
)

func TestGetCanteenPictures(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "error1",
			args: args{
				name: "樱花食堂",
			},
			want: []string{
				"http://cdn.redrock.team/hello-student_yhst0.jpg",
				"http://cdn.redrock.team/hello-student_yhst1.jpg",
				"http://cdn.redrock.team/hello-student_yhst2.jpg",
				"http://cdn.redrock.team/hello-student_yhst3.jpg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCanteenPictures(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCanteenPictures() = %v, want %v", got, tt.want)
			}
		})
	}
}
