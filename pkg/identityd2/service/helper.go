package metathings_identityd2_service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"

	id_helper "github.com/nayotta/metathings/pkg/common/id"
	pb_helper "github.com/nayotta/metathings/pkg/common/protobuf"
	storage "github.com/nayotta/metathings/pkg/identityd2/storage"
	pb "github.com/nayotta/metathings/pkg/proto/identityd2"
)

const (
	NONEXPIRATION = 100 * 365 * 24 * time.Hour // 100 years
)

func new_token(dom_id, ent_id, cred_id *string, expire time.Duration) *storage.Token {
	id := id_helper.NewId()
	now := time.Now()
	expires_at := now.Add(expire)

	return &storage.Token{
		Id:           &id,
		DomainId:     dom_id,
		EntityId:     ent_id,
		CredentialId: cred_id,
		IssuedAt:     &now,
		ExpiresAt:    &expires_at,
		Text:         &id, // token text is token id now.
	}
}

func generate_secret(siz int32) string {
	buf := make([]byte, siz)

	for i := int32(0); i < siz; i++ {
		buf[i] = byte(rand.Intn(256))
	}

	return base64.StdEncoding.EncodeToString(buf)
}

func must_parse_extra(x map[string]*wrappers.StringValue) string {
	var buf []byte
	var err error

	if x == nil {
		return `{}`
	}

	extra_map := pb_helper.ExtractStringMap(x)
	if buf, err = json.Marshal(extra_map); err != nil {
		return `{}`
	}

	return string(buf)
}

func copy_extra(x *string) map[string]string {
	if x == nil {
		return map[string]string{}
	}

	y := map[string]string{}
	if err := json.Unmarshal([]byte(*x), &y); err != nil {
		y = map[string]string{}
	}

	return y
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
		Extra:    copy_extra(x.Extra),
	}

	return y
}

func copy_action(x *storage.Action) *pb.Action {
	y := &pb.Action{
		Id:          *x.Id,
		Name:        *x.Name,
		Alias:       *x.Alias,
		Description: *x.Description,
		Extra:       copy_extra(x.Extra),
	}

	return y
}

func copy_role(x *storage.Role) *pb.Role {
	y := &pb.Role{
		Id:          *x.Id,
		Name:        *x.Name,
		Alias:       *x.Alias,
		Description: copy_string(x.Description),
		Extra:       copy_extra(x.Extra),
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
		Extra:   copy_extra(x.Extra),
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
		Extra:       copy_extra(x.Extra),
	}

	return y
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

func ensure_get_domain_id(x domain_getter) error {
	if x.GetDomain() == nil || x.GetDomain().GetId() == nil {
		return errors.New("domain.id is empty")
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
