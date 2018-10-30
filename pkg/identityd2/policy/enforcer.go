package metathings_identityd2_policy

import "errors"

var (
	ErrPermissionDenied = errors.New("permission denied")
)

type Enforcer interface {
	Enforce(domain, groups, subject, object, action interface{}) error
	AddPolicy(domain, group, subject, object, action interface{}) error
	RemovePolicy(domain, group, subject, object, action interface{}) error
	AddSubjectToRole(domain, group, subject, role interface{}) error
	RemoveSubjectFromRole(domain, group, subject, role interface{}) error
}
