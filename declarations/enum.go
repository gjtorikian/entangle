package declarations

// Enumeration value declaration.
type EnumValue struct {
	// Value.
	Value int64

	// Name.
	Name string

	// Documentation paragraphs.
	Documentation []string
}

// Enumeration declaration.
type Enum struct {
	// Struct name.
	Name string

	// Documentation paragraphs.
	Documentation []string

	// Values.
	//
	// Mapping of values to representation.
	Values map[int64]EnumValue
}

// New enumeration declaration.
func NewEnum(name string, documentation []string) *Enum {
	return &Enum{
		Name:          name,
		Documentation: documentation,
		Values:        make(map[int64]EnumValue),
	}
}

// Add a value.
func (e *Enum) AddValue(value int64, name string, documentation []string) {
	e.Values[value] = EnumValue{
		Value:         value,
		Name:          name,
		Documentation: documentation,
	}
}

// Determine if a value is in use.
func (e *Enum) ValueInUse(value int64) bool {
	_, inUse := e.Values[value]
	return inUse
}
