package declarations

import (
	"sort"
)

// Function argument declaration.
type FunctionArgument struct {
	// Index.
	Index uint

	// Name.
	Name string

	// Type.
	Type Type
}

// Function declaration.
type Function struct {
	// Service name.
	Name string

	// Documentation paragraphs.
	Documentation []string

	// FunctionArguments.
	//
	// Do not modify this slice directly. Always use AddFunctionArgument.
	Arguments []*FunctionArgument

	// Return type.
	//
	// If no return type is defined, this is considered a void function.
	ReturnType Type

	// FunctionArgument name mapping.
	argumentNameMapping map[string]*FunctionArgument

	// FunctionArgument index mapping.
	argumentIndexMapping map[uint]*FunctionArgument
}

// New struct declaration.
func NewFunction(name string, documentation []string) *Function {
	return &Function{
		Name:                 name,
		Documentation:        documentation,
		Arguments:            []*FunctionArgument{},
		argumentNameMapping:  map[string]*FunctionArgument{},
		argumentIndexMapping: map[uint]*FunctionArgument{},
	}
}

// Add a argument to a function declaration.
//
// The caller is expected to have validated that neither the name nor index are
// in use before calling AddFunctionArgument.
func (s *Function) AddArgument(index uint, name string, argumentType Type) {
	argument := &FunctionArgument{
		Index: index,
		Name:  name,
		Type:  argumentType,
	}

	s.Arguments = append(s.Arguments, argument)
	s.argumentNameMapping[name] = argument
	s.argumentIndexMapping[index] = argument
}

// Determine if a argument index is in use.
func (s *Function) ArgumentIndexInUse(index uint) bool {
	_, inUse := s.argumentIndexMapping[index]
	return inUse
}

// Determine if a argument name is in use.
func (s *Function) ArgumentNameInUse(name string) bool {
	_, inUse := s.argumentNameMapping[name]
	return inUse
}

// Functions by name.
type functionsByName []*Function

func (l functionsByName) Len() int {
	return len(l)
}

func (l functionsByName) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l functionsByName) Less(i, j int) bool {
	return l[i].Name < l[j].Name
}

// Function arguments by index.
type functionArgumentsByIndex []*FunctionArgument

func (l functionArgumentsByIndex) Len() int {
	return len(l)
}

func (l functionArgumentsByIndex) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l functionArgumentsByIndex) Less(i, j int) bool {
	return l[i].Index < l[j].Index
}

// Sorted list of function arguments by index.
func (f *Function) ArgumentsSortedByIndex() []*FunctionArgument {
	unsorted := make([]*FunctionArgument, len(f.Arguments))

	idx := 0
	for _, exc := range f.Arguments {
		unsorted[idx] = exc
		idx++
	}

	sort.Sort(functionArgumentsByIndex(unsorted))

	return unsorted
}

// Minimum length of deserialized array.
func (f *Function) MinimumDeserializedLength() (minimum int) {
	minIndex := uint(0)

	for _, arg := range f.Arguments {
		if !arg.Type.Nilable() && minIndex < arg.Index {
			minIndex = arg.Index
		}
	}

	return int(minIndex)
}

// Length of serialized array.
func (f *Function) SerializedLength() (length int) {
	maxIndex := uint(0)

	for _, arg := range f.Arguments {
		if arg.Index > maxIndex {
			maxIndex = arg.Index
		}
	}

	return int(maxIndex)
}
