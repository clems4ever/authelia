package session

import (
	"testing"

	"github.com/fasthttp/session/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldEncryptAndDecrypt(t *testing.T) {
	payload := session.Dict{}
	payload.Set("key", "value")

	dst, err := payload.MarshalMsg(nil)
	require.NoError(t, err)

	serializer := NewEncryptingSerializer("asecret")
	encryptedDst, err := serializer.Encode(payload)
	require.NoError(t, err)

	assert.NotEqual(t, dst, encryptedDst)

	decodedPayload := session.Dict{}
	err = serializer.Decode(&decodedPayload, encryptedDst)
	require.NoError(t, err)

	assert.Equal(t, "value", decodedPayload.Get("key"))
}

func TestShouldSupportUnencryptedSessionForBackwardCompatibility(t *testing.T) {
	payload := session.Dict{}
	payload.Set("key", "value")

	dst, err := payload.MarshalMsg(nil)
	require.NoError(t, err)

	serializer := NewEncryptingSerializer("asecret")

	decodedPayload := session.Dict{}
	err = serializer.Decode(&decodedPayload, dst)
	require.NoError(t, err)

	assert.Equal(t, "value", decodedPayload.Get("key"))
}
