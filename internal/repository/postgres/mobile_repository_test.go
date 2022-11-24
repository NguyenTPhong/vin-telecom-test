package postgres

import (
	"reflect"
	"testing"
	"vinbigdata/internal/delivery/http/model"
	"vinbigdata/internal/repository/entity"
	db2 "vinbigdata/test/db"

	"gorm.io/gorm"
)

func TestMobileRepository_SaveUserCall(t *testing.T) {

	type args struct {
		userCall *entity.UserCalls
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "happy case",
			args: args{
				userCall: &entity.UserCalls{
					UserName: "user1",
					Duration: 62000,
					Block:    3,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, cleanFunc := db2.GetTestDB()
			defer cleanFunc()
			r := NewMobileRepository(db)
			if err := r.SaveUserCall(tt.args.userCall); (err != nil) != tt.wantErr {
				t.Errorf("SaveUserCall() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMobileRepository_GetUserCallSummary(t *testing.T) {

	type args struct {
		userName string
	}
	tests := []struct {
		name     string
		args     args
		mockData func(db *gorm.DB)
		want     *model.BillingData
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "happy case, empty value",
			args: args{
				userName: "phongnt",
			},
			mockData: func(db *gorm.DB) {},
			wantErr:  false,
			want:     &model.BillingData{},
		},
		{
			name: "happy case: single record",
			args: args{
				userName: "phongnt",
			},
			mockData: func(db *gorm.DB) {
				db.Create(&entity.UserCalls{
					UserName: "phongnt",
					Block:    2,
					Duration: 32000,
				})
			},
			wantErr: false,
			want: &model.BillingData{
				BlockCount: 2,
				CallCount:  1,
			},
		},
		{
			name: "happy case: multi records",
			args: args{
				userName: "phongnt",
			},
			mockData: func(db *gorm.DB) {
				db.Create(&entity.UserCalls{
					UserName: "phongnt",
					Block:    2,
					Duration: 32000,
				})

				db.Create(&entity.UserCalls{
					UserName: "phongnt",
					Block:    3,
					Duration: 64000,
				})
			},
			wantErr: false,
			want: &model.BillingData{
				BlockCount: 5,
				CallCount:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, cleanFunc := db2.GetTestDB()
			defer cleanFunc()
			r := NewMobileRepository(db)
			tt.mockData(db)
			got, err := r.GetUserCallSummary(tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserCallSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserCallSummary() got = %v, want %v", got, tt.want)
			}
		})
	}
}
