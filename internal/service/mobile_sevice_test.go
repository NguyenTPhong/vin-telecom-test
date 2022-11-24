package service

import (
	"fmt"
	"reflect"
	"testing"
	"vinbigdata/internal/delivery/http/model"
	"vinbigdata/internal/repository/entity"
	"vinbigdata/internal/service/mocks"
)

func TestMobileService_SaveUserCall(t *testing.T) {
	type fields struct {
		mobileRepo MobileRepository
	}
	type args struct {
		userName string
		duration int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "happy case: odd block",
			fields: fields{
				mobileRepo: func() MobileRepository {
					mockRepo := &mocks.MobileRepository{}
					// mock result here
					mockRepo.On("SaveUserCall", &entity.UserCalls{
						UserName: "phongnt",
						Duration: 30000,
						Block:    1,
					}).Return(nil)
					return mockRepo
				}(),
			},
			args: args{
				userName: "phongnt",
				duration: 30000,
			},
			wantErr: false,
		},
		{
			name: "happy case: even block",
			fields: fields{
				mobileRepo: func() MobileRepository {
					mockRepo := &mocks.MobileRepository{}
					// mock result here
					mockRepo.On("SaveUserCall", &entity.UserCalls{
						UserName: "phongnt",
						Duration: 32000,
						Block:    2,
					}).Return(nil)
					return mockRepo
				}(),
			},
			args: args{
				userName: "phongnt",
				duration: 32000,
			},
			wantErr: false,
		},
		{
			name: "err case: insert error",
			fields: fields{
				mobileRepo: func() MobileRepository {
					mockRepo := &mocks.MobileRepository{}
					// mock result here
					mockRepo.On("SaveUserCall", &entity.UserCalls{
						UserName: "phongnt",
						Duration: 32000,
						Block:    2,
					}).Return(fmt.Errorf("error occured"))
					return mockRepo
				}(),
			},
			args: args{
				userName: "phongnt",
				duration: 32000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MobileService{
				mobileRepo: tt.fields.mobileRepo,
			}
			if err := s.SaveUserCall(tt.args.userName, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("SaveUserCall() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMobileService_GetUserBilling(t *testing.T) {
	type fields struct {
		mobileRepo MobileRepository
	}
	type args struct {
		userName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.BillingData
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "happy case",
			fields: fields{
				mobileRepo: func() MobileRepository {
					mockRepo := &mocks.MobileRepository{}
					// mock result here
					mockRepo.On("GetUserCallSummary", "phongnt").Return(&model.BillingData{
						BlockCount: 1,
						CallCount:  1,
					}, nil)
					return mockRepo
				}(),
			},
			args: args{
				userName: "phongnt",
			},
			wantErr: false,
			want: &model.BillingData{
				BlockCount: 1,
				CallCount:  1,
			},
		},
		{
			name: "error case",
			fields: fields{
				mobileRepo: func() MobileRepository {
					mockRepo := &mocks.MobileRepository{}
					// mock result here
					mockRepo.On("GetUserCallSummary", "phongnt").Return(nil, fmt.Errorf("error occured"))
					return mockRepo
				}(),
			},
			args: args{
				userName: "phongnt",
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MobileService{
				mobileRepo: tt.fields.mobileRepo,
			}
			got, err := s.GetUserBilling(tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserBilling() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserBilling() got = %v, want %v", got, tt.want)
			}
		})
	}
}
