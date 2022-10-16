package listwatcher

import (
	"context"
	"github.com/hypergig/etcd-explorer/internal/etcdtree"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	etcd "go.etcd.io/etcd/client/v3"
)

const (
	eventName = "etcEvent"
)

type Service struct {
	ctx    context.Context
	client *etcd.Client
	root   *etcdtree.Root
}

func New() *Service {
	return &Service{}
}

// Model service no purpose but to generate ts for EtcdTree
func (s *Service) Model(_ *etcdtree.Root) {}

func (s *Service) SetCtx(ctx context.Context) {
	s.ctx = ctx
}

func (s *Service) emitEvent() {
	runtime.EventsEmit(s.ctx, eventName, s.root.Tree)
}

func (s *Service) warm() (int64, error) {
	runtime.LogInfo(s.ctx, "warming cache")
	rsp, err := s.client.Get(
		s.ctx,
		"",
		etcd.WithPrefix(),
		etcd.WithKeysOnly(),
	)
	if err != nil {
		return 0, err
	}
	for _, kv := range rsp.Kvs {
		if kv == nil {
			continue
		}
		s.root.Add(string(kv.Key))
	}
	runtime.LogInfo(s.ctx, "warming etcd cache done")
	return rsp.Header.Revision, nil
}

func (s *Service) StartService(endpoint string) error {
	var err error
	s.root = etcdtree.New("/")

	// setup client
	s.client, err = etcd.New(etcd.Config{
		Endpoints: []string{endpoint},
		Context:   s.ctx,
	})
	if err != nil {
		return err
	}

	// warm
	rev, err := s.warm()
	if err != nil {
		return err
	}
	s.emitEvent()

	// start
	go s.startEventService(rev)

	return nil
}

func (s *Service) startEventService(rev int64) {
	defer s.client.Close()
	runtime.LogInfof(s.ctx, "starting etcd event service")

	watch := s.client.Watch(
		s.ctx,
		"",
		etcd.WithPrefix(),
		etcd.WithRev(rev+1),
	)
	for {
		select {
		case <-s.ctx.Done():
			runtime.LogInfo(s.ctx, "context canceled")
			return
		case w, ok := <-watch:
			runtime.LogInfo(s.ctx, "received event")
			if !ok {
				runtime.LogInfo(s.ctx, "etcd watch channel closed")
				return
			}
			for _, event := range w.Events {
				switch event.Type {
				case etcd.EventTypePut:
					s.root.Add(string(event.Kv.Key))
				case etcd.EventTypeDelete:
					panic("not implemented")
				default:
					runtime.LogFatalf(s.ctx, "unrecognized event time: %s", event.Type)
				}
			}
			s.emitEvent()
		}
	}
}
