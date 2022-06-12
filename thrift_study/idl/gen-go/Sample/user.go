// Code generated by Thrift Compiler (0.15.0). DO NOT EDIT.

package Sample

import (
	"bytes"
	"context"
	"fmt"
	"time"
	thrift "github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

type Data map[string]string

func DataPtr(v Data) *Data { return &v }

// Attributes:
//  - ID
//  - Name
//  - Address
type User struct {
  ID int32        `thrift:"id,1,required" db:"id" json:"id"`
  Name string     `thrift:"name,2,required" db:"name" json:"name"`
  Address *string `thrift:"address,3" db:"address" json:"address,omitempty"`
}

func NewUser() *User {
  return &User{}
}


func (p *User) GetID() int32 {
  return p.ID
}

func (p *User) GetName() string {
  return p.Name
}

// Address  是一个 option 字段
var User_Address_DEFAULT string
func (p *User) GetAddress() string {
  if !p.IsSetAddress() {
    return User_Address_DEFAULT
  }
return *p.Address
}
func (p *User) IsSetAddress() bool {
  return p.Address != nil
}

func (p *User) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetID bool = false;
  var issetName bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetID = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetName = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ID is not set"));
  }
  if !issetName{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Name is not set"));
  }
  return nil
}

func (p *User)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ID = v
}
  return nil
}

func (p *User)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *User)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Address = &v
}
  return nil
}

func (p *User) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "User"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *User) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "id", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.ID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err) }
  return err
}

func (p *User) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "name", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:name: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:name: ", p), err) }
  return err
}

func (p *User) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetAddress() {
    if err := oprot.WriteFieldBegin(ctx, "address", thrift.STRING, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:address: ", p), err) }
    if err := oprot.WriteString(ctx, string(*p.Address)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.address (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:address: ", p), err) }
  }
  return err
}

func (p *User) Equals(other *User) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.ID != other.ID { return false }
  if p.Name != other.Name { return false }
  if p.Address != other.Address {
    if p.Address == nil || other.Address == nil {
      return false
    }
    if (*p.Address) != (*other.Address) { return false }
  }
  return true
}

func (p *User) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("User(%+v)", *p)
}

// Attributes:
//  - ErrCode
//  - ErrMsg
//  - Data
type Response struct {
  ErrCode int32 `thrift:"errCode,1,required" db:"errCode" json:"errCode"`
  ErrMsg string `thrift:"errMsg,2,required" db:"errMsg" json:"errMsg"`
  Data Data `thrift:"data,3,required" db:"data" json:"data"`
}

func NewResponse() *Response {
  return &Response{}
}


func (p *Response) GetErrCode() int32 {
  return p.ErrCode
}

func (p *Response) GetErrMsg() string {
  return p.ErrMsg
}

func (p *Response) GetData() Data {
  return p.Data
}
func (p *Response) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetErrCode bool = false;
  var issetErrMsg bool = false;
  var issetData bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetErrCode = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(ctx, iprot); err != nil {
          return err
        }
        issetErrMsg = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.MAP {
        if err := p.ReadField3(ctx, iprot); err != nil {
          return err
        }
        issetData = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetErrCode{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ErrCode is not set"));
  }
  if !issetErrMsg{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ErrMsg is not set"));
  }
  if !issetData{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Data is not set"));
  }
  return nil
}

func (p *Response)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ErrCode = v
}
  return nil
}

func (p *Response)  ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(ctx); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.ErrMsg = v
}
  return nil
}

func (p *Response)  ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
  _, _, size, err := iprot.ReadMapBegin(ctx)
  if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
  }
  tMap := make(Data, size)
  p.Data =  tMap
  for i := 0; i < size; i ++ {
var _key0 string
    if v, err := iprot.ReadString(ctx); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _key0 = v
}
var _val1 string
    if v, err := iprot.ReadString(ctx); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _val1 = v
}
    p.Data[_key0] = _val1
  }
  if err := iprot.ReadMapEnd(ctx); err != nil {
    return thrift.PrependError("error reading map end: ", err)
  }
  return nil
}

