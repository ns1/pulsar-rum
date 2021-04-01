// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bulkbeacon.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

////
// Classification of the device type on which a measurement was taken. Default is UNKNOWN.
type DeviceType int32

const (
	DeviceType_UNKNOWN         DeviceType = 0
	DeviceType_DESKTOP         DeviceType = 1
	DeviceType_MOBILE          DeviceType = 2
	DeviceType_GENERIC_SET_TOP DeviceType = 3
)

var DeviceType_name = map[int32]string{
	0: "UNKNOWN",
	1: "DESKTOP",
	2: "MOBILE",
	3: "GENERIC_SET_TOP",
}

var DeviceType_value = map[string]int32{
	"UNKNOWN":         0,
	"DESKTOP":         1,
	"MOBILE":          2,
	"GENERIC_SET_TOP": 3,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}

func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{0}
}

////
// Specifies the network location from which a measurement was taken. Either
// IP address OR Geo+ASN information maybe be provided -- but not both.  If both
// are provided then Geo+ASN takes precedence.
type Location struct {
	// IPv4 or IPv6 host address.
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// ISO 3166-1 (alpha-2) country code. If providing Geo+ASN info this value is
	// REQUIRED.
	GeoCountry string `protobuf:"bytes,2,opt,name=geo_country,json=geoCountry,proto3" json:"geo_country,omitempty"`
	// ISO 3166-2 subdivision code. If providing Geo+ASN info this value is OPTIONAL.
	GeoSubdiv string `protobuf:"bytes,3,opt,name=geo_subdiv,json=geoSubdiv,proto3" json:"geo_subdiv,omitempty"`
	// Public ASN. If providing Geo+ASN info this value is OPTIONAL.
	Asn                  int32    `protobuf:"varint,4,opt,name=asn,proto3" json:"asn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{0}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Location) GetGeoCountry() string {
	if m != nil {
		return m.GeoCountry
	}
	return ""
}

func (m *Location) GetGeoSubdiv() string {
	if m != nil {
		return m.GeoSubdiv
	}
	return ""
}

func (m *Location) GetAsn() int32 {
	if m != nil {
		return m.Asn
	}
	return 0
}

////
// Unification of attributes that identify the originator (provider) of a measurement and
// the owner of the data.
type Attribution struct {
	// REQUIRED. Associated measurement will we discarded without this.
	Jobid string `protobuf:"bytes,1,opt,name=jobid,proto3" json:"jobid,omitempty"`
	// REQUIRED. Associated measurement will we discarded without this.
	Location *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// OPTIONAL.
	DeviceType           DeviceType `protobuf:"varint,3,opt,name=device_type,json=deviceType,proto3,enum=bulkbeacon.DeviceType" json:"device_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Attribution) Reset()         { *m = Attribution{} }
func (m *Attribution) String() string { return proto.CompactTextString(m) }
func (*Attribution) ProtoMessage()    {}
func (*Attribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{1}
}

func (m *Attribution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attribution.Unmarshal(m, b)
}
func (m *Attribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attribution.Marshal(b, m, deterministic)
}
func (m *Attribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribution.Merge(m, src)
}
func (m *Attribution) XXX_Size() int {
	return xxx_messageInfo_Attribution.Size(m)
}
func (m *Attribution) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribution.DiscardUnknown(m)
}

var xxx_messageInfo_Attribution proto.InternalMessageInfo

func (m *Attribution) GetJobid() string {
	if m != nil {
		return m.Jobid
	}
	return ""
}

func (m *Attribution) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Attribution) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_UNKNOWN
}

