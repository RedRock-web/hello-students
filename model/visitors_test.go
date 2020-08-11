// @program: hello-students
// @author: edte
// @create: 2020-08-11 21:09
package model

import (
	"testing"
)

func TestAddVisitor(t *testing.T) {
	type args struct {
		visitor Visitor
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常1",
			args: args{
				visitor: Visitor{
					IP:     "127.0.0.1",
					Gender: "女",
					Region: "ha",
					Major:  "test",
				},
			},
			wantErr: false,
		},
		{
			name: "正常2",
			args: args{
				visitor: Visitor{
					IP:     "127.0.0.1",
					Gender: "男",
					Region: "hsdfa",
					Major:  "test",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitModel()
			if err := AddVisitor(tt.args.visitor); (err != nil) != tt.wantErr {
				t.Errorf("AddVisitor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetVisitorsNumber(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{
			name:    "正常",
			want:    71,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		InitModel()
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVisitorsNumber()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVisitorsNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetVisitorsNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}
