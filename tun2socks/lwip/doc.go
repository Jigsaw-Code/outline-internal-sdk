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

/*
The tun2socks/lwip package uses the lwIP (A Lightweight TCP/IP stack) library to translate between IP packets and
TCP/UDP protocols. The library is singleton, so only one instance can be created per process.

To create the instance with TCP/UDP handlers:

	// tcpHandler will be used to handle TCP streams, and udpHandler to handle UDP packets
	t2s, err := lwip.NewTun2SocksDevice(tcpHandler, udpHandler)
	if err != nil {
		// handle error
	}

InitInstance can only be called once. You will get an error if you call it more than once.
*/
package lwip