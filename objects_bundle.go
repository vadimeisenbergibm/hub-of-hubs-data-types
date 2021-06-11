package datatypes

func NewObjectBundle() *ObjectsBundle {
	return &ObjectsBundle {
		Objects:        make([]interface{}, 0),
		DeletedObjects: make([]interface{}, 0),
	}
}

type ObjectsBundle struct {
	Objects        []interface{} `json:"objects"`
	DeletedObjects []interface{} `json:"deletedObjects"`
}