func (p *Response) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "Response"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
    if err := p.writeField2(ctx, oprot); err != nil { return err }
    if err := p.writeField3(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Response) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "errCode", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:errCode: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.ErrCode)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.errCode (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:errCode: ", p), err) }
  return err
}

func (p *Response) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "errMsg", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:errMsg: ", p), err) }
  if err := oprot.WriteString(ctx, string(p.ErrMsg)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.errMsg (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:errMsg: ", p), err) }
  return err
}

func (p *Response) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "data", thrift.MAP, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:data: ", p), err) }
  if err := oprot.WriteMapBegin(ctx, thrift.STRING, thrift.STRING, len(p.Data)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
  }
  for k, v := range p.Data {
    if err := oprot.WriteString(ctx, string(k)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    if err := oprot.WriteString(ctx, string(v)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
  }
  if err := oprot.WriteMapEnd(ctx); err != nil {
    return thrift.PrependError("error writing map end: ", err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:data: ", p), err) }
  return err
}

func (p *Response) Equals(other *Response) bool {
  if p == other {
    return true
  } else if p == nil || other == nil {
    return false
  }
  if p.ErrCode != other.ErrCode { return false }
  if p.ErrMsg != other.ErrMsg { return false }
  if len(p.Data) != len(other.Data) { return false }
  for k, _tgt := range p.Data {
    _src2 := other.Data[k]
    if _tgt != _src2 { return false }
  }
  return true
}

func (p *Response) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Response(%+v)", *p)
}

type Greeter interface {
  // Parameters:
  //  - User
  SayHello(ctx context.Context, user *User) (_r *Response, _err error)
  // Parameters:
  //  - UID
  GetUser(ctx context.Context, uid int32) (_r *Response, _err error)
}

type GreeterClient struct {
  c thrift.TClient
  meta thrift.ResponseMeta
}

func NewGreeterClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *GreeterClient {
  return &GreeterClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewGreeterClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *GreeterClient {
  return &GreeterClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewGreeterClient(c thrift.TClient) *GreeterClient {
  return &GreeterClient{
    c: c,
  }
}

func (p *GreeterClient) Client_() thrift.TClient {
  return p.c
}

func (p *GreeterClient) LastResponseMeta_() thrift.ResponseMeta {
  return p.meta
}

func (p *GreeterClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
  p.meta = meta
}

// Parameters:
//  - User
func (p *GreeterClient) SayHello(ctx context.Context, user *User) (_r *Response, _err error) {
  var _args3 GreeterSayHelloArgs
  _args3.User = user
  var _result5 GreeterSayHelloResult
  var _meta4 thrift.ResponseMeta
  _meta4, _err = p.Client_().Call(ctx, "SayHello", &_args3, &_result5)
  p.SetLastResponseMeta_(_meta4)
  if _err != nil {
    return
  }
  if _ret6 := _result5.GetSuccess(); _ret6 != nil {
    return _ret6, nil
  }
  return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "SayHello failed: unknown result")
}

// Parameters:
//  - UID
func (p *GreeterClient) GetUser(ctx context.Context, uid int32) (_r *Response, _err error) {
  var _args7 GreeterGetUserArgs
  _args7.UID = uid
  var _result9 GreeterGetUserResult
  var _meta8 thrift.ResponseMeta
  _meta8, _err = p.Client_().Call(ctx, "GetUser", &_args7, &_result9)
  p.SetLastResponseMeta_(_meta8)
  if _err != nil {
    return
  }
  if _ret10 := _result9.GetSuccess(); _ret10 != nil {
    return _ret10, nil
  }
  return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "GetUser failed: unknown result")
}

// 自定义实现的 Processer， 实现 了 TProcessor 接口
type GreeterProcessor struct {
  processorMap map[string]thrift.TProcessorFunction     // 保存了 function name 和 对应的 TProcessorFunction
  handler Greeter                                       // Greeter 接口
}
var  _ thrift.TProcessor = (*GreeterProcessor)(nil)

func (p *GreeterProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *GreeterProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *GreeterProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewGreeterProcessor(handler Greeter) *GreeterProcessor {

  self11 := &GreeterProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self11.processorMap["SayHello"] = &greeterProcessorSayHello{handler:handler}
  self11.processorMap["GetUser"] = &greeterProcessorGetUser{handler:handler}
return self11
}

func (p *GreeterProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
 // 通过 TProtocol 对象读取消息的 seq 和 需要调用的服务func
  name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
  if err2 != nil { return false, thrift.WrapTException(err2) }

  // 如果这个 func 存在，就直接调用 对应的 func 去处理这个请求接口
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }

  iprot.Skip(ctx, thrift.STRUCT)
  iprot.ReadMessageEnd(ctx)

  // 说明要调用的 func 并不存在，写入 这个 异常
  x12 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
  x12.Write(ctx, oprot)
  oprot.WriteMessageEnd(ctx)
  oprot.Flush(ctx)
  return false, x12

}

// 如果调用的是 SayHello 这个方法
type greeterProcessorSayHello struct {
  handler Greeter
}
var  _ thrift.TProcessorFunction = (*greeterProcessorSayHello)(nil)

func (p *greeterProcessorSayHello) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := GreeterSayHelloArgs{}
  var err2 error
  // args.Read(ctx, iprot) 设置 args 这个请求参数 ，如果读取请求参数失败了，就说明有 PROTOCOL_ERROR 错误了
  if err2 = args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "SayHello", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)  // 每 5 微秒 检查一次 连接是否还在建立
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel()    // 如果连接没有建立了，就调用 cancel func 取消当前的上下文
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := GreeterSayHelloResult{}
  var retval *Response
  // 实际调用目标方法，如果出现问题了，说明就服务端内部的问题了
  if retval, err2 = p.handler.SayHello(ctx, args.User); err2 != nil {
    tickerCancel()
    if err2 == thrift.ErrAbandonRequest {
      return false, thrift.WrapTException(err2)
    }
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing SayHello: " + err2.Error())
    oprot.WriteMessageBegin(ctx, "SayHello", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return true, thrift.WrapTException(err2)
  } else {
    result.Success = retval
  }

  // 代码到这 说明 处理这个请求以及结束了，而且没有问题，在这里写入 reply
  tickerCancel()
  if err2 = oprot.WriteMessageBegin(ctx, "SayHello", thrift.REPLY, seqId); err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err != nil {
    return
  }
  return true, err
}

type greeterProcessorGetUser struct {
  handler Greeter
}

func (p *greeterProcessorGetUser) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := GreeterGetUserArgs{}
  var err2 error
  if err2 = args.Read(ctx, iprot); err2 != nil {
    iprot.ReadMessageEnd(ctx)
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
    oprot.WriteMessageBegin(ctx, "GetUser", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return false, thrift.WrapTException(err2)
  }
  iprot.ReadMessageEnd(ctx)

  tickerCancel := func() {}
  // Start a goroutine to do server side connectivity check.
  if thrift.ServerConnectivityCheckInterval > 0 {
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    var tickerCtx context.Context
    tickerCtx, tickerCancel = context.WithCancel(context.Background())
    defer tickerCancel()
    go func(ctx context.Context, cancel context.CancelFunc) {
      ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
      defer ticker.Stop()
      for {
        select {
        case <-ctx.Done():
          return
        case <-ticker.C:
          if !iprot.Transport().IsOpen() {
            cancel()
            return
          }
        }
      }
    }(tickerCtx, cancel)
  }

  result := GreeterGetUserResult{}
  var retval *Response
  if retval, err2 = p.handler.GetUser(ctx, args.UID); err2 != nil {
    tickerCancel()
    if err2 == thrift.ErrAbandonRequest {
      return false, thrift.WrapTException(err2)
    }
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetUser: " + err2.Error())
    oprot.WriteMessageBegin(ctx, "GetUser", thrift.EXCEPTION, seqId)
    x.Write(ctx, oprot)
    oprot.WriteMessageEnd(ctx)
    oprot.Flush(ctx)
    return true, thrift.WrapTException(err2)
  } else {
    result.Success = retval
  }
  tickerCancel()
  if err2 = oprot.WriteMessageBegin(ctx, "GetUser", thrift.REPLY, seqId); err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = thrift.WrapTException(err2)
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - User
type GreeterSayHelloArgs struct {
  User *User `thrift:"user,1,required" db:"user" json:"user"`
}

func NewGreeterSayHelloArgs() *GreeterSayHelloArgs {
  return &GreeterSayHelloArgs{}
}

var GreeterSayHelloArgs_User_DEFAULT *User
func (p *GreeterSayHelloArgs) GetUser() *User {
  if !p.IsSetUser() {
    return GreeterSayHelloArgs_User_DEFAULT
  }
return p.User
}
func (p *GreeterSayHelloArgs) IsSetUser() bool {
  return p.User != nil
}

func (p *GreeterSayHelloArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetUser bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetUser = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetUser{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field User is not set"));
  }
  return nil
}

func (p *GreeterSayHelloArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  p.User = &User{}
  if err := p.User.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.User), err)
  }
  return nil
}

