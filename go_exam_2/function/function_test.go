package function

import "testing"

func TestValidateDigitThailandCitizenID(t *testing.T) {
	type args struct {
		idNo string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test CitizenIDLength 12 digits must be fail",
			args: args{
				idNo: "123456789012",
			},
			wantErr: true,
		},
		{
			name: "Test CitizenIDLength 13 digits must be success",
			args: args{
				idNo: "1516712728928",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateThailandCitizenID(tt.args.idNo); (err != nil) != tt.wantErr {
				t.Errorf("ValidateThailandCitizenID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNumericThailandCitizenID(t *testing.T) {
	type args struct {
		idNo string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test CitizenID with charecter must be fail",
			args: args{
				idNo: "a12233445566b",
			},
			wantErr: true,
		},
		{
			name: "Test CitizenID numeric only must be success",
			args: args{
				idNo: "1516712728928",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateThailandCitizenID(tt.args.idNo); (err != nil) != tt.wantErr {
				t.Errorf("ValidateThailandCitizenID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePatternThailandCitizenID(t *testing.T) {
	type args struct {
		idNo string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test wrong pattern must be fail",
			args: args{
				idNo: "1516712728912",
			},
			wantErr: true,
		},
		{
			name: "Test true pattern must be success",
			args: args{
				idNo: "1516712728928",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateThailandCitizenID(tt.args.idNo); (err != nil) != tt.wantErr {
				t.Errorf("ValidateThailandCitizenID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
