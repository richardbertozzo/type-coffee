package entity

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	validLink, _ := NewImageLink("https://rollingstone.uol.com.br/media/_versions/godzilla-kingking-reprod-twitter-cortada_widelg.jpg")
	caracs, _ := NewCaracteristics([]string{"fraco"})

	caracsZero, _ := NewCaracteristics([]string{})

	type args struct {
		uuid string
		name string
		d    string
		l    Link
		c    []Caracteristic
	}
	tests := []struct {
		name    string
		args    args
		want    Coffee
		wantErr bool
	}{
		{
			name: "Create a valid coffee",
			args: args{
				uuid: "1234",
				name: "Café iguaçu",
				d:    "Café lavado",
				l:    validLink,
				c:    caracs,
			},
			want: Coffee{
				ID:             "1234",
				Name:           "Café iguaçu",
				Description:    "Café lavado",
				Image:          validLink,
				caracteristics: caracs,
			},
			wantErr: false,
		},
		{
			name: "Create a invalid coffee",
			args: args{
				uuid: "1234",
				name: "Café iguaçu",
				d:    "Café lavado",
				l:    validLink,
				c:    caracsZero,
			},
			want:    Coffee{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.uuid, tt.args.name, tt.args.d, tt.args.l, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
