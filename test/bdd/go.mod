// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/edge-adapter/test/bdd

go 1.15

require (
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/cucumber/godog v0.9.0
	github.com/fsouza/go-dockerclient v1.6.5
	github.com/google/uuid v1.1.2
	github.com/hyperledger/aries-framework-go v0.1.6-0.20210305165454-c65334e35380
	github.com/hyperledger/aries-framework-go-ext/component/vdr/sidetree v0.0.0-20210303194824-a55a12f8d063
	github.com/hyperledger/aries-framework-go-ext/component/vdr/trustbloc v0.0.0-20210303194824-a55a12f8d063
	github.com/ory/hydra-client-go v1.4.10
	github.com/piprate/json-gold v0.4.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	github.com/tidwall/gjson v1.6.7
	github.com/trustbloc/edge-adapter v0.0.0
	github.com/trustbloc/edge-core v0.1.6-0.20210305000733-14a89fe44ae8
	github.com/trustbloc/hub-router v0.1.6-0.20210303154803-0289dfd31b46
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	gotest.tools/v3 v3.0.3 // indirect
)

replace (
	github.com/trustbloc/edge-adapter => ../..
	// https://github.com/ory/dockertest/issues/208#issuecomment-686820414
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6
)
