package metathings_identityd2_policy

type Enforcer interface {
	Enforce(subject, domain, object, action string) bool
}
