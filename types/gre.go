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
	"strconv"
	"strings"
)

func (a GRE) CSVHeader() []string {
	return filter([]string{
		"Timestamp",
		"ChecksumPresent",   // bool
		"RoutingPresent",    // bool
		"KeyPresent",        // bool
		"SeqPresent",        // bool
		"StrictSourceRoute", // bool
		"AckPresent",        // bool
		"RecursionControl",  // int32
		"Flags",             // int32
		"Version",           // int32
		"Protocol",          // int32
		"Checksum",          // int32
		"Offset",            // int32
		"Key",               // uint32
		"Seq",               // uint32
		"Ack",               // uint32
		"Routing",           // *GRERouting
	})
}

func (a GRE) CSVRecord() []string {
	return filter([]string{
		formatTimestamp(a.Timestamp),
		strconv.FormatBool(a.ChecksumPresent),   // bool
		strconv.FormatBool(a.RoutingPresent),    // bool
		strconv.FormatBool(a.KeyPresent),        // bool
		strconv.FormatBool(a.SeqPresent),        // bool
		strconv.FormatBool(a.StrictSourceRoute), // bool
		strconv.FormatBool(a.AckPresent),        // bool
		formatInt32(a.RecursionControl),         // int32
		formatInt32(a.Flags),                    // int32
		formatInt32(a.Version),                  // int32
		formatInt32(a.Protocol),                 // int32
		formatInt32(a.Checksum),                 // int32
		formatInt32(a.Offset),                   // int32
		formatUint32(a.Key),                     // uint32
		formatUint32(a.Seq),                     // uint32
		formatUint32(a.Ack),                     // uint32
		a.Routing.GetString(),                   // *GRERouting
	})
}

func (a GRE) NetcapTimestamp() string {
	return a.Timestamp
}

func (r *GRERouting) GetString() string {

	if r == nil {
		return ""
	}

	var b strings.Builder

	b.WriteString(begin)
	b.WriteString(formatInt32(r.AddressFamily))
	b.WriteString(sep)
	b.WriteString(formatInt32(r.SREOffset))
	b.WriteString(sep)
	b.WriteString(formatInt32(r.SRELength))
	b.WriteString(sep)
	b.WriteString(hex.EncodeToString(r.RoutingInformation))
	b.WriteString(sep)
	b.WriteString(r.Next.GetString())
	b.WriteString(end)

	return b.String()
}
