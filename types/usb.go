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
)

func (u USB) CSVHeader() []string {
	return filter([]string{
		"Timestamp",
		"ID",
		"EventType",
		"TransferType",
		"Direction",
		"EndpointNumber",
		"DeviceAddress",
		"BusID",
		"TimestampSec",
		"TimestampUsec",
		"Setup",
		"Data",
		"Status",
		"UrbLength",
		"UrbDataLength",
		"UrbInterval",
		"UrbStartFrame",
		"UrbCopyOfTransferFlags",
		"IsoNumDesc",
		"Payload",
	})
}

func (u USB) CSVRecord() []string {
	return filter([]string{
		formatTimestamp(u.Timestamp), // string
		formatUint64(u.ID),
		formatInt32(u.EventType),
		formatInt32(u.TransferType),
		formatInt32(u.Direction),
		formatInt32(u.EndpointNumber),
		formatInt32(u.DeviceAddress),
		formatInt32(u.BusID),
		formatInt64(u.TimestampSec),
		formatInt32(u.TimestampUsec),
		strconv.FormatBool(u.Setup),
		strconv.FormatBool(u.Data),
		formatInt32(u.Status),
		formatUint32(u.UrbLength),
		formatUint32(u.UrbDataLength),
		formatUint32(u.UrbInterval),
		formatUint32(u.UrbStartFrame),
		formatUint32(u.UrbCopyOfTransferFlags),
		formatUint32(u.IsoNumDesc),
		hex.EncodeToString(u.Payload),
	})
}

func (f USB) NetcapTimestamp() string {
	return f.Timestamp
}