////
// Detailed measurement of performance. Time values are absolute durations.
type DetailedPerformance struct {
	DnsLookupMs        int32 `protobuf:"varint,1,opt,name=dns_lookup_ms,json=dnsLookupMs,proto3" json:"dns_lookup_ms,omitempty"`
	TransportConnectMs int32 `protobuf:"varint,2,opt,name=transport_connect_ms,json=transportConnectMs,proto3" json:"transport_connect_ms,omitempty"`
	TlsConnectMs       int32 `protobuf:"varint,3,opt,name=tls_connect_ms,json=tlsConnectMs,proto3" json:"tls_connect_ms,omitempty"`
	WaitingMs          int32 `protobuf:"varint,4,opt,name=waiting_ms,json=waitingMs,proto3" json:"waiting_ms,omitempty"`
	// One of TTFB or TTLB is REQUIRED. Providing both is valid.
	TimeToFirstByteMs    int32    `protobuf:"varint,5,opt,name=time_to_first_byte_ms,json=timeToFirstByteMs,proto3" json:"time_to_first_byte_ms,omitempty"`
	TimeToLastByteMs     int32    `protobuf:"varint,6,opt,name=time_to_last_byte_ms,json=timeToLastByteMs,proto3" json:"time_to_last_byte_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailedPerformance) Reset()         { *m = DetailedPerformance{} }
func (m *DetailedPerformance) String() string { return proto.CompactTextString(m) }
func (*DetailedPerformance) ProtoMessage()    {}
func (*DetailedPerformance) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{2}
}

func (m *DetailedPerformance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailedPerformance.Unmarshal(m, b)
}
func (m *DetailedPerformance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailedPerformance.Marshal(b, m, deterministic)
}
func (m *DetailedPerformance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailedPerformance.Merge(m, src)
}
func (m *DetailedPerformance) XXX_Size() int {
	return xxx_messageInfo_DetailedPerformance.Size(m)
}
func (m *DetailedPerformance) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailedPerformance.DiscardUnknown(m)
}

var xxx_messageInfo_DetailedPerformance proto.InternalMessageInfo

func (m *DetailedPerformance) GetDnsLookupMs() int32 {
	if m != nil {
		return m.DnsLookupMs
	}
	return 0
}

func (m *DetailedPerformance) GetTransportConnectMs() int32 {
	if m != nil {
		return m.TransportConnectMs
	}
	return 0
}

func (m *DetailedPerformance) GetTlsConnectMs() int32 {
	if m != nil {
		return m.TlsConnectMs
	}
	return 0
}

func (m *DetailedPerformance) GetWaitingMs() int32 {
	if m != nil {
		return m.WaitingMs
	}
	return 0
}

func (m *DetailedPerformance) GetTimeToFirstByteMs() int32 {
	if m != nil {
		return m.TimeToFirstByteMs
	}
	return 0
}

func (m *DetailedPerformance) GetTimeToLastByteMs() int32 {
	if m != nil {
		return m.TimeToLastByteMs
	}
	return 0
}

////
// A simple, single latency value. Use this instead of DetailedPerformance if your
// measurement process captures only end-to-end timings.
type SimpleLatency struct {
	ValueMs              int32    `protobuf:"varint,1,opt,name=value_ms,json=valueMs,proto3" json:"value_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleLatency) Reset()         { *m = SimpleLatency{} }
func (m *SimpleLatency) String() string { return proto.CompactTextString(m) }
func (*SimpleLatency) ProtoMessage()    {}
func (*SimpleLatency) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{3}
}

func (m *SimpleLatency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleLatency.Unmarshal(m, b)
}
func (m *SimpleLatency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleLatency.Marshal(b, m, deterministic)
}
func (m *SimpleLatency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleLatency.Merge(m, src)
}
func (m *SimpleLatency) XXX_Size() int {
	return xxx_messageInfo_SimpleLatency.Size(m)
}
func (m *SimpleLatency) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleLatency.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleLatency proto.InternalMessageInfo

func (m *SimpleLatency) GetValueMs() int32 {
	if m != nil {
		return m.ValueMs
	}
	return 0
}

////
// A single value of uninterpreted dimension. Use this if your measurement process
// does not capture values that fit in either DetailedPerformance or SimpleLatency
// categories.
type Score struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Score) Reset()         { *m = Score{} }
func (m *Score) String() string { return proto.CompactTextString(m) }
func (*Score) ProtoMessage()    {}
func (*Score) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{4}
}

func (m *Score) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Score.Unmarshal(m, b)
}
func (m *Score) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Score.Marshal(b, m, deterministic)
}
func (m *Score) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Score.Merge(m, src)
}
func (m *Score) XXX_Size() int {
	return xxx_messageInfo_Score.Size(m)
}
func (m *Score) XXX_DiscardUnknown() {
	xxx_messageInfo_Score.DiscardUnknown(m)
}

var xxx_messageInfo_Score proto.InternalMessageInfo

func (m *Score) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

