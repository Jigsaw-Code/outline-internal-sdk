// Copyright 2023 Jigsaw Operations LLC
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

package lwip2transport

import (
	"context"
	"errors"
	"net"
	"os"
	"testing"

	"github.com/Jigsaw-Code/outline-internal-sdk/network"
	"github.com/Jigsaw-Code/outline-internal-sdk/transport"
	"github.com/stretchr/testify/require"
)

func TestStackClosedWriteError(t *testing.T) {
	h := &errTcpUdpHandler{err: errors.New("not supported")}
	t2s := reConfigurelwIPDeviceForTest(t, h, h)

	t2s.stack.Close() // close the underlying stack without calling Close
	n, err := t2s.Write([]byte{0x01})
	require.Exactly(t, 0, n)
	require.ErrorIs(t, err, network.ErrClosed)
	require.ErrorIs(t, err, os.ErrClosed) // network.ErrClosed should wrap os.ErrClosed
}

func reConfigurelwIPDeviceForTest(t *testing.T, sd transport.StreamDialer, pl transport.PacketListener) *lwIPDevice {
	t2s, err := ConfigureDevice(sd, pl)
	require.NoError(t, err)
	t2sInternal, ok := t2s.(*lwIPDevice)
	require.True(t, ok)
	return t2sInternal
}

type errTcpUdpHandler struct {
	err error
}

func (h *errTcpUdpHandler) Dial(ctx context.Context, raddr string) (transport.StreamConn, error) {
	return nil, h.err
}

func (h *errTcpUdpHandler) ListenPacket(ctx context.Context) (net.PacketConn, error) {
	return nil, h.err
}