func (p *GreeterSayHelloArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "SayHello_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GreeterSayHelloArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "user", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:user: ", p), err) }
  if err := p.User.Write(ctx, oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.User), err)
  }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:user: ", p), err) }
  return err
}

func (p *GreeterSayHelloArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GreeterSayHelloArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GreeterSayHelloResult struct {
  Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGreeterSayHelloResult() *GreeterSayHelloResult {
  return &GreeterSayHelloResult{}
}

var GreeterSayHelloResult_Success_DEFAULT *Response
func (p *GreeterSayHelloResult) GetSuccess() *Response {
  if !p.IsSetSuccess() {
    return GreeterSayHelloResult_Success_DEFAULT
  }
return p.Success
}
func (p *GreeterSayHelloResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GreeterSayHelloResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *GreeterSayHelloResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  p.Success = &Response{}
  if err := p.Success.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *GreeterSayHelloResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "SayHello_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GreeterSayHelloResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *GreeterSayHelloResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GreeterSayHelloResult(%+v)", *p)
}

// Attributes:
//  - UID
type GreeterGetUserArgs struct {
  UID int32 `thrift:"uid,1,required" db:"uid" json:"uid"`
}

func NewGreeterGetUserArgs() *GreeterGetUserArgs {
  return &GreeterGetUserArgs{}
}


func (p *GreeterGetUserArgs) GetUID() int32 {
  return p.UID
}
func (p *GreeterGetUserArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetUID bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(ctx, iprot); err != nil {
          return err
        }
        issetUID = true
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetUID{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field UID is not set"));
  }
  return nil
}

func (p *GreeterGetUserArgs)  ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(ctx); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.UID = v
}
  return nil
}

