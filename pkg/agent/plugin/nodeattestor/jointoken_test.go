package nodeattestor_test

import (
	"context"
	"errors"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
	"github.com/spiffe/spire/test/spiretest"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestJoinToken(t *testing.T) {
	log, _ := test.NewNullLogger()
	attestor := nodeattestor.JoinToken(log, "foo")

	t.Run("success", func(t *testing.T) {
		err := attestor.Attest(context.Background(), &fakeStream{
			expectData: nodeattestor.AttestationData{Type: "join_token", Payload: []byte("foo")},
		})
		require.NoError(t, err)
	})

	t.Run("attestation fails", func(t *testing.T) {
		err := attestor.Attest(context.Background(), &fakeStream{
			attestationErr: errors.New("ohno"),
			expectData:     nodeattestor.AttestationData{Type: "join_token", Payload: []byte("bar")},
			challenges:     challenges("hello"),
		})
		// ServerStream errors are not the responsibility of the plugin, so
		// we shouldn't wrap them. ServerStream implementations are responsible
		// for the shape of those errors.
		spiretest.RequireGRPCStatus(t, err, codes.Unknown, "ohno")
	})

	t.Run("server issues unexpected challenge", func(t *testing.T) {
		err := attestor.Attest(context.Background(), &fakeStream{
			expectData: nodeattestor.AttestationData{Type: "join_token", Payload: []byte("foo")},
			challenges: challenges("hello"),
		})
		spiretest.RequireGRPCStatus(t, err, codes.Internal, "nodeattestor(join_token): server issued unexpected challenge")
	})
}
