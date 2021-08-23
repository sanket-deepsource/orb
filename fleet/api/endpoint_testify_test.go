package api_test

import (
	"fmt"
	"github.com/mainflux/mainflux"
	mfsdk "github.com/mainflux/mainflux/pkg/sdk/go"
	"github.com/ns1labs/orb/fleet"
	flmocks "github.com/ns1labs/orb/fleet/mocks"
	"github.com/ns1labs/orb/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"testing"
	"time"
)

var (
	validName, _ = types.NewIdentifier("eu-agents")
	agentToMock  = fleet.Agent{
		Name:          validName,
		MFOwnerID:     "user@example.com",
		MFThingID:     "1",
		MFKeyID:       "1",
		MFChannelID:   "2",
		Created:       time.Time{},
		OrbTags:       nil,
		AgentTags:     nil,
		AgentMetadata: nil,
		State:         0,
		LastHBData:    nil,
		LastHB:        time.Time{},
	}
)

func newServiceTestifyMock(auth mainflux.AuthServiceClient, url string, agentRepo fleet.AgentRepository) fleet.Service {
	agentGroupRepo := flmocks.NewAgentGroupRepository()
	agentComms := flmocks.NewFleetCommService()
	logger, _ := zap.NewDevelopment()
	config := mfsdk.Config{
		BaseURL: url,
	}

	mfsdk := mfsdk.NewSDK(config)
	return fleet.NewFleetService(logger, auth, agentRepo, agentGroupRepo, agentComms, mfsdk)

}

func TestCreateAgentGroupTestifyMock(t *testing.T) {

	agentRepo := &flmocks.AgentRepository{}
	agentRepo.On("Save", mock.Anything, agentToMock).Return(nil).Once()

	cli := newClientServerTestifyMock(t, agentRepo)
	defer cli.server.Close()

	cases := map[string]struct {
		req         string
		contentType string
		auth        string
		status      int
		location    string
	}{
		"add a valid agent": {
			req:         validJson,
			contentType: contentType,
			auth:        token,
			status:      http.StatusCreated,
			location:    "/agents",
		},
	}

	for desc, tc := range cases {
		req := testRequest{
			client:      cli.server.Client(),
			method:      http.MethodPost,
			url:         fmt.Sprintf("%s/agents", cli.server.URL),
			contentType: tc.contentType,
			token:       tc.auth,
			body:        strings.NewReader(tc.req),
		}
		res, err := req.make()
		assert.Nil(t, err, fmt.Sprintf("unexpected erro %s", err))
		assert.Equal(t, tc.status, res.StatusCode, fmt.Sprintf("%s: expected status code %d got %d", desc, tc.status, res.StatusCode))
	}

}

func newClientServerTestifyMock(t *testing.T, agentRepo fleet.AgentRepository) clientServer {
	t.Helper()
	users := flmocks.NewAuthService(map[string]string{token: email})

	thingsServer := newThingsServer(newThingsService(users))
	fleetService := newServiceTestifyMock(users, thingsServer.URL, agentRepo)
	fleetServer := newServer(fleetService)

	return clientServer{
		service: fleetService,
		server:  fleetServer,
	}
}
