package metathings_toolkit_tag

type TagToolkit interface {
	Tag(id string, tags []string) error
	Untag(id string, tags []string) error
	Remove(id string) error
	Get(id string) ([]string, error)
	Query(tags []string) ([]string, error)
}

var tag_toolkit_factories map[string]func(...interface{}) (TagToolkit, error)

func register_tag_toolkit_factory(driver string, fty func(...interface{}) (TagToolkit, error)) {
	if tag_toolkit_factories == nil {
		tag_toolkit_factories = make(map[string]func(...interface{}) (TagToolkit, error))
	}

	tag_toolkit_factories[driver] = fty
}

func NewTagToolkit(driver string, args ...interface{}) (TagToolkit, error) {
	fty, ok := tag_toolkit_factories[driver]
	if !ok {
		return nil, ErrUnknownTagToolkitDriver
	}

	tag, err := fty(args...)
	if err != nil {
		return nil, err
	}

	return tag, nil
}