////
// Data for a single Attribution. That is: some JobID, from some network location, on
// some type of device.
type Payload struct {
	// Timestamp (in unix epoch nanoseconds) the data in this object was observed.
	// If not provided the timestamp that a message containing this Payload was
	// received by NS1 servers will be used.
	TimstampEpochNs uint64 `protobuf:"fixed64,1,opt,name=timstamp_epoch_ns,json=timstampEpochNs,proto3" json:"timstamp_epoch_ns,omitempty"`
	// Final HTTP status code resulting from fetching this data. If you don't know
	// this value (or can't observe it) use 200 to indicate success and 650 to indicate failure.
	StatusCode int32 `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// Overrides the attribution set on the Measurement level. This is OPTIONAL.
	Attribution *Attribution `protobuf:"bytes,3,opt,name=attribution,proto3" json:"attribution,omitempty"`
	// Time-to-live value in seconds. This determines how long we will use this data for
	// making routing decisions for the associated attribution.
	// This is OPTIONAL, unless the payload corresponds to a Static Beacon:
	// - If is not a static beacon and is not provided, the default value from the
	// associated Pulsar app/job will be used.
	// - If it's a static beacon and is not provided, the beacon will be rejected.
	DataTtl int32 `protobuf:"varint,7,opt,name=data_ttl,json=dataTtl,proto3" json:"data_ttl,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Payload_Simple
	//	*Payload_Detailed
	//	*Payload_Score
	Value                isPayload_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{5}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetTimstampEpochNs() uint64 {
	if m != nil {
		return m.TimstampEpochNs
	}
	return 0
}

func (m *Payload) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Payload) GetAttribution() *Attribution {
	if m != nil {
		return m.Attribution
	}
	return nil
}

func (m *Payload) GetDataTtl() int32 {
	if m != nil {
		return m.DataTtl
	}
	return 0
}

type isPayload_Value interface {
	isPayload_Value()
}

type Payload_Simple struct {
	Simple *SimpleLatency `protobuf:"bytes,4,opt,name=simple,proto3,oneof"`
}

type Payload_Detailed struct {
	Detailed *DetailedPerformance `protobuf:"bytes,5,opt,name=detailed,proto3,oneof"`
}

type Payload_Score struct {
	Score *Score `protobuf:"bytes,6,opt,name=score,proto3,oneof"`
}

func (*Payload_Simple) isPayload_Value() {}

func (*Payload_Detailed) isPayload_Value() {}

func (*Payload_Score) isPayload_Value() {}

func (m *Payload) GetValue() isPayload_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Payload) GetSimple() *SimpleLatency {
	if x, ok := m.GetValue().(*Payload_Simple); ok {
		return x.Simple
	}
	return nil
}

func (m *Payload) GetDetailed() *DetailedPerformance {
	if x, ok := m.GetValue().(*Payload_Detailed); ok {
		return x.Detailed
	}
	return nil
}

func (m *Payload) GetScore() *Score {
	if x, ok := m.GetValue().(*Payload_Score); ok {
		return x.Score
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Payload_Simple)(nil),
		(*Payload_Detailed)(nil),
		(*Payload_Score)(nil),
	}
}

////
// A container for multiple Payloads and some top-level defaults.
type Measurement struct {
	// Default attribution for payloads that do not provide their own Attribution object.
	// This is OPTIONAL but not providing it means any Payload not containing an Attribution
	// will be discarded.
	Attribution *Attribution `protobuf:"bytes,1,opt,name=attribution,proto3" json:"attribution,omitempty"`
	// May contain data from any number of JobIDs but those jobs must all be owned by the
	// same AppID -- which is implied by the Beacon object that contains this Measurement.
	Payloads             []*Payload `protobuf:"bytes,2,rep,name=payloads,proto3" json:"payloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Measurement) Reset()         { *m = Measurement{} }
func (m *Measurement) String() string { return proto.CompactTextString(m) }
func (*Measurement) ProtoMessage()    {}
func (*Measurement) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{6}
}

func (m *Measurement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Measurement.Unmarshal(m, b)
}
func (m *Measurement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Measurement.Marshal(b, m, deterministic)
}
func (m *Measurement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Measurement.Merge(m, src)
}
func (m *Measurement) XXX_Size() int {
	return xxx_messageInfo_Measurement.Size(m)
}
func (m *Measurement) XXX_DiscardUnknown() {
	xxx_messageInfo_Measurement.DiscardUnknown(m)
}

var xxx_messageInfo_Measurement proto.InternalMessageInfo

func (m *Measurement) GetAttribution() *Attribution {
	if m != nil {
		return m.Attribution
	}
	return nil
}

func (m *Measurement) GetPayloads() []*Payload {
	if m != nil {
		return m.Payloads
	}
	return nil
}

type Beacon struct {
	Appid                string         `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Signature            []byte         `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Measurements         []*Measurement `protobuf:"bytes,3,rep,name=measurements,proto3" json:"measurements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Beacon) Reset()         { *m = Beacon{} }
func (m *Beacon) String() string { return proto.CompactTextString(m) }
func (*Beacon) ProtoMessage()    {}
func (*Beacon) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{7}
}

