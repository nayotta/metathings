package metathings_identityd2_service

import (
	"context"
	"encoding/base64"
	"errors"
	"math/rand"
	"time"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/proto/identityd2"
)

const (
	NONEXPIRATION = 7 * 24 * time.Hour // 7 days
)

func new_token(dom_id, ent_id, cred_id *string, expire time.Duration) *storage.Token {
	id := id_helper.NewId()
	now := time.Now()
	expires_at := now.Add(expire)

	return &storage.Token{
		Id:            &id,
		DomainId:      dom_id,
		EntityId:      ent_id,
		CredentialId:  cred_id,
		IssuedAt:      &now,
		ExpiresAt:     &expires_at,
		ExpiresPeriod: (*int64)(&expire),
		Text:          &id, // token text is token id now.
	}
}

func generate_secret(siz int32) string {
	buf := make([]byte, siz)

	for i := int32(0); i < siz; i++ {
		buf[i] = byte(rand.Intn(256))
	}

	return base64.StdEncoding.EncodeToString(buf)
}

func copy_string(x *string) string {
	if x == nil {
		return ""
	}

	return *x
}

func copy_domain_children(xs []*storage.Domain) []*pb.Domain {
	ys := []*pb.Domain{}

	for _, x := range xs {
		ys = append(ys, &pb.Domain{
			Id: *x.Id,
		})
	}

	return ys
}

func copy_domain(x *storage.Domain) *pb.Domain {
	var parent *pb.Domain = nil
	if x.Parent != nil {
		parent = &pb.Domain{
			Id: *x.Parent.Id,
		}
	}

	y := &pb.Domain{
		Id:       *x.Id,
		Name:     *x.Name,
		Alias:    *x.Alias,
		Parent:   parent,
		Children: copy_domain_children(x.Children),
		Extra:    pb_helper.CopyExtra(x.Extra),
	}

	return y
}

func copy_action(x *storage.Action) *pb.Action {
	y := &pb.Action{
		Id:          *x.Id,
		Name:        *x.Name,
		Alias:       *x.Alias,
		Description: *x.Description,
		Extra:       pb_helper.CopyExtra(x.Extra),
	}

	return y
}

func copy_action_view(x *storage.Action) *pb.Action {
	return &pb.Action{Id: *x.Id}
}

func copy_actions_view(xs []*storage.Action) []*pb.Action {
	ys := []*pb.Action{}
	for _, x := range xs {
		ys = append(ys, copy_action_view(x))
	}

	return ys
}

func copy_role(x *storage.Role) *pb.Role {
	y := &pb.Role{
		Id:          *x.Id,
		Name:        *x.Name,
		Alias:       *x.Alias,
		Actions:     copy_actions_view(x.Actions),
		Description: copy_string(x.Description),
		Extra:       pb_helper.CopyExtra(x.Extra),
	}

	return y
}

func copy_entity(x *storage.Entity) *pb.Entity {
	domains := []*pb.Domain{}
	for _, d := range x.Domains {
		domains = append(domains, &pb.Domain{
			Id: *d.Id,
		})
	}

	groups := []*pb.Group{}
	for _, g := range x.Groups {
		groups = append(groups, &pb.Group{
			Id: *g.Id,
		})
	}

	roles := []*pb.Role{}
	for _, r := range x.Roles {
		roles = append(roles, &pb.Role{
			Id: *r.Id,
		})
	}

	y := &pb.Entity{
		Id:      *x.Id,
		Domains: domains,
		Groups:  groups,
		Roles:   roles,
		Name:    *x.Name,
		Alias:   *x.Alias,
		Extra:   pb_helper.CopyExtra(x.Extra),
	}

	return y
}

func copy_group(x *storage.Group) *pb.Group {
	roles := []*pb.Role{}
	for _, r := range x.Roles {
		roles = append(roles, &pb.Role{
			Id: *r.Id,
		})
	}

	subjects := []*pb.Entity{}
	for _, s := range x.Subjects {
		subjects = append(subjects, &pb.Entity{
			Id: *s.Id,
		})
	}

	objects := []*pb.Entity{}
	for _, o := range x.Objects {
		objects = append(objects, &pb.Entity{
			Id: *o.Id,
		})
	}

	y := &pb.Group{
		Id:          *x.Id,
		Roles:       roles,
		Subjects:    subjects,
		Objects:     objects,
		Name:        *x.Name,
		Alias:       *x.Alias,
		Description: *x.Description,
		Extra:       pb_helper.CopyExtra(x.Extra),
	}

	return y
}

