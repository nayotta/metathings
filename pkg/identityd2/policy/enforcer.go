package metathings_identityd2_policy

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
)

type Enforcer interface {
	Enforce(domain, group, subject, object, action interface{}) error
	AddGroup(domain, group string) error
	RemoveGroup(domain, group string) error
	AddSubjectToRole(subject, role string) error
	RemoveSubjectFromRole(subject, role string) error
	AddObjectToKind(object, kind string) error
	RemoveObjectFromKind(object, kind string) error
}
