/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package governance

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/hyperledger/aries-framework-go/component/storageutil/mem"
	mockstorage "github.com/hyperledger/aries-framework-go/component/storageutil/mock"
	"github.com/hyperledger/aries-framework-go/spi/storage"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		p, err := New("", nil, mem.NewProvider(), nil, "")
		require.NoError(t, err)
		require.NotNil(t, p)
	})

	t.Run("test failed to open store", func(t *testing.T) {
		t.Parallel()

		p, err := New("", nil, &mockstorage.Provider{
			ErrOpenStore: fmt.Errorf("failed to open store"),
		}, nil, "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to open store")
		require.Nil(t, p)
	})
}

func TestProvider_IssueCredential(t *testing.T) {
	t.Parallel()

	t.Run("test error governance vc already issued", func(t *testing.T) {
		t.Parallel()

		memProvider := mem.NewProvider()

		store, err := memProvider.OpenStore(storeName)
		require.NoError(t, err)

		err = store.Put(fmt.Sprintf(governanceVCKey, "p1"), []byte("value"))
		require.NoError(t, err)

		p, err := New("", nil, memProvider, nil, "")
		require.NoError(t, err)

		_, err = p.IssueCredential("did:example:123", "p1")
		require.Error(t, err)
		require.Contains(t, err.Error(), "governance vc already issued")
	})

	t.Run("test failed to get governance vc from store", func(t *testing.T) {
		t.Parallel()

		p, err := New("", nil, &mockstorage.Provider{
			OpenStoreReturn: &mockstorage.Store{ErrGet: fmt.Errorf("failed to get")},
		},
			nil, "")
		require.NoError(t, err)

		_, err = p.IssueCredential("did:example:123", "p1")
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to get")
	})

	t.Run("test failed to send http request", func(t *testing.T) {
		t.Parallel()

		p, err := New("", nil, mem.NewProvider(),
			map[string]string{vcsGovernanceRequestTokenName: "token"}, "")
		require.NoError(t, err)

		p.httpClient = &mockHTTPClient{respErr: fmt.Errorf("failed to send http request")}

		_, err = p.IssueCredential("did:example:123", "p1")
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to send http request")
	})

	t.Run("test vcs return 500 status code", func(t *testing.T) {
		t.Parallel()

		p, err := New("", nil, mem.NewProvider(),
			map[string]string{vcsGovernanceRequestTokenName: "token"}, "")
		require.NoError(t, err)

		p.httpClient = &mockHTTPClient{respValue: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("failed ot issue vc"))),
		}}

		_, err = p.IssueCredential("did:example:123", "p1")
		require.Error(t, err)
		require.Contains(t, err.Error(), "http request: 500 failed ot issue vc")
	})

	t.Run("test put vc in db", func(t *testing.T) {
		t.Parallel()

		p, err := New("", nil, &mockstorage.Provider{
			OpenStoreReturn: &mockstorage.Store{ErrGet: storage.ErrDataNotFound, ErrPut: fmt.Errorf("error put")},
		},
			map[string]string{vcsGovernanceRequestTokenName: "token"}, "")
		require.NoError(t, err)

		p.httpClient = &mockHTTPClient{respValue: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("vc success"))),
		}}

		_, err = p.IssueCredential("did:example:123", "p1")
		require.Error(t, err)
		require.Contains(t, err.Error(), "error put")
	})

	t.Run("test success", func(t *testing.T) {
		t.Parallel()

		p, err := New("", nil, mem.NewProvider(),
			map[string]string{vcsGovernanceRequestTokenName: "token"}, "")
		require.NoError(t, err)

		p.httpClient = &mockHTTPClient{respValue: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte("vc success"))),
		}}

		data, err := p.IssueCredential("did:example:123", "p1")
		require.NoError(t, err)
		require.Equal(t, []byte("vc success"), data)
	})
}

type mockHTTPClient struct {
	respValue *http.Response
	respErr   error
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.respErr != nil {
		return nil, m.respErr
	}

	return m.respValue, nil
}
