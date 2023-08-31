// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package simple

import (
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:    "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination",
		Iface:   reflect.TypeOf((*Destination)(nil)).Elem(),
		Impl:    reflect.TypeOf(destination{}),
		Routed:  true,
		NoRetry: []int{2, 3},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return destination_local_stub{impl: impl.(Destination), tracer: tracer, getAllMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination", Method: "GetAll", Remote: false}), getpidMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination", Method: "Getpid", Remote: false}), recordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination", Method: "Record", Remote: false}), routedRecordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination", Method: "RoutedRecord", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return destination_client_stub{stub: stub, getAllMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination", Method: "GetAll", Remote: true}), getpidMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination", Method: "Getpid", Remote: true}), recordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination", Method: "Record", Remote: true}), routedRecordMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination", Method: "RoutedRecord", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return destination_server_stub{impl: impl.(Destination), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return destination_reflect_stub{caller: caller}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:      "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Server",
		Iface:     reflect.TypeOf((*Server)(nil)).Elem(),
		Impl:      reflect.TypeOf(server{}),
		Listeners: []string{"hello"},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return server_local_stub{impl: impl.(Server), tracer: tracer, addressMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Server", Method: "Address", Remote: false}), proxyAddressMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Server", Method: "ProxyAddress", Remote: false}), shutdownMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Server", Method: "Shutdown", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return server_client_stub{stub: stub, addressMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Server", Method: "Address", Remote: true}), proxyAddressMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Server", Method: "ProxyAddress", Remote: true}), shutdownMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Server", Method: "Shutdown", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return server_server_stub{impl: impl.(Server), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return server_reflect_stub{caller: caller}
		},
		RefData: "⟦1e2dce71:wEaVeRlIsTeNeRs:github.com/ServiceWeaver/weaver/weavertest/internal/simple/Server→hello⟧\n",
	})
	codegen.Register(codegen.Registration{
		Name:    "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Source",
		Iface:   reflect.TypeOf((*Source)(nil)).Elem(),
		Impl:    reflect.TypeOf(source{}),
		NoRetry: []int{0},
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return source_local_stub{impl: impl.(Source), tracer: tracer, emitMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Source", Method: "Emit", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return source_client_stub{stub: stub, emitMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/simple/Source", Method: "Emit", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return source_server_stub{impl: impl.(Source), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return source_reflect_stub{caller: caller}
		},
		RefData: "⟦bf914175:wEaVeReDgE:github.com/ServiceWeaver/weaver/weavertest/internal/simple/Source→github.com/ServiceWeaver/weaver/weavertest/internal/simple/Destination⟧\n",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[Destination] = (*destination)(nil)
var _ weaver.InstanceOf[Server] = (*server)(nil)
var _ weaver.InstanceOf[Source] = (*source)(nil)

// weaver.Router checks.
var _ weaver.RoutedBy[destRouter] = (*destination)(nil)
var _ weaver.Unrouted = (*server)(nil)
var _ weaver.Unrouted = (*source)(nil)

// Component "destination", router "destRouter" checks.
type __destination_destRouter_if_youre_seeing_this_you_probably_forgot_to_run_weaver_generate struct {
	destRouter
	__destination_destRouter_embedding
}

type __destination_destRouter_embedding struct{}

func (__destination_destRouter_embedding) GetAll() {}
func (__destination_destRouter_embedding) Getpid() {}
func (__destination_destRouter_embedding) Record() {}

var _ func(_ context.Context, file string, msg string) string = (&destRouter{}).RoutedRecord                 // routed
var _ = (&__destination_destRouter_if_youre_seeing_this_you_probably_forgot_to_run_weaver_generate{}).GetAll // unrouted
var _ = (&__destination_destRouter_if_youre_seeing_this_you_probably_forgot_to_run_weaver_generate{}).Getpid // unrouted
var _ = (&__destination_destRouter_if_youre_seeing_this_you_probably_forgot_to_run_weaver_generate{}).Record // unrouted

// Local stub implementations.

type destination_local_stub struct {
	impl                Destination
	tracer              trace.Tracer
	getAllMetrics       *codegen.MethodMetrics
	getpidMetrics       *codegen.MethodMetrics
	recordMetrics       *codegen.MethodMetrics
	routedRecordMetrics *codegen.MethodMetrics
}

// Check that destination_local_stub implements the Destination interface.
var _ Destination = (*destination_local_stub)(nil)

func (s destination_local_stub) GetAll(ctx context.Context, a0 string) (r0 []string, err error) {
	// Update metrics.
	begin := s.getAllMetrics.Begin()
	defer func() { s.getAllMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "simple.Destination.GetAll", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetAll(ctx, a0)
}

func (s destination_local_stub) Getpid(ctx context.Context) (r0 int, err error) {
	// Update metrics.
	begin := s.getpidMetrics.Begin()
	defer func() { s.getpidMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "simple.Destination.Getpid", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Getpid(ctx)
}

func (s destination_local_stub) Record(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	begin := s.recordMetrics.Begin()
	defer func() { s.recordMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "simple.Destination.Record", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Record(ctx, a0, a1)
}

func (s destination_local_stub) RoutedRecord(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	begin := s.routedRecordMetrics.Begin()
	defer func() { s.routedRecordMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "simple.Destination.RoutedRecord", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.RoutedRecord(ctx, a0, a1)
}

type server_local_stub struct {
	impl                Server
	tracer              trace.Tracer
	addressMetrics      *codegen.MethodMetrics
	proxyAddressMetrics *codegen.MethodMetrics
	shutdownMetrics     *codegen.MethodMetrics
}

// Check that server_local_stub implements the Server interface.
var _ Server = (*server_local_stub)(nil)

func (s server_local_stub) Address(ctx context.Context) (r0 string, err error) {
	// Update metrics.
	begin := s.addressMetrics.Begin()
	defer func() { s.addressMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "simple.Server.Address", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Address(ctx)
}

func (s server_local_stub) ProxyAddress(ctx context.Context) (r0 string, err error) {
	// Update metrics.
	begin := s.proxyAddressMetrics.Begin()
	defer func() { s.proxyAddressMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "simple.Server.ProxyAddress", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.ProxyAddress(ctx)
}

func (s server_local_stub) Shutdown(ctx context.Context) (err error) {
	// Update metrics.
	begin := s.shutdownMetrics.Begin()
	defer func() { s.shutdownMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "simple.Server.Shutdown", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Shutdown(ctx)
}

type source_local_stub struct {
	impl        Source
	tracer      trace.Tracer
	emitMetrics *codegen.MethodMetrics
}

// Check that source_local_stub implements the Source interface.
var _ Source = (*source_local_stub)(nil)

func (s source_local_stub) Emit(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	begin := s.emitMetrics.Begin()
	defer func() { s.emitMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "simple.Source.Emit", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Emit(ctx, a0, a1)
}

// Client stub implementations.

type destination_client_stub struct {
	stub                codegen.Stub
	getAllMetrics       *codegen.MethodMetrics
	getpidMetrics       *codegen.MethodMetrics
	recordMetrics       *codegen.MethodMetrics
	routedRecordMetrics *codegen.MethodMetrics
}

// Check that destination_client_stub implements the Destination interface.
var _ Destination = (*destination_client_stub)(nil)

func (s destination_client_stub) GetAll(ctx context.Context, a0 string) (r0 []string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getAllMetrics.Begin()
	defer func() { s.getAllMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "simple.Destination.GetAll", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_string_4af10117(dec)
	err = dec.Error()
	return
}

func (s destination_client_stub) Getpid(ctx context.Context) (r0 int, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getpidMetrics.Begin()
	defer func() { s.getpidMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "simple.Destination.Getpid", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 1, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.Int()
	err = dec.Error()
	return
}

func (s destination_client_stub) Record(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.recordMetrics.Begin()
	defer func() { s.recordMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "simple.Destination.Record", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += (4 + len(a1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.String(a1)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s destination_client_stub) RoutedRecord(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.routedRecordMetrics.Begin()
	defer func() { s.routedRecordMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "simple.Destination.RoutedRecord", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += (4 + len(a1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.String(a1)

	// Set the shardKey.
	var r destRouter
	shardKey := _hashDestination(r.RoutedRecord(ctx, a0, a1))

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 3, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type server_client_stub struct {
	stub                codegen.Stub
	addressMetrics      *codegen.MethodMetrics
	proxyAddressMetrics *codegen.MethodMetrics
	shutdownMetrics     *codegen.MethodMetrics
}

// Check that server_client_stub implements the Server interface.
var _ Server = (*server_client_stub)(nil)

func (s server_client_stub) Address(ctx context.Context) (r0 string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.addressMetrics.Begin()
	defer func() { s.addressMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "simple.Server.Address", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

func (s server_client_stub) ProxyAddress(ctx context.Context) (r0 string, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.proxyAddressMetrics.Begin()
	defer func() { s.proxyAddressMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "simple.Server.ProxyAddress", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 1, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = dec.String()
	err = dec.Error()
	return
}

func (s server_client_stub) Shutdown(ctx context.Context) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.shutdownMetrics.Begin()
	defer func() { s.shutdownMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "simple.Server.Shutdown", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 2, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type source_client_stub struct {
	stub        codegen.Stub
	emitMetrics *codegen.MethodMetrics
}

// Check that source_client_stub implements the Source interface.
var _ Source = (*source_client_stub)(nil)

func (s source_client_stub) Emit(ctx context.Context, a0 string, a1 string) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.emitMetrics.Begin()
	defer func() { s.emitMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "simple.Source.Emit", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	size += (4 + len(a1))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	enc.String(a1)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][20]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.20.0 (codegen
version v0.20.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type destination_server_stub struct {
	impl    Destination
	addLoad func(key uint64, load float64)
}

// Check that destination_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*destination_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s destination_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetAll":
		return s.getAll
	case "Getpid":
		return s.getpid
	case "Record":
		return s.record
	case "RoutedRecord":
		return s.routedRecord
	default:
		return nil
	}
}

func (s destination_server_stub) getAll(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetAll(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_string_4af10117(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s destination_server_stub) getpid(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Getpid(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Int(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s destination_server_stub) record(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Record(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s destination_server_stub) routedRecord(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()
	var r destRouter
	s.addLoad(_hashDestination(r.RoutedRecord(ctx, a0, a1)), 1.0)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.RoutedRecord(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type server_server_stub struct {
	impl    Server
	addLoad func(key uint64, load float64)
}

// Check that server_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*server_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s server_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Address":
		return s.address
	case "ProxyAddress":
		return s.proxyAddress
	case "Shutdown":
		return s.shutdown
	default:
		return nil
	}
}

func (s server_server_stub) address(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Address(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s server_server_stub) proxyAddress(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.ProxyAddress(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.String(r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s server_server_stub) shutdown(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Shutdown(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type source_server_stub struct {
	impl    Source
	addLoad func(key uint64, load float64)
}

// Check that source_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*source_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s source_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Emit":
		return s.emit
	default:
		return nil
	}
}

func (s source_server_stub) emit(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()
	var a1 string
	a1 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Emit(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type destination_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that destination_reflect_stub implements the Destination interface.
var _ Destination = (*destination_reflect_stub)(nil)

func (s destination_reflect_stub) GetAll(ctx context.Context, a0 string) (r0 []string, err error) {
	err = s.caller("GetAll", ctx, []any{a0}, []any{&r0})
	return
}

func (s destination_reflect_stub) Getpid(ctx context.Context) (r0 int, err error) {
	err = s.caller("Getpid", ctx, []any{}, []any{&r0})
	return
}

func (s destination_reflect_stub) Record(ctx context.Context, a0 string, a1 string) (err error) {
	err = s.caller("Record", ctx, []any{a0, a1}, []any{})
	return
}

func (s destination_reflect_stub) RoutedRecord(ctx context.Context, a0 string, a1 string) (err error) {
	err = s.caller("RoutedRecord", ctx, []any{a0, a1}, []any{})
	return
}

type server_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that server_reflect_stub implements the Server interface.
var _ Server = (*server_reflect_stub)(nil)

func (s server_reflect_stub) Address(ctx context.Context) (r0 string, err error) {
	err = s.caller("Address", ctx, []any{}, []any{&r0})
	return
}

func (s server_reflect_stub) ProxyAddress(ctx context.Context) (r0 string, err error) {
	err = s.caller("ProxyAddress", ctx, []any{}, []any{&r0})
	return
}

func (s server_reflect_stub) Shutdown(ctx context.Context) (err error) {
	err = s.caller("Shutdown", ctx, []any{}, []any{})
	return
}

type source_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that source_reflect_stub implements the Source interface.
var _ Source = (*source_reflect_stub)(nil)

func (s source_reflect_stub) Emit(ctx context.Context, a0 string, a1 string) (err error) {
	err = s.caller("Emit", ctx, []any{a0, a1}, []any{})
	return
}

// Router methods.

// _hashDestination returns a 64 bit hash of the provided value.
func _hashDestination(r string) uint64 {
	var h codegen.Hasher
	h.WriteString(string(r))
	return h.Sum64()
}

// _orderedCodeDestination returns an order-preserving serialization of the provided value.
func _orderedCodeDestination(r string) codegen.OrderedCode {
	var enc codegen.OrderedEncoder
	enc.WriteString(string(r))
	return enc.Encode()
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}
