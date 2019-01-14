/*
 * NETCAP - Traffic Analysis Framework
 * Copyright (c) 2017 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package types

import "encoding/hex"

func (a NortelDiscovery) CSVHeader() []string {
	return filter([]string{
		"Timestamp",
		"IPAddress", // string
		"SegmentID", // []byte
		"Chassis",   // int32
		"Backplane", // int32
		"State",     // int32
		"NumLinks",  // int32
	})
}

func (a NortelDiscovery) CSVRecord() []string {
	return filter([]string{
		formatTimestamp(a.Timestamp),
		a.IPAddress,                     // string
		hex.EncodeToString(a.SegmentID), // []byte
		formatInt32(a.Chassis),          // int32
		formatInt32(a.Backplane),        // int32
		formatInt32(a.State),            // int32
		formatInt32(a.NumLinks),         // int32
	})
}

func (a NortelDiscovery) NetcapTimestamp() string {
	return a.Timestamp
}