func (m *Beacon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Beacon.Unmarshal(m, b)
}
func (m *Beacon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Beacon.Marshal(b, m, deterministic)
}
func (m *Beacon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Beacon.Merge(m, src)
}
func (m *Beacon) XXX_Size() int {
	return xxx_messageInfo_Beacon.Size(m)
}
func (m *Beacon) XXX_DiscardUnknown() {
	xxx_messageInfo_Beacon.DiscardUnknown(m)
}

var xxx_messageInfo_Beacon proto.InternalMessageInfo

func (m *Beacon) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *Beacon) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Beacon) GetMeasurements() []*Measurement {
	if m != nil {
		return m.Measurements
	}
	return nil
}

type Beacons struct {
	Beacons              []*Beacon `protobuf:"bytes,1,rep,name=beacons,proto3" json:"beacons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Beacons) Reset()         { *m = Beacons{} }
func (m *Beacons) String() string { return proto.CompactTextString(m) }
func (*Beacons) ProtoMessage()    {}
func (*Beacons) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{8}
}

func (m *Beacons) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Beacons.Unmarshal(m, b)
}
func (m *Beacons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Beacons.Marshal(b, m, deterministic)
}
func (m *Beacons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Beacons.Merge(m, src)
}
func (m *Beacons) XXX_Size() int {
	return xxx_messageInfo_Beacons.Size(m)
}
func (m *Beacons) XXX_DiscardUnknown() {
	xxx_messageInfo_Beacons.DiscardUnknown(m)
}

var xxx_messageInfo_Beacons proto.InternalMessageInfo

func (m *Beacons) GetBeacons() []*Beacon {
	if m != nil {
		return m.Beacons
	}
	return nil
}

type IngestResult struct {
	Processed            uint32   `protobuf:"varint,1,opt,name=processed,proto3" json:"processed,omitempty"`
	Failed               uint32   `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngestResult) Reset()         { *m = IngestResult{} }
func (m *IngestResult) String() string { return proto.CompactTextString(m) }
func (*IngestResult) ProtoMessage()    {}
func (*IngestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_67b32dc0c8a719ce, []int{9}
}

func (m *IngestResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngestResult.Unmarshal(m, b)
}
func (m *IngestResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngestResult.Marshal(b, m, deterministic)
}
func (m *IngestResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngestResult.Merge(m, src)
}
func (m *IngestResult) XXX_Size() int {
	return xxx_messageInfo_IngestResult.Size(m)
}
func (m *IngestResult) XXX_DiscardUnknown() {
	xxx_messageInfo_IngestResult.DiscardUnknown(m)
}

var xxx_messageInfo_IngestResult proto.InternalMessageInfo

func (m *IngestResult) GetProcessed() uint32 {
	if m != nil {
		return m.Processed
	}
	return 0
}

func (m *IngestResult) GetFailed() uint32 {
	if m != nil {
		return m.Failed
	}
	return 0
}

func init() {
	proto.RegisterEnum("bulkbeacon.DeviceType", DeviceType_name, DeviceType_value)
	proto.RegisterType((*Location)(nil), "bulkbeacon.Location")
	proto.RegisterType((*Attribution)(nil), "bulkbeacon.Attribution")
	proto.RegisterType((*DetailedPerformance)(nil), "bulkbeacon.DetailedPerformance")
	proto.RegisterType((*SimpleLatency)(nil), "bulkbeacon.SimpleLatency")
	proto.RegisterType((*Score)(nil), "bulkbeacon.Score")
	proto.RegisterType((*Payload)(nil), "bulkbeacon.Payload")
	proto.RegisterType((*Measurement)(nil), "bulkbeacon.Measurement")
	proto.RegisterType((*Beacon)(nil), "bulkbeacon.Beacon")
	proto.RegisterType((*Beacons)(nil), "bulkbeacon.Beacons")
	proto.RegisterType((*IngestResult)(nil), "bulkbeacon.IngestResult")
}

