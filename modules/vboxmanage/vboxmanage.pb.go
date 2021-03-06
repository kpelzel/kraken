// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vboxmanage.proto

package vboxmanage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Config struct {
	Servers              map[string]*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PollingInterval      string             `protobuf:"bytes,2,opt,name=polling_interval,json=pollingInterval,proto3" json:"polling_interval,omitempty"`
	NameUrl              string             `protobuf:"bytes,3,opt,name=name_url,json=nameUrl,proto3" json:"name_url,omitempty"`
	ServerUrl            string             `protobuf:"bytes,4,opt,name=server_url,json=serverUrl,proto3" json:"server_url,omitempty"`
	UuidUrl              string             `protobuf:"bytes,5,opt,name=uuid_url,json=uuidUrl,proto3" json:"uuid_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b0aa3846a76f79, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetServers() map[string]*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

func (m *Config) GetPollingInterval() string {
	if m != nil {
		return m.PollingInterval
	}
	return ""
}

func (m *Config) GetNameUrl() string {
	if m != nil {
		return m.NameUrl
	}
	return ""
}

func (m *Config) GetServerUrl() string {
	if m != nil {
		return m.ServerUrl
	}
	return ""
}

func (m *Config) GetUuidUrl() string {
	if m != nil {
		return m.UuidUrl
	}
	return ""
}

func (*Config) XXX_MessageName() string {
	return "VBoxManage.Config"
}

type Server struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b0aa3846a76f79, []int{1}
}
func (m *Server) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Server.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return m.Size()
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Server) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Server) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (*Server) XXX_MessageName() string {
	return "VBoxManage.Server"
}
func init() {
	proto.RegisterType((*Config)(nil), "VBoxManage.Config")
	golang_proto.RegisterType((*Config)(nil), "VBoxManage.Config")
	proto.RegisterMapType((map[string]*Server)(nil), "VBoxManage.Config.ServersEntry")
	golang_proto.RegisterMapType((map[string]*Server)(nil), "VBoxManage.Config.ServersEntry")
	proto.RegisterType((*Server)(nil), "VBoxManage.Server")
	golang_proto.RegisterType((*Server)(nil), "VBoxManage.Server")
}

func init() { proto.RegisterFile("vboxmanage.proto", fileDescriptor_96b0aa3846a76f79) }
func init() { golang_proto.RegisterFile("vboxmanage.proto", fileDescriptor_96b0aa3846a76f79) }

var fileDescriptor_96b0aa3846a76f79 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0xb3, 0xe5, 0xff, 0x40, 0x94, 0xec, 0xa9, 0x90, 0x58, 0x09, 0x27, 0x3c, 0x58, 0x12,
	0xbc, 0xf8, 0xe7, 0x62, 0x30, 0x1e, 0x3c, 0xe8, 0xa1, 0x06, 0x0f, 0x5e, 0x48, 0xab, 0x4b, 0xdd,
	0xb8, 0x74, 0x9b, 0xa5, 0xdb, 0xc0, 0x73, 0xf8, 0x42, 0x1e, 0x39, 0xfa, 0x08, 0xa6, 0xbc, 0x88,
	0xd9, 0xd9, 0x1a, 0xb8, 0xcd, 0xf7, 0x7d, 0xbf, 0x99, 0xcc, 0x64, 0xa0, 0x9b, 0x47, 0x72, 0xbd,
	0x0c, 0x93, 0x30, 0x66, 0x7e, 0xaa, 0x64, 0x26, 0x29, 0xbc, 0x4c, 0xe5, 0xfa, 0x11, 0x9d, 0xfe,
	0x79, 0xcc, 0xb3, 0x0f, 0x1d, 0xf9, 0x6f, 0x72, 0x39, 0x8e, 0x65, 0x2c, 0xc7, 0x88, 0x44, 0x7a,
	0x81, 0x0a, 0x05, 0x56, 0xb6, 0x75, 0xf8, 0xe5, 0x40, 0xfd, 0x4e, 0x26, 0x0b, 0x1e, 0xd3, 0x2b,
	0x68, 0xac, 0x98, 0xca, 0x99, 0x5a, 0xb9, 0x64, 0x50, 0x19, 0xb5, 0x27, 0xa7, 0xfe, 0x7e, 0xae,
	0x6f, 0x21, 0xff, 0xd9, 0x12, 0xf7, 0x49, 0xa6, 0x36, 0xc1, 0x3f, 0x4f, 0xcf, 0xa0, 0x9b, 0x4a,
	0x21, 0x78, 0x12, 0xcf, 0x79, 0x92, 0x31, 0x95, 0x87, 0xc2, 0x75, 0x06, 0x64, 0xd4, 0x0a, 0x8e,
	0x4b, 0xff, 0xa1, 0xb4, 0x69, 0x0f, 0x9a, 0x49, 0xb8, 0x64, 0x73, 0xad, 0x84, 0x5b, 0x41, 0xa4,
	0x61, 0xf4, 0x4c, 0x09, 0x7a, 0x02, 0x60, 0x07, 0x62, 0x58, 0xc5, 0xb0, 0x65, 0x1d, 0x13, 0xf7,
	0xa0, 0xa9, 0x35, 0x7f, 0xc7, 0xb0, 0x66, 0x3b, 0x8d, 0x9e, 0x29, 0xd1, 0x7f, 0x82, 0xce, 0xe1,
	0x62, 0xb4, 0x0b, 0x95, 0x4f, 0xb6, 0x71, 0x09, 0x52, 0xa6, 0xa4, 0x23, 0xa8, 0xe5, 0xa1, 0xd0,
	0x0c, 0xd7, 0x6a, 0x4f, 0xe8, 0xe1, 0x69, 0xb6, 0x35, 0xb0, 0xc0, 0xb5, 0x73, 0x49, 0x86, 0xb7,
	0x50, 0xb7, 0x26, 0xa5, 0x50, 0x35, 0xeb, 0x95, 0xa3, 0xb0, 0xa6, 0x47, 0xe0, 0xf0, 0xb4, 0xbc,
	0xcf, 0xe1, 0xa9, 0x61, 0x52, 0xa9, 0x32, 0x3c, 0xa7, 0x16, 0x60, 0x3d, 0x1d, 0x6e, 0x0b, 0x8f,
	0xfc, 0x14, 0x1e, 0xf9, 0x2d, 0x3c, 0xf2, 0xbd, 0xf3, 0xc8, 0x76, 0xe7, 0x91, 0xd7, 0x8e, 0x7f,
	0xb3, 0x7f, 0x5e, 0x54, 0xc7, 0x17, 0x5c, 0xfc, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x0c, 0xa5,
	0xca, 0xd1, 0x01, 0x00, 0x00,
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.UuidUrl) > 0 {
		i -= len(m.UuidUrl)
		copy(dAtA[i:], m.UuidUrl)
		i = encodeVarintVboxmanage(dAtA, i, uint64(len(m.UuidUrl)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ServerUrl) > 0 {
		i -= len(m.ServerUrl)
		copy(dAtA[i:], m.ServerUrl)
		i = encodeVarintVboxmanage(dAtA, i, uint64(len(m.ServerUrl)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NameUrl) > 0 {
		i -= len(m.NameUrl)
		copy(dAtA[i:], m.NameUrl)
		i = encodeVarintVboxmanage(dAtA, i, uint64(len(m.NameUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PollingInterval) > 0 {
		i -= len(m.PollingInterval)
		copy(dAtA[i:], m.PollingInterval)
		i = encodeVarintVboxmanage(dAtA, i, uint64(len(m.PollingInterval)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Servers) > 0 {
		for k := range m.Servers {
			v := m.Servers[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintVboxmanage(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintVboxmanage(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintVboxmanage(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Server) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Server) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Server) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Port != 0 {
		i = encodeVarintVboxmanage(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Ip) > 0 {
		i -= len(m.Ip)
		copy(dAtA[i:], m.Ip)
		i = encodeVarintVboxmanage(dAtA, i, uint64(len(m.Ip)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintVboxmanage(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVboxmanage(dAtA []byte, offset int, v uint64) int {
	offset -= sovVboxmanage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Servers) > 0 {
		for k, v := range m.Servers {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovVboxmanage(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovVboxmanage(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovVboxmanage(uint64(mapEntrySize))
		}
	}
	l = len(m.PollingInterval)
	if l > 0 {
		n += 1 + l + sovVboxmanage(uint64(l))
	}
	l = len(m.NameUrl)
	if l > 0 {
		n += 1 + l + sovVboxmanage(uint64(l))
	}
	l = len(m.ServerUrl)
	if l > 0 {
		n += 1 + l + sovVboxmanage(uint64(l))
	}
	l = len(m.UuidUrl)
	if l > 0 {
		n += 1 + l + sovVboxmanage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Server) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovVboxmanage(uint64(l))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovVboxmanage(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovVboxmanage(uint64(m.Port))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVboxmanage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVboxmanage(x uint64) (n int) {
	return sovVboxmanage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVboxmanage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Servers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVboxmanage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVboxmanage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Servers == nil {
				m.Servers = make(map[string]*Server)
			}
			var mapkey string
			var mapvalue *Server
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowVboxmanage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVboxmanage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthVboxmanage
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthVboxmanage
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowVboxmanage
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthVboxmanage
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthVboxmanage
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Server{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipVboxmanage(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthVboxmanage
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Servers[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollingInterval", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVboxmanage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVboxmanage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PollingInterval = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NameUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVboxmanage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVboxmanage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NameUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVboxmanage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVboxmanage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UuidUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVboxmanage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVboxmanage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UuidUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVboxmanage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVboxmanage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Server) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVboxmanage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Server: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Server: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVboxmanage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVboxmanage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVboxmanage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVboxmanage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVboxmanage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVboxmanage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVboxmanage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVboxmanage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVboxmanage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthVboxmanage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVboxmanage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVboxmanage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVboxmanage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVboxmanage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVboxmanage = fmt.Errorf("proto: unexpected end of group")
)
