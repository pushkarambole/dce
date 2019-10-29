package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/Optum/Redbox/pkg/api/response"
	"github.com/Optum/Redbox/pkg/common"
	commonMock "github.com/Optum/Redbox/pkg/common/mocks"
	"github.com/Optum/Redbox/pkg/db"
	mockDB "github.com/Optum/Redbox/pkg/db/mocks"
	"github.com/Optum/Redbox/pkg/provision"
	provisionMock "github.com/Optum/Redbox/pkg/provision/mocks"
	"github.com/aws/aws-lambda-go/events"
)

func TestMain(m *testing.M) {
	os.Setenv("PROVISION_TOPIC", "Test_Provision_Topic")
	os.Exit(m.Run())
}

func TestCreateController_Call(t *testing.T) {
	type fields struct {
		Dao           db.DBer
		Provisioner   provision.Provisioner
		SNS           common.Notificationer
		LeaseTopicARN *string
	}
	type args struct {
		ctx context.Context
		req *events.APIGatewayProxyRequest
	}

	leaseTopicARN := "some-topic-arn"
	messageID := "message123456789"

	mockDB := &mockDB.DBer{}
	mockProv := &provisionMock.Provisioner{}
	mockSNS := &commonMock.Notificationer{}

	// Set up the mocks...
	mockProv.On("FindActiveLeaseForPrincipal", "123456").Return(createActiveLease(), nil)
	mockDB.On("GetReadyAccount").Return(createAccount(), nil)
	mockProv.On("FindLeaseWithAccount", "123456", "987654321").Return(createActiveLease(), nil)
	mockProv.On("ActivateAccount",
		true, "123456", "987654321", float64(50), "USD", mock.Anything, mock.Anything).Return(createActiveLease(), nil)
	mockDB.On("TransitionAccountStatus", "987654321", db.Ready, db.Leased).Return(createAccount(), nil)
	mockSNS.On("PublishMessage", &leaseTopicARN, mock.Anything, true).Return(&messageID, nil)

	testFields := &fields{
		Dao:           mockDB,
		Provisioner:   mockProv,
		SNS:           mockSNS,
		LeaseTopicARN: &leaseTopicARN,
	}

	successResponse := createSuccessCreateResponse()
	badRequestResponse := response.ClientBadRequestError(fmt.Sprintf("Failed to Parse Request Body: %s", ""))
	pastRequestResponse := response.BadRequestError("Requested lease has a desired expiry date less than today: 1570627876")

	successArgs := &args{ctx: context.Background(), req: createSuccessfulCreateRequest()}
	pastArgs := &args{ctx: context.Background(), req: createPastCreateRequest()}
	badArgs := &args{ctx: context.Background(), req: createBadCreateRequest()}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{name: "Bad request.", fields: *testFields, args: *badArgs, want: badRequestResponse, wantErr: false},
		{name: "Past request.", fields: *testFields, args: *pastArgs, want: pastRequestResponse, wantErr: false},
		{name: "Successful create.", fields: *testFields, args: *successArgs, want: *successResponse, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DbSvc = tt.fields.Dao
			Provisioner = tt.fields.Provisioner
			SnsSvc = tt.fields.SNS
			leaseTopicARN = *tt.fields.LeaseTopicARN

			got, err := Handler(tt.args.ctx, *tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateController.Call() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateController.Call() = %v, want %v", got, tt.want)
			}
		})
	}

}

func createSuccessfulCreateRequest() *events.APIGatewayProxyRequest {
	createLeaseRequest := &CreateLeaseRequest{
		PrincipalID:              "123456",
		AccountID:                "987654321",
		BudgetAmount:             50,
		BudgetCurrency:           "USD",
		BudgetNotificationEmails: []string{"user3@example.com", "user2@example.com"},
		ExpiresOn:                time.Now().AddDate(0, 0, 7).Unix(),
	}
	requestBodyBytes, _ := json.Marshal(createLeaseRequest)
	return &events.APIGatewayProxyRequest{
		Body:       string(requestBodyBytes),
		HTTPMethod: http.MethodPost,
		Path:       "/leases",
	}
}

func createBadCreateRequest() *events.APIGatewayProxyRequest {
	return &events.APIGatewayProxyRequest{}
}

func createPastCreateRequest() *events.APIGatewayProxyRequest {
	createLeaseRequest := &CreateLeaseRequest{
		PrincipalID:              "123456",
		AccountID:                "987654321",
		BudgetAmount:             50,
		BudgetCurrency:           "USD",
		BudgetNotificationEmails: []string{"user3@example.com", "user2@example.com"},
		ExpiresOn:                1570627876,
	}
	requestBodyBytes, _ := json.Marshal(createLeaseRequest)
	return &events.APIGatewayProxyRequest{
		Body:       string(requestBodyBytes),
		HTTPMethod: http.MethodPost,
		Path:       "/leases",
	}
}

func createActiveLease() *db.RedboxLease {
	return &db.RedboxLease{}
}

func createAccount() *db.RedboxAccount {
	return &db.RedboxAccount{
		ID: "987654321",
	}
}

func createSuccessCreateResponse() *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       "{\"accountId\":\"\",\"principalId\":\"\",\"id\":\"\",\"leaseStatus\":\"\",\"leaseStatusReason\":\"\",\"createdOn\":0,\"lastModifiedOn\":0,\"budgetAmount\":0,\"budgetCurrency\":\"\",\"budgetNotificationEmails\":null,\"leaseStatusModifiedOn\":0,\"expiresOn\":0}",
	}
}