func init() { proto.RegisterFile("bulkbeacon.proto", fileDescriptor_67b32dc0c8a719ce) }

var fileDescriptor_67b32dc0c8a719ce = []byte{
	// 792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6f, 0xe3, 0x36,
	0x10, 0xf5, 0x47, 0x2d, 0xdb, 0x23, 0x67, 0xd7, 0x4b, 0xbb, 0x5b, 0x6f, 0xd1, 0x60, 0x03, 0xa1,
	0x87, 0x34, 0x28, 0xd2, 0x40, 0x39, 0x04, 0x41, 0xd1, 0x43, 0xfc, 0xd1, 0x24, 0x88, 0xed, 0x18,
	0xb4, 0x8b, 0x1e, 0x05, 0x5a, 0x62, 0x5c, 0x35, 0x32, 0x29, 0x88, 0x54, 0x0a, 0xa1, 0xff, 0xa1,
	0x87, 0xfe, 0xb6, 0xfe, 0xa0, 0x82, 0xa4, 0x6c, 0xc9, 0xc9, 0x69, 0x6f, 0x9e, 0x37, 0x6f, 0x38,
	0x33, 0xef, 0x8d, 0x05, 0xdd, 0x75, 0x1a, 0x3d, 0xaf, 0x29, 0xf1, 0x39, 0x3b, 0x8f, 0x13, 0x2e,
	0x39, 0x82, 0x02, 0x71, 0xfe, 0x86, 0xd6, 0x94, 0xfb, 0x44, 0x86, 0x9c, 0xa1, 0x63, 0x80, 0x30,
	0xf6, 0x48, 0x10, 0x24, 0x54, 0x88, 0x41, 0xf5, 0xa4, 0x7a, 0xda, 0xc6, 0xed, 0x30, 0xbe, 0x31,
	0x00, 0xfa, 0x0c, 0xf6, 0x86, 0x72, 0xcf, 0xe7, 0x29, 0x93, 0x49, 0x36, 0xa8, 0xe9, 0x3c, 0x6c,
	0x28, 0x1f, 0x19, 0x44, 0xd5, 0x2b, 0x82, 0x48, 0xd7, 0x41, 0xf8, 0x32, 0xa8, 0x9b, 0xfa, 0x0d,
	0xe5, 0x4b, 0x0d, 0xa0, 0x2e, 0xd4, 0x89, 0x60, 0x83, 0xaf, 0x4e, 0xaa, 0xa7, 0x0d, 0xac, 0x7e,
	0x3a, 0xff, 0x54, 0xc1, 0xbe, 0x91, 0x32, 0x09, 0xd7, 0xa9, 0x1e, 0xa0, 0x0f, 0x8d, 0x3f, 0xf9,
	0x3a, 0x0c, 0xf2, 0xde, 0x26, 0x40, 0x17, 0xd0, 0x8a, 0xf2, 0x11, 0x75, 0x53, 0xdb, 0xed, 0x9f,
	0x97, 0x76, 0xda, 0x8d, 0x8f, 0xf7, 0x2c, 0x74, 0x05, 0x76, 0x40, 0x5f, 0x42, 0x9f, 0x7a, 0x32,
	0x8b, 0xa9, 0x9e, 0xe4, 0x9d, 0xfb, 0xb1, 0x5c, 0x34, 0xd6, 0xe9, 0x55, 0x16, 0x53, 0x0c, 0xc1,
	0xfe, 0xb7, 0xf3, 0x6f, 0x0d, 0x7a, 0x63, 0x2a, 0x49, 0x18, 0xd1, 0x60, 0x41, 0x93, 0x27, 0x9e,
	0x6c, 0x09, 0xf3, 0x29, 0x72, 0xe0, 0x28, 0x60, 0xc2, 0x8b, 0x38, 0x7f, 0x4e, 0x63, 0x6f, 0x6b,
	0xc4, 0x69, 0x60, 0x3b, 0x60, 0x62, 0xaa, 0xb1, 0x99, 0x40, 0x17, 0xd0, 0x97, 0x09, 0x61, 0x22,
	0xe6, 0x89, 0xf4, 0x7c, 0xce, 0x18, 0xf5, 0xa5, 0xa2, 0xd6, 0x34, 0x15, 0xed, 0x73, 0x23, 0x93,
	0x9a, 0x09, 0xf4, 0x3d, 0xbc, 0x93, 0x91, 0x28, 0x73, 0xeb, 0x9a, 0xdb, 0x91, 0x91, 0x28, 0x58,
	0xc7, 0x00, 0x7f, 0x91, 0x50, 0x86, 0x6c, 0xa3, 0x18, 0x46, 0xbd, 0x76, 0x8e, 0xe8, 0xb6, 0x5f,
	0xcb, 0x70, 0x4b, 0x3d, 0xc9, 0xbd, 0xa7, 0x30, 0x11, 0xd2, 0x5b, 0x67, 0x92, 0x2a, 0x66, 0x43,
	0x33, 0x3f, 0xa8, 0xe4, 0x8a, 0xff, 0xaa, 0x52, 0xc3, 0x4c, 0xd2, 0x99, 0x40, 0xe7, 0xd0, 0xdf,
	0x55, 0x44, 0xa4, 0x54, 0x60, 0xe9, 0x82, 0xae, 0x29, 0x98, 0x92, 0x1d, 0xdf, 0x39, 0x83, 0xa3,
	0x65, 0xb8, 0x8d, 0x23, 0x3a, 0x25, 0x92, 0x32, 0x3f, 0x43, 0x9f, 0xa0, 0xf5, 0x42, 0xa2, 0x94,
	0x16, 0x42, 0x34, 0x75, 0x3c, 0x13, 0xce, 0x31, 0x34, 0x96, 0x3e, 0x4f, 0xa8, 0xb2, 0x52, 0x63,
	0x9a, 0x50, 0xc5, 0x26, 0x70, 0xfe, 0xab, 0x41, 0x73, 0x41, 0xb2, 0x88, 0x93, 0x00, 0x9d, 0x81,
	0x9a, 0x4d, 0x48, 0xb2, 0x8d, 0x3d, 0x1a, 0x73, 0xff, 0x0f, 0x8f, 0x99, 0xe7, 0x2c, 0xfc, 0x7e,
	0x97, 0x98, 0x28, 0x7c, 0xae, 0x4f, 0x4f, 0x48, 0x22, 0x53, 0x25, 0x56, 0x40, 0x73, 0x49, 0xc1,
	0x40, 0x23, 0x1e, 0x50, 0x74, 0x0d, 0x36, 0x29, 0x0e, 0x49, 0xeb, 0x68, 0xbb, 0xdf, 0x94, 0x1d,
	0x2f, 0xdd, 0x19, 0x2e, 0x73, 0xd5, 0x36, 0x01, 0x91, 0xc4, 0x93, 0x32, 0x1a, 0x34, 0xcd, 0x36,
	0x2a, 0x5e, 0xc9, 0x08, 0x5d, 0x82, 0x25, 0xf4, 0xe6, 0x5a, 0x76, 0xdb, 0xfd, 0x54, 0x7e, 0xf0,
	0x40, 0x93, 0xbb, 0x0a, 0xce, 0xa9, 0xe8, 0x17, 0x68, 0x05, 0xf9, 0x09, 0x69, 0x0f, 0x6c, 0xf7,
	0xf3, 0xe1, 0xe5, 0xbd, 0x39, 0xaf, 0xbb, 0x0a, 0xde, 0x97, 0xa0, 0x1f, 0xa0, 0x21, 0x94, 0x82,
	0xda, 0x0e, 0xdb, 0xfd, 0x70, 0xd0, 0x52, 0x25, 0xee, 0x2a, 0xd8, 0x30, 0x86, 0xcd, 0x5c, 0x63,
	0x27, 0x03, 0x7b, 0x46, 0x89, 0x48, 0x13, 0xba, 0xa5, 0x4c, 0xbe, 0x16, 0xa3, 0xfa, 0x05, 0x62,
	0xfc, 0x04, 0xad, 0xd8, 0xf8, 0xa3, 0x0e, 0xb7, 0x7e, 0x6a, 0xbb, 0xbd, 0x72, 0x5d, 0xee, 0x1d,
	0xde, 0x93, 0x9c, 0x0c, 0xac, 0xa1, 0xce, 0x29, 0xc7, 0x49, 0x1c, 0x17, 0x7f, 0x5e, 0x1d, 0xa0,
	0xef, 0xa0, 0x2d, 0xc2, 0x0d, 0x23, 0x32, 0x4d, 0x8c, 0x6f, 0x1d, 0x5c, 0x00, 0xe8, 0x67, 0xe8,
	0x6c, 0x8b, 0xc1, 0xd5, 0xfd, 0xd7, 0x5f, 0x8f, 0x5a, 0x5a, 0x0c, 0x1f, 0x90, 0x9d, 0x2b, 0x68,
	0x9a, 0xd6, 0x02, 0xfd, 0x08, 0x4d, 0x43, 0x57, 0x17, 0xa4, 0x9e, 0x40, 0xe5, 0x27, 0x0c, 0x0b,
	0xef, 0x28, 0xce, 0x18, 0x3a, 0xf7, 0x6c, 0x43, 0x85, 0xc4, 0x54, 0xa4, 0x91, 0x54, 0x33, 0xc6,
	0x09, 0xf7, 0xa9, 0x10, 0xd4, 0x4c, 0x7f, 0x84, 0x0b, 0x00, 0x7d, 0x04, 0xeb, 0xc9, 0xb8, 0x59,
	0xd3, 0xa9, 0x3c, 0x3a, 0xbb, 0x05, 0x28, 0xbe, 0x22, 0xc8, 0x86, 0xe6, 0x6f, 0xf3, 0x87, 0xf9,
	0xe3, 0xef, 0xf3, 0x6e, 0x45, 0x05, 0xe3, 0xc9, 0xf2, 0x61, 0xf5, 0xb8, 0xe8, 0x56, 0x11, 0x80,
	0x35, 0x7b, 0x1c, 0xde, 0x4f, 0x27, 0xdd, 0x1a, 0xea, 0xc1, 0xfb, 0xdb, 0xc9, 0x7c, 0x82, 0xef,
	0x47, 0xde, 0x72, 0xb2, 0xf2, 0x14, 0xa1, 0xee, 0x2e, 0xa0, 0xb7, 0x48, 0x23, 0x41, 0x92, 0x31,
	0x91, 0xc4, 0x0c, 0xa6, 0xac, 0xb8, 0x06, 0xcb, 0x04, 0xa8, 0xf7, 0x76, 0x19, 0xf1, 0xed, 0xa0,
	0x0c, 0x96, 0xd7, 0x71, 0x2a, 0x6b, 0x4b, 0x7f, 0xe7, 0x2f, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x74, 0x38, 0xca, 0x22, 0xfb, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PulsarDataIngestionClient is the client API for PulsarDataIngestion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PulsarDataIngestionClient interface {
	Ingest(ctx context.Context, in *Beacons, opts ...grpc.CallOption) (*IngestResult, error)
}

type pulsarDataIngestionClient struct {
	cc *grpc.ClientConn
}

func NewPulsarDataIngestionClient(cc *grpc.ClientConn) PulsarDataIngestionClient {
	return &pulsarDataIngestionClient{cc}
}

func (c *pulsarDataIngestionClient) Ingest(ctx context.Context, in *Beacons, opts ...grpc.CallOption) (*IngestResult, error) {
	out := new(IngestResult)
	err := c.cc.Invoke(ctx, "/bulkbeacon.PulsarDataIngestion/Ingest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PulsarDataIngestionServer is the server API for PulsarDataIngestion service.
type PulsarDataIngestionServer interface {
	Ingest(context.Context, *Beacons) (*IngestResult, error)
}

// UnimplementedPulsarDataIngestionServer can be embedded to have forward compatible implementations.
type UnimplementedPulsarDataIngestionServer struct {
}

func (*UnimplementedPulsarDataIngestionServer) Ingest(ctx context.Context, req *Beacons) (*IngestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ingest not implemented")
}

func RegisterPulsarDataIngestionServer(s *grpc.Server, srv PulsarDataIngestionServer) {
	s.RegisterService(&_PulsarDataIngestion_serviceDesc, srv)
}

func _PulsarDataIngestion_Ingest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Beacons)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PulsarDataIngestionServer).Ingest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bulkbeacon.PulsarDataIngestion/Ingest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PulsarDataIngestionServer).Ingest(ctx, req.(*Beacons))
	}
	return interceptor(ctx, in, info, handler)
}

var _PulsarDataIngestion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bulkbeacon.PulsarDataIngestion",
	HandlerType: (*PulsarDataIngestionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ingest",
			Handler:    _PulsarDataIngestion_Ingest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bulkbeacon.proto",
}