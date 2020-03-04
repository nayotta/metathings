package metathings_evaluatord_sdk

type Resource interface {
	GetId() string
	GetType() string
}

type resource struct {
	id  string
	typ string
}

func (r *resource) GetId() string {
	return r.id
}

func (r *resource) GetType() string {
	return r.typ
}

func NewResource(id, typ string) Resource {
	return &resource{
		id:  id,
		typ: typ,
	}
}
