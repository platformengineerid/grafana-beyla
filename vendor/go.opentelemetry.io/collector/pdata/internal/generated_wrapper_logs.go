// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package internal

import (
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlplogs "go.opentelemetry.io/collector/pdata/internal/data/protogen/logs/v1"
)

type ResourceLogsSlice struct {
	orig *[]*otlplogs.ResourceLogs
}

func GetOrigResourceLogsSlice(ms ResourceLogsSlice) *[]*otlplogs.ResourceLogs {
	return ms.orig
}

func NewResourceLogsSlice(orig *[]*otlplogs.ResourceLogs) ResourceLogsSlice {
	return ResourceLogsSlice{orig: orig}
}

type ResourceLogs struct {
	orig *otlplogs.ResourceLogs
}

func GetOrigResourceLogs(ms ResourceLogs) *otlplogs.ResourceLogs {
	return ms.orig
}

func NewResourceLogs(orig *otlplogs.ResourceLogs) ResourceLogs {
	return ResourceLogs{orig: orig}
}

type ScopeLogsSlice struct {
	orig *[]*otlplogs.ScopeLogs
}

func GetOrigScopeLogsSlice(ms ScopeLogsSlice) *[]*otlplogs.ScopeLogs {
	return ms.orig
}

func NewScopeLogsSlice(orig *[]*otlplogs.ScopeLogs) ScopeLogsSlice {
	return ScopeLogsSlice{orig: orig}
}

type ScopeLogs struct {
	orig *otlplogs.ScopeLogs
}

func GetOrigScopeLogs(ms ScopeLogs) *otlplogs.ScopeLogs {
	return ms.orig
}

func NewScopeLogs(orig *otlplogs.ScopeLogs) ScopeLogs {
	return ScopeLogs{orig: orig}
}

type LogRecordSlice struct {
	orig *[]*otlplogs.LogRecord
}

func GetOrigLogRecordSlice(ms LogRecordSlice) *[]*otlplogs.LogRecord {
	return ms.orig
}

func NewLogRecordSlice(orig *[]*otlplogs.LogRecord) LogRecordSlice {
	return LogRecordSlice{orig: orig}
}

type LogRecord struct {
	orig *otlplogs.LogRecord
}

func GetOrigLogRecord(ms LogRecord) *otlplogs.LogRecord {
	return ms.orig
}

func NewLogRecord(orig *otlplogs.LogRecord) LogRecord {
	return LogRecord{orig: orig}
}

func GenerateTestResourceLogsSlice() ResourceLogsSlice {
	orig := []*otlplogs.ResourceLogs{}
	tv := NewResourceLogsSlice(&orig)
	FillTestResourceLogsSlice(tv)
	return tv
}

func FillTestResourceLogsSlice(tv ResourceLogsSlice) {
	*tv.orig = make([]*otlplogs.ResourceLogs, 7)
	for i := 0; i < 7; i++ {
		(*tv.orig)[i] = &otlplogs.ResourceLogs{}
		FillTestResourceLogs(NewResourceLogs((*tv.orig)[i]))
	}
}

func GenerateTestResourceLogs() ResourceLogs {
	orig := otlplogs.ResourceLogs{}
	tv := NewResourceLogs(&orig)
	FillTestResourceLogs(tv)
	return tv
}

func FillTestResourceLogs(tv ResourceLogs) {
	FillTestResource(NewResource(&tv.orig.Resource))
	tv.orig.SchemaUrl = "https://opentelemetry.io/schemas/1.5.0"
	FillTestScopeLogsSlice(NewScopeLogsSlice(&tv.orig.ScopeLogs))
}

func GenerateTestScopeLogsSlice() ScopeLogsSlice {
	orig := []*otlplogs.ScopeLogs{}
	tv := NewScopeLogsSlice(&orig)
	FillTestScopeLogsSlice(tv)
	return tv
}

func FillTestScopeLogsSlice(tv ScopeLogsSlice) {
	*tv.orig = make([]*otlplogs.ScopeLogs, 7)
	for i := 0; i < 7; i++ {
		(*tv.orig)[i] = &otlplogs.ScopeLogs{}
		FillTestScopeLogs(NewScopeLogs((*tv.orig)[i]))
	}
}

func GenerateTestScopeLogs() ScopeLogs {
	orig := otlplogs.ScopeLogs{}
	tv := NewScopeLogs(&orig)
	FillTestScopeLogs(tv)
	return tv
}

func FillTestScopeLogs(tv ScopeLogs) {
	FillTestInstrumentationScope(NewInstrumentationScope(&tv.orig.Scope))
	tv.orig.SchemaUrl = "https://opentelemetry.io/schemas/1.5.0"
	FillTestLogRecordSlice(NewLogRecordSlice(&tv.orig.LogRecords))
}

func GenerateTestLogRecordSlice() LogRecordSlice {
	orig := []*otlplogs.LogRecord{}
	tv := NewLogRecordSlice(&orig)
	FillTestLogRecordSlice(tv)
	return tv
}

func FillTestLogRecordSlice(tv LogRecordSlice) {
	*tv.orig = make([]*otlplogs.LogRecord, 7)
	for i := 0; i < 7; i++ {
		(*tv.orig)[i] = &otlplogs.LogRecord{}
		FillTestLogRecord(NewLogRecord((*tv.orig)[i]))
	}
}

func GenerateTestLogRecord() LogRecord {
	orig := otlplogs.LogRecord{}
	tv := NewLogRecord(&orig)
	FillTestLogRecord(tv)
	return tv
}

func FillTestLogRecord(tv LogRecord) {
	tv.orig.ObservedTimeUnixNano = 1234567890
	tv.orig.TimeUnixNano = 1234567890
	tv.orig.TraceId = data.TraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1})
	tv.orig.SpanId = data.SpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1})
	tv.orig.Flags = 1
	tv.orig.SeverityText = "INFO"
	tv.orig.SeverityNumber = otlplogs.SeverityNumber(5)
	FillTestValue(NewValue(&tv.orig.Body))
	FillTestMap(NewMap(&tv.orig.Attributes))
	tv.orig.DroppedAttributesCount = uint32(17)
}