func copy_groups(xs []*storage.Group) []*pb.Group {
	var ys []*pb.Group

	for _, x := range xs {
		ys = append(ys, copy_group(x))
	}

	return ys
}

func copy_credential_with_secret(x *storage.Credential) *pb.Credential {
	roles := []*pb.Role{}
	for _, r := range x.Roles {
		roles = append(roles, &pb.Role{
			Id: *r.Id,
		})
	}

	expires_at := pb_helper.FromTime(*x.ExpiresAt)

	y := &pb.Credential{
		Id: *x.Id,
		Domain: &pb.Domain{
			Id: *x.DomainId,
		},
		Roles: roles,
		Entity: &pb.Entity{
			Id: *x.EntityId,
		},
		Name:        *x.Name,
		Alias:       *x.Alias,
		Secret:      *x.Secret,
		Description: *x.Description,
		ExpiresAt:   &expires_at,
	}

	return y
}

func copy_credential(x *storage.Credential) *pb.Credential {
	cred := copy_credential_with_secret(x)
	cred.Secret = ""
	return cred
}

func copy_token(x *storage.Token) *pb.Token {
	issued_at := pb_helper.FromTime(*x.IssuedAt)
	expires_at := pb_helper.FromTime(*x.ExpiresAt)

	var roles []*pb.Role
	for _, r := range x.Roles {
		roles = append(roles, &pb.Role{
			Id: *r.Id,
		})
	}

	var credential *pb.Credential
	if x.Credential != nil {
		credential = &pb.Credential{
			Id: *x.Credential.Id,
		}
	}

	var groups []*pb.Group
	for _, g := range x.Groups {
		groups = append(groups, &pb.Group{
			Id: *g.Id,
		})
	}

	y := &pb.Token{
		Id:        *x.Id,
		IssuedAt:  &issued_at,
		ExpiresAt: &expires_at,
		Entity: &pb.Entity{
			Id: *x.EntityId,
		},
		Roles:  roles,
		Groups: groups,
		Domain: &pb.Domain{
			Id: *x.DomainId,
		},
		Credential: credential,
		Text:       *x.Text,
	}

	return y
}

func role_in_entity(ent *storage.Entity, role_id string) bool {
	for _, r := range ent.Roles {
		if *r.Id == role_id {
			return true
		}
	}

	return false
}

func role_in_group(grp *storage.Group, role_id string) bool {
	for _, r := range grp.Roles {
		if *r.Id == role_id {
			return true
		}
	}

	return false
}

func domain_in_entity(ent *storage.Entity, dom_id string) bool {
	for _, d := range ent.Domains {
		if *d.Id == dom_id {
			return true
		}
	}

	return false
}

func domain_in_credential(cred *storage.Credential, dom_id string) bool {
	return *cred.Domain.Id == dom_id
}

type entity_getter interface {
	GetEntity() *pb.OpEntity
}

type domain_getter interface {
	GetDomain() *pb.OpDomain
}

type credential_getter interface {
	GetCredential() *pb.OpCredential
}

type role_getter interface {
	GetRole() *pb.OpRole
}

type subject_getter interface {
	GetSubject() *pb.OpEntity
}

type object_getter interface {
	GetObject() *pb.OpEntity
}

type group_getter interface {
	GetGroup() *pb.OpGroup
}

type action_getter interface {
	GetAction() *pb.OpAction
}

type token_getter interface {
	GetToken() *pb.OpToken
}

func ensure_get_domain_id(x domain_getter) error {
	if x.GetDomain() == nil || x.GetDomain().GetId() == nil {
		return errors.New("domain.id is empty")
	}
	return nil
}

func ensure_get_domain_parent_id(x domain_getter) error {
	if dom := x.GetDomain(); dom != nil {
		if par := dom.GetParent(); par != nil {
			if id := par.GetId(); id != nil {
				if s := id.GetValue(); s != "" {
					return nil
				}
			}
		}
	}

	return errors.New("domain.parent.id is empty")
}

