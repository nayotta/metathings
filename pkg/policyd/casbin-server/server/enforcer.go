// Copyright 2018 The casbin Authors. All Rights Reserved.
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

package server

import (
	"errors"

	"github.com/casbin/casbin"
	"github.com/casbin/casbin/persist"
	pb "github.com/nayotta/metathings/pkg/proto/policyd"
	"golang.org/x/net/context"
)

// Server is used to implement proto.CasbinServer.
type Server struct {
	enforcerMap map[int]*casbin.Enforcer
	adapterMap  map[int]persist.Adapter
}

func NewServer() *Server {
	s := Server{}

	s.enforcerMap = map[int]*casbin.Enforcer{}
	s.adapterMap = map[int]persist.Adapter{}

	return &s
}

func (s *Server) getEnforcer(handle int) (*casbin.Enforcer, error) {
	if _, ok := s.enforcerMap[handle]; ok {
		return s.enforcerMap[handle], nil
	} else {
		return nil, errors.New("enforcer not found")
	}
}

func (s *Server) GetEnforcer(handle int) (*casbin.Enforcer, error) {
	return s.getEnforcer(handle)
}

func (s *Server) getAdapter(handle int) (persist.Adapter, error) {
	if _, ok := s.adapterMap[handle]; ok {
		return s.adapterMap[handle], nil
	} else {
		return nil, errors.New("adapter not found")
	}
}

func (s *Server) addEnforcer(e *casbin.Enforcer) int {
	cnt := len(s.enforcerMap)
	s.enforcerMap[cnt] = e
	return cnt
}

func (s *Server) addAdapter(a persist.Adapter) int {
	cnt := len(s.adapterMap)
	s.adapterMap[cnt] = a
	return cnt
}

func (s *Server) NewEnforcer(ctx context.Context, in *pb.NewEnforcerRequest) (*pb.NewEnforcerReply, error) {
	a, err := s.getAdapter(int(in.AdapterHandle))
	if err != nil {
		return &pb.NewEnforcerReply{Handler: 0}, err
	}

	e := casbin.NewEnforcer(casbin.NewModel(in.ModelText), a)
	h := s.addEnforcer(e)

	return &pb.NewEnforcerReply{Handler: int32(h)}, nil
}

func (s *Server) NewAdapter(ctx context.Context, in *pb.NewAdapterRequest) (*pb.NewAdapterReply, error) {
	a, err := newAdapter(in)
	if err != nil {
		return nil, err
	}

	h := s.addAdapter(a)

	return &pb.NewAdapterReply{Handler: int32(h)}, nil
}

func (s *Server) parseEnforceRequest(in *pb.EnforceRequest) []interface{} {
	params := make([]interface{}, 0, len(in.Params))
	for e := range in.Params {
		params = append(params, in.Params[e])
	}

	return params
}

func (s *Server) enforce(handler int, params []interface{}) (bool, error) {
	e, err := s.getEnforcer(handler)
	if err != nil {
		return false, err
	}

	return e.Enforce(params...), nil
}

func (s *Server) Enforce(ctx context.Context, in *pb.EnforceRequest) (*pb.BoolReply, error) {
	var res bool
	var err error

	params := s.parseEnforceRequest(in)
	if res, err = s.enforce(int(in.EnforcerHandler), params); err != nil {
		return &pb.BoolReply{Res: false}, err
	}

	return &pb.BoolReply{Res: res}, nil
}

func (s *Server) EnforceBucket(ctx context.Context, in *pb.EnforceBucketRequest) (*pb.BoolReply, error) {
	var res bool
	var err error

	for _, req := range in.Requests {
		params := s.parseEnforceRequest(req)
		if res, err = s.enforce(int(req.EnforcerHandler), params); err != nil {
			return &pb.BoolReply{Res: false}, err
		} else if res {
			return &pb.BoolReply{Res: true}, nil
		}
	}

	return &pb.BoolReply{Res: false}, nil
}

