// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package xds

import (
	"testing"

	envoy_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/consul/agent/structs"
)

func TestInjectHTTPHeaderManipToVirtualHost(t *testing.T) {
	tests := []struct {
		name        string
		reqHeaders  *structs.HTTPHeaderModifiers
		respHeaders *structs.HTTPHeaderModifiers
		expectReqAdd    int
		expectReqRemove int
		expectRespAdd   int
		expectRespRemove int
	}{
		{
			name:        "nil headers",
			reqHeaders:  nil,
			respHeaders: nil,
		},
		{
			name: "request headers only - set",
			reqHeaders: &structs.HTTPHeaderModifiers{
				Set: map[string]string{
					"x-source-service": "web",
				},
			},
			respHeaders:     nil,
			expectReqAdd:    1,
			expectReqRemove: 0,
		},
		{
			name: "request headers - add and remove",
			reqHeaders: &structs.HTTPHeaderModifiers{
				Add: map[string]string{
					"x-trace-id": "abc123",
				},
				Remove: []string{"x-internal-only"},
			},
			respHeaders:     nil,
			expectReqAdd:    1,
			expectReqRemove: 1,
		},
		{
			name:       "response headers only",
			reqHeaders: nil,
			respHeaders: &structs.HTTPHeaderModifiers{
				Set: map[string]string{
					"x-served-by": "consul-mesh",
				},
				Remove: []string{"server"},
			},
			expectRespAdd:    1,
			expectRespRemove: 1,
		},
		{
			name: "both request and response headers",
			reqHeaders: &structs.HTTPHeaderModifiers{
				Set: map[string]string{
					"x-source": "web",
				},
				Add: map[string]string{
					"x-request-start": "t=now",
				},
				Remove: []string{"x-debug"},
			},
			respHeaders: &structs.HTTPHeaderModifiers{
				Set: map[string]string{
					"x-envoy-upstream-service-time": "true",
				},
			},
			expectReqAdd:     2,
			expectReqRemove:  1,
			expectRespAdd:    1,
			expectRespRemove: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vh := &envoy_route_v3.VirtualHost{
				Name:    "test",
				Domains: []string{"*"},
			}

			err := injectHTTPHeaderManipToVirtualHost(tt.reqHeaders, tt.respHeaders, vh)
			require.NoError(t, err)

			require.Len(t, vh.RequestHeadersToAdd, tt.expectReqAdd)
			require.Len(t, vh.RequestHeadersToRemove, tt.expectReqRemove)
			require.Len(t, vh.ResponseHeadersToAdd, tt.expectRespAdd)
			require.Len(t, vh.ResponseHeadersToRemove, tt.expectRespRemove)
		})
	}
}

func TestMakeHeadersValueOptions_AppendAction(t *testing.T) {
	// Verify that "add" mode uses APPEND and "set" mode uses OVERWRITE
	addOpts := makeHeadersValueOptions(map[string]string{"x-foo": "bar"}, true)
	require.Len(t, addOpts, 1)
	// Default is APPEND_IF_EXISTS_OR_ADD (value 0)
	require.Equal(t, envoy_core_v3.HeaderValueOption_APPEND_IF_EXISTS_OR_ADD, addOpts[0].AppendAction)

	setOpts := makeHeadersValueOptions(map[string]string{"x-foo": "bar"}, false)
	require.Len(t, setOpts, 1)
	require.Equal(t, envoy_core_v3.HeaderValueOption_OVERWRITE_IF_EXISTS_OR_ADD, setOpts[0].AppendAction)
}
