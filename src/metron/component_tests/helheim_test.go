// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package component_test

import (
	"context"
	"plumbing"

	"google.golang.org/grpc/metadata"
)

type mockDopplerIngestorServer struct {
	PusherCalled chan bool
	PusherInput  struct {
		Arg0 chan plumbing.DopplerIngestor_PusherServer
	}
	PusherOutput struct {
		Ret0 chan error
	}
}

func newMockDopplerIngestorServer() *mockDopplerIngestorServer {
	m := &mockDopplerIngestorServer{}
	m.PusherCalled = make(chan bool, 100)
	m.PusherInput.Arg0 = make(chan plumbing.DopplerIngestor_PusherServer, 100)
	m.PusherOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockDopplerIngestorServer) Pusher(arg0 plumbing.DopplerIngestor_PusherServer) error {
	m.PusherCalled <- true
	m.PusherInput.Arg0 <- arg0
	return <-m.PusherOutput.Ret0
}

type mockDopplerIngestor_PusherServer struct {
	SendAndCloseCalled chan bool
	SendAndCloseInput  struct {
		Arg0 chan *plumbing.PushResponse
	}
	SendAndCloseOutput struct {
		Ret0 chan error
	}
	RecvCalled chan bool
	RecvOutput struct {
		Ret0 chan *plumbing.EnvelopeData
		Ret1 chan error
	}
	SendHeaderCalled chan bool
	SendHeaderInput  struct {
		Arg0 chan metadata.MD
	}
	SendHeaderOutput struct {
		Ret0 chan error
	}
	SetTrailerCalled chan bool
	SetTrailerInput  struct {
		Arg0 chan metadata.MD
	}
	ContextCalled chan bool
	ContextOutput struct {
		Ret0 chan context.Context
	}
	SendMsgCalled chan bool
	SendMsgInput  struct {
		M_ chan interface{}
	}
	SendMsgOutput struct {
		Ret0 chan error
	}
	RecvMsgCalled chan bool
	RecvMsgInput  struct {
		M_ chan interface{}
	}
	RecvMsgOutput struct {
		Ret0 chan error
	}
}

func newMockDopplerIngestor_PusherServer() *mockDopplerIngestor_PusherServer {
	m := &mockDopplerIngestor_PusherServer{}
	m.SendAndCloseCalled = make(chan bool, 100)
	m.SendAndCloseInput.Arg0 = make(chan *plumbing.PushResponse, 100)
	m.SendAndCloseOutput.Ret0 = make(chan error, 100)
	m.RecvCalled = make(chan bool, 100)
	m.RecvOutput.Ret0 = make(chan *plumbing.EnvelopeData, 100)
	m.RecvOutput.Ret1 = make(chan error, 100)
	m.SendHeaderCalled = make(chan bool, 100)
	m.SendHeaderInput.Arg0 = make(chan metadata.MD, 100)
	m.SendHeaderOutput.Ret0 = make(chan error, 100)
	m.SetTrailerCalled = make(chan bool, 100)
	m.SetTrailerInput.Arg0 = make(chan metadata.MD, 100)
	m.ContextCalled = make(chan bool, 100)
	m.ContextOutput.Ret0 = make(chan context.Context, 100)
	m.SendMsgCalled = make(chan bool, 100)
	m.SendMsgInput.M_ = make(chan interface{}, 100)
	m.SendMsgOutput.Ret0 = make(chan error, 100)
	m.RecvMsgCalled = make(chan bool, 100)
	m.RecvMsgInput.M_ = make(chan interface{}, 100)
	m.RecvMsgOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockDopplerIngestor_PusherServer) SendAndClose(arg0 *plumbing.PushResponse) error {
	m.SendAndCloseCalled <- true
	m.SendAndCloseInput.Arg0 <- arg0
	return <-m.SendAndCloseOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServer) Recv() (*plumbing.EnvelopeData, error) {
	m.RecvCalled <- true
	return <-m.RecvOutput.Ret0, <-m.RecvOutput.Ret1
}
func (m *mockDopplerIngestor_PusherServer) SendHeader(arg0 metadata.MD) error {
	m.SendHeaderCalled <- true
	m.SendHeaderInput.Arg0 <- arg0
	return <-m.SendHeaderOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServer) SetTrailer(arg0 metadata.MD) {
	m.SetTrailerCalled <- true
	m.SetTrailerInput.Arg0 <- arg0
}
func (m *mockDopplerIngestor_PusherServer) Context() context.Context {
	m.ContextCalled <- true
	return <-m.ContextOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServer) SendMsg(m_ interface{}) error {
	m.SendMsgCalled <- true
	m.SendMsgInput.M_ <- m_
	return <-m.SendMsgOutput.Ret0
}
func (m *mockDopplerIngestor_PusherServer) RecvMsg(m_ interface{}) error {
	m.RecvMsgCalled <- true
	m.RecvMsgInput.M_ <- m_
	return <-m.RecvMsgOutput.Ret0
}