func (p *GreeterGetUserArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "GetUser_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GreeterGetUserArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin(ctx, "uid", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:uid: ", p), err) }
  if err := oprot.WriteI32(ctx, int32(p.UID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.uid (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:uid: ", p), err) }
  return err
}

func (p *GreeterGetUserArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GreeterGetUserArgs(%+v)", *p)
}

// Attributes:
//  - Success
type GreeterGetUserResult struct {
  Success *Response `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewGreeterGetUserResult() *GreeterGetUserResult {
  return &GreeterGetUserResult{}
}

var GreeterGetUserResult_Success_DEFAULT *Response
func (p *GreeterGetUserResult) GetSuccess() *Response {
  if !p.IsSetSuccess() {
    return GreeterGetUserResult_Success_DEFAULT
  }
return p.Success
}
func (p *GreeterGetUserResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GreeterGetUserResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(ctx, iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(ctx, fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(ctx, fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(ctx); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(ctx); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *GreeterGetUserResult)  ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
  p.Success = &Response{}
  if err := p.Success.Read(ctx, iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *GreeterGetUserResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin(ctx, "GetUser_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(ctx, oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(ctx); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(ctx); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *GreeterGetUserResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(ctx, oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(ctx); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *GreeterGetUserResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GreeterGetUserResult(%+v)", *p)
}


