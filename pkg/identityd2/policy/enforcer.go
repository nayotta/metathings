package metathings_identityd2_policy

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
)

type Enforcer interface {
	Enforce(domain, group, subject, object, action interface{}) error
	AddGroup(domain, group string) error
	RemoveGroup(domain, group string) error
	AddSubjectToRole(domain, group, subject, role string) error
	RemoveSubjectFromRole(domain, group, subject, role string) error
	AddObjectToKind(domain, group, object, kind string) error
	RemoveObjectFromKind(domain, group, object, kind string) error
}
