package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"vinbigdata/internal/delivery/http/controller/mocks"
	"vinbigdata/internal/delivery/http/model"
	error2 "vinbigdata/package/error"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestTelecomController_SaveCall(t *testing.T) {
	type fields struct {
		mobilService MobileService
	}
	type args struct {
		mockCtx func(ctx *gin.Context)
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedResponse string
		expectedStatus   int
	}{
		// TODO: Add test cases.
		{
			name: "error case: invalid payload",
			fields: fields{
				mobilService: func() MobileService {
					mockService := &mocks.MobileService{}
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					ctx.Request = httptest.NewRequest(
						http.MethodPost,
						"/mobile/phongnt/call",
						strings.NewReader(""),
					)
				},
			},
			expectedResponse: "{\"code\":400,\"message\":\"unexpected end of JSON input\"}",
			expectedStatus:   http.StatusBadRequest,
		},
		{
			name: "error case: service return err",
			fields: fields{
				mobilService: func() MobileService {
					mockService := &mocks.MobileService{}
					mockService.On("SaveUserCall", "phongnt", int64(30000)).Return(error2.NewCError(500, "error occurred"))
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					payload := model.CallRequest{
						Duration: 30000,
					}
					bodyString, _ := json.Marshal(payload)
					ctx.Request = httptest.NewRequest(
						http.MethodPost,
						"/mobile/phongnt/call",
						strings.NewReader(string(bodyString)),
					)
					ctx.AddParam("user_name", "phongnt")
				},
			},
			expectedResponse: "{\"code\":500,\"message\":\"error occurred\"}",
			expectedStatus:   http.StatusBadRequest,
		},
		{
			name: "happy case: case",
			fields: fields{
				mobilService: func() MobileService {
					mockService := &mocks.MobileService{}
					mockService.On("SaveUserCall", "phongnt", int64(30000)).Return(nil)
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					payload := model.CallRequest{
						Duration: 30000,
					}
					bodyString, _ := json.Marshal(payload)
					ctx.Request = httptest.NewRequest(
						http.MethodPost,
						"/mobile/phongnt/call",
						strings.NewReader(string(bodyString)),
					)
					ctx.AddParam("user_name", "phongnt")
				},
			},
			expectedResponse: "null",
			expectedStatus:   http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			c := &TelecomController{
				mobilService: tt.fields.mobilService,
			}
			tt.args.mockCtx(ctx)
			c.SaveCall(ctx)
			assert.Equal(t, tt.expectedResponse, w.Body.String())
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestTelecomController_GetBill(t *testing.T) {
	type fields struct {
		mobilService MobileService
	}
	type args struct {
		mockCtx func(ctx *gin.Context)
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedResponse string
		expectedStatus   int
	}{
		// TODO: Add test cases.
		{
			name: "error case",
			fields: fields{
				mobilService: func() MobileService {
					mockService := &mocks.MobileService{}
					mockService.On("GetUserBilling", "phongnt").Return(nil, error2.NewCError(500, "error occurred"))
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					ctx.Request = httptest.NewRequest(
						http.MethodGet,
						"/mobile/phongnt/billing",
						strings.NewReader(""),
					)
					ctx.AddParam("user_name", "phongnt")
				},
			},
			expectedResponse: "{\"code\":500,\"message\":\"error occurred\"}",
			expectedStatus:   http.StatusBadRequest,
		},
		{
			name: "happy case",
			fields: fields{
				mobilService: func() MobileService {
					mockService := &mocks.MobileService{}
					mockService.On("GetUserBilling", "phongnt").Return(&model.BillingData{
						CallCount:  1,
						BlockCount: 1,
					}, nil)
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					ctx.Request = httptest.NewRequest(
						http.MethodGet,
						"/mobile/phongnt/billing",
						strings.NewReader(""),
					)
					ctx.AddParam("user_name", "phongnt")
				},
			},
			expectedResponse: "{\"call_count\":1,\"block_count\":1}",
			expectedStatus:   http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			c := &TelecomController{
				mobilService: tt.fields.mobilService,
			}
			tt.args.mockCtx(ctx)
			c.GetBill(ctx)
			assert.Equal(t, tt.expectedResponse, w.Body.String())
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
