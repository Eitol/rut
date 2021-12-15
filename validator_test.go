package rut

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		rut string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "VALID",
			args: args{
				rut: "6473297-8",
			},
			wantErr: nil,
		},
		{
			name: "ErrInvalidVD",
			args: args{
				rut: "6473297-7",
			},
			wantErr: ErrInvalidVD,
		},
		{
			name: "ErrToShort",
			args: args{
				rut: "6",
			},
			wantErr: ErrToShort,
		},
		{
			name: "ErrToLong",
			args: args{
				rut: "11111111111111111111",
			},
			wantErr: ErrToLong,
		},
		{
			name: "ErrInvalidNumber",
			args: args{
				rut: "f473297-7",
			},
			wantErr: ErrInvalidNumber,
		},
		{
			name: "valid_with_spaces",
			args: args{
				rut: " 6473297-8 ",
			},
			wantErr: nil,
		},
		{
			name: "valid_with_dots",
			args: args{
				rut: "6.473.297-8",
			},
			wantErr: nil,
		},
		{
			name: "valid_with_dots_and_spaces",
			args: args{
				rut: "6.473.297-8",
			},
			wantErr: nil,
		},
		{
			name: "valid_with_dots_and_spaces_and_lowercase",
			args: args{
				rut: "26.349.413-k",
			},
			wantErr: nil,
		},
		{
			name: "valid_without_dots__lowercase",
			args: args{
				rut: "26349413k",
			},
			wantErr: nil,
		},
		{
			name: "valid_without_dots__lowercase",
			args: args{
				rut: "26349413K",
			},
			wantErr: nil,
		},
		{
			name: "valid_without_dots__lowercase",
			args: args{
				rut: "263494130",
			},
			wantErr: ErrInvalidVD,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.args.rut)
			if err == nil && tt.wantErr != nil || err != nil && tt.wantErr == nil {
				t.Fatalf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Fatalf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
