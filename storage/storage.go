// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package storage provides the storage interfaces required by the various
// pieces of the CT monitor.
package storage

import (
	"context"
	"time"

	ct "github.com/google/certificate-transparency-go"
	"github.com/google/certificate-transparency-go/x509"
	"github.com/google/monologue/apicall"
	"github.com/google/monologue/ctlog"
)

// APICallWriter is an interface for storing individual calls to CT API
// endpoints.
type APICallWriter interface {
	WriteAPICall(ctx context.Context, l *ctlog.Log, apiCall *apicall.APICall) error
}

// STHWriter is an interface for storing STHs received from a CT Log.
type STHWriter interface {
	WriteSTH(ctx context.Context, l *ctlog.Log, sth *ct.SignedTreeHead, receivedAt time.Time, errs []error) error
}

// SCTWriter is an interface for storing SCTs received from a CT Log.
type SCTWriter interface {
	WriteSCT(ctx context.Context, l *ctlog.Log, sct *ct.SignedCertificateTimestamp, receivedAt time.Time, chain []*x509.Certificate, errs []error) error
}

// RootsWriter is an interface for storing root certificates retrieved from a CT get-roots call.
type RootsWriter interface {
	WriteRoots(ctx context.Context, l *ctlog.Log, roots []*x509.Certificate, receivedAt time.Time) error
}
