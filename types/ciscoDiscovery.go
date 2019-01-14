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

import (
	"encoding/hex"
	"strings"
)

func (a CiscoDiscovery) CSVHeader() []string {
	return filter([]string{
		"Timestamp",
		"Version",  // int32
		"TTL",      // int32
		"Checksum", // int32
		"Values",   // []*CiscoDiscoveryValue
	})
}

func (a CiscoDiscovery) CSVRecord() []string {
	var vals []string
	for _, v := range a.Values {
		vals = append(vals, v.ToString())
	}
	return filter([]string{
		formatTimestamp(a.Timestamp),
		formatInt32(a.Version),  // int32
		formatInt32(a.TTL),      // int32
		formatInt32(a.Checksum), // int32
		join(vals...),           // []*CiscoDiscoveryValue
	})
}

func (a CiscoDiscovery) NetcapTimestamp() string {
	return a.Timestamp
}

func (v CiscoDiscoveryValue) ToString() string {

	var b strings.Builder
	b.WriteString(begin)
	b.WriteString(formatInt32(v.Type))
	b.WriteString(sep)
	b.WriteString(formatInt32(v.Length))
	b.WriteString(sep)
	b.WriteString(hex.EncodeToString(v.Value))
	b.WriteString(end)
	return b.String()
}