func (s *Server) LoadPolicy(ctx context.Context, in *pb.EmptyRequest) (*pb.EmptyReply, error) {
	e, err := s.getEnforcer(int(in.Handler))
	if err != nil {
		return &pb.EmptyReply{}, err
	}

	err = e.LoadPolicy()

	return &pb.EmptyReply{}, err
}

func (s *Server) SavePolicy(ctx context.Context, in *pb.EmptyRequest) (*pb.EmptyReply, error) {
	e, err := s.getEnforcer(int(in.Handler))
	if err != nil {
		return &pb.EmptyReply{}, err
	}

	err = e.SavePolicy()

	return &pb.EmptyReply{}, err
}

func (s *Server) AddPolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	in.PType = "p"

	return s.AddNamedPolicy(ctx, in)
}

func (s *Server) AddNamedPolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	return &pb.BoolReply{Res: e.AddNamedPolicy(in.PType, in.Params)}, err
}

func (s *Server) RemovePolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	res := e.RemovePolicy(in.Params)

	return &pb.BoolReply{Res: res}, err
}

func (s *Server) RemoveNamedPolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	res := e.RemoveNamedPolicy(in.PType, in.Params)

	return &pb.BoolReply{Res: res}, err
}

// RemoveFilteredPolicy removes an authorization rule from the current policy, field filters can be specified.
func (s *Server) RemoveFilteredPolicy(ctx context.Context, in *pb.FilteredPolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	return &pb.BoolReply{Res: e.RemoveFilteredNamedPolicy("p", int(in.FieldIndex), in.FieldValues...)}, nil
}

// RemoveFilteredNamedPolicy removes an authorization rule from the current named policy, field filters can be specified.
func (s *Server) RemoveFilteredNamedPolicy(ctx context.Context, in *pb.FilteredPolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	return &pb.BoolReply{Res: e.RemoveFilteredNamedPolicy(in.PType, int(in.FieldIndex), in.FieldValues...)}, nil
}

// AddGroupingPolicy adds a role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (s *Server) AddGroupingPolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	in.PType = "g"

	return s.AddNamedGroupingPolicy(ctx, in)
}

// AddNamedGroupingPolicy adds a named role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func (s *Server) AddNamedGroupingPolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	return &pb.BoolReply{Res: e.AddNamedGroupingPolicy(in.PType, in.Params)}, nil
}

// RemoveGroupingPolicy removes a role inheritance rule from the current policy.
func (s *Server) RemoveGroupingPolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	return &pb.BoolReply{Res: e.RemoveNamedGroupingPolicy("g", in.Params)}, nil
}

// RemoveNamedGroupingPolicy removes a role inheritance rule from the current named policy.
func (s *Server) RemoveNamedGroupingPolicy(ctx context.Context, in *pb.PolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	return &pb.BoolReply{Res: e.RemoveNamedGroupingPolicy(in.PType, in.Params)}, nil
}

// RemoveFilteredGroupingPolicy removes a role inheritance rule from the current policy, field filters can be specified.
func (s *Server) RemoveFilteredGroupingPolicy(ctx context.Context, in *pb.FilteredPolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	return &pb.BoolReply{Res: e.RemoveFilteredNamedGroupingPolicy("g", int(in.FieldIndex), in.FieldValues...)}, nil
}

// RemoveFilteredNamedGroupingPolicy removes a role inheritance rule from the current named policy, field filters can be specified.
func (s *Server) RemoveFilteredNamedGroupingPolicy(ctx context.Context, in *pb.FilteredPolicyRequest) (*pb.BoolReply, error) {
	e, err := s.getEnforcer(int(in.EnforcerHandler))
	if err != nil {
		return &pb.BoolReply{}, err
	}

	return &pb.BoolReply{Res: e.RemoveFilteredNamedGroupingPolicy(in.PType, int(in.FieldIndex), in.FieldValues...)}, nil
}