func ensure_get_domain_name(x domain_getter) error {
	if x.GetDomain() == nil || x.GetDomain().GetName() == nil {
		return errors.New("domain.name is empty")
	}

	return nil
}

func ensure_get_group_id(x group_getter) error {
	if x.GetGroup() == nil || x.GetGroup().GetId() == nil {
		return errors.New("group.id is empty")
	}
	return nil
}

func ensure_get_subject_id(x subject_getter) error {
	if x.GetSubject() == nil || x.GetSubject().GetId() == nil {
		return errors.New("subject.id is empty")
	}
	return nil
}

func ensure_get_object_id(x object_getter) error {
	if x.GetObject() == nil || x.GetObject().GetId() == nil {
		return errors.New("object.id is empty")
	}
	return nil
}

func ensure_get_credential_id(x credential_getter) error {
	if x.GetCredential() == nil || x.GetCredential().GetId() == nil {
		return errors.New("credential.id is empty")
	}
	return nil
}

func ensure_get_entity_id(x entity_getter) error {
	if x.GetEntity() == nil || x.GetEntity().GetId() == nil {
		return errors.New("entity.id is empty")
	}
	return nil
}

func ensure_get_role_id(x role_getter) error {
	if x.GetRole() == nil || x.GetRole().GetId() == nil {
		return errors.New("role.id is empty")
	}
	return nil
}

func ensure_get_action_id(x action_getter) error {
	if x.GetAction() == nil || x.GetAction().GetId() == nil {
		return errors.New("action.id is empty")
	}
	return nil
}

func ensure_get_action_name(x action_getter) error {
	if x.GetAction() == nil || x.GetAction().GetName() == nil {
		return errors.New("action.name is empty")
	}
	return nil
}

func ensure_group_exists_s(ctx context.Context, s storage.Storage) func(group_getter) error {
	return func(x group_getter) error {
		if exist, err := s.ExistGroup(ctx, x.GetGroup().GetId().GetValue()); err != nil {
			return err
		} else if !exist {
			return errors.New("group not found")
		}
		return nil
	}
}

func ensure_subject_not_exists_in_group_s(ctx context.Context, s storage.Storage) func(subject_getter, group_getter) error {
	return func(x subject_getter, y group_getter) error {
		if exist, err := s.SubjectExistsInGroup(ctx, x.GetSubject().GetId().GetValue(), y.GetGroup().GetId().GetValue()); err != nil {
			return err
		} else if exist {
			return errors.New("subject exists in group")
		}

		return nil
	}
}

func ensure_subject_exists_in_group_s(ctx context.Context, s storage.Storage) func(subject_getter, group_getter) error {
	return func(x subject_getter, y group_getter) error {
		if exist, err := s.SubjectExistsInGroup(ctx, x.GetSubject().GetId().GetValue(), y.GetGroup().GetId().GetValue()); err != nil {
			return err
		} else if !exist {
			return errors.New("subject not exists in group")
		}

		return nil
	}
}

func ensure_object_not_exists_in_group_s(ctx context.Context, s storage.Storage) func(object_getter, group_getter) error {
	return func(x object_getter, y group_getter) error {
		if exist, err := s.ObjectExistsInGroup(ctx, x.GetObject().GetId().GetValue(), y.GetGroup().GetId().GetValue()); err != nil {
			return err
		} else if exist {
			return errors.New("object exists in group")
		}

		return nil
	}
}

func ensure_object_exists_in_group_s(ctx context.Context, s storage.Storage) func(object_getter, group_getter) error {
	return func(x object_getter, y group_getter) error {
		if exist, err := s.ObjectExistsInGroup(ctx, x.GetObject().GetId().GetValue(), y.GetGroup().GetId().GetValue()); err != nil {
			return err
		} else if !exist {
			return errors.New("object not exists in group")
		}

		return nil
	}
}

func ensure_get_token_text(x token_getter) error {
	if x.GetToken() == nil || x.GetToken().GetText() == nil {
		return errors.New("token.text is empty")
	}

	return nil
}
