package config

import (
	"fmt"
	"strings"
	"time"
)

// MockConfig implements config.KeyedReader
type MockConfig struct {
	expectations map[string][]*Expectation
}

// Expectation models an expectation
type Expectation struct {
	value interface{}
}

// MakeMockConfig makes a MockConfig, which implements config.KeyedReader
func MakeMockConfig() *MockConfig {
	return &MockConfig{
		expectations: make(map[string][]*Expectation),
	}
}

// Get gets a generic property, returning an interface{}
func (m *MockConfig) Get(prop string) interface{} {
	return m.mustUseExpectation("Get", prop).value
}

// GetBool gets a bool value
func (m *MockConfig) GetBool(prop string) bool {
	e := m.mustUseExpectation("GetBool", prop)
	v, ok := e.value.(bool)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to a bool (this is a bug)", e.value))
	}
	return v
}

// GetDuration gets a time.Duration value
func (m *MockConfig) GetDuration(prop string) time.Duration {
	e := m.mustUseExpectation("GetDuration", prop)
	v, ok := e.value.(time.Duration)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to time.Duration (this is a bug)",
			e.value))
	}
	return v
}

// GetFloat64 gets a float64 value
func (m *MockConfig) GetFloat64(prop string) float64 {
	e := m.mustUseExpectation("GetFloat64", prop)
	v, ok := e.value.(float64)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to float64 (this is a bug)", e.value))
	}
	return v
}

// GetInt gets an int value
func (m *MockConfig) GetInt(prop string) int {
	e := m.mustUseExpectation("GetInt", prop)
	v, ok := e.value.(int)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to int (this is a bug)", e.value))
	}
	return v
}

// GetInt64 gets an int64 value
func (m *MockConfig) GetInt64(prop string) int64 {
	e := m.mustUseExpectation("GetInt64", prop)
	v, ok := e.value.(int64)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to int64 (this is a bug)", e.value))
	}
	return v
}

// GetSizeInBytes gets a uint value
func (m *MockConfig) GetSizeInBytes(prop string) uint {
	e := m.mustUseExpectation("GetSizeInBytes", prop)
	v, ok := e.value.(uint)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to uint (this is a bug)", e.value))
	}
	return v
}

// GetString murders your family... or gets a string value, who knows really
func (m *MockConfig) GetString(prop string) string {
	e := m.mustUseExpectation("GetString", prop)
	v, ok := e.value.(string)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to string (this is a bug)", e.value))
	}
	return v
}

// GetStringMap gets a map[string]interface{}
func (m *MockConfig) GetStringMap(prop string) map[string]interface{} {
	e := m.mustUseExpectation("GetStringMap", prop)
	v, ok := e.value.(map[string]interface{})
	if !ok {
		panic(fmt.Errorf(
			`Unable to coerce expected value %v
			 to map[string]interface{} (this is a bug)`, e.value))
	}
	return v
}

// GetStringMapString gets a map[string]string
func (m *MockConfig) GetStringMapString(prop string) map[string]string {
	e := m.mustUseExpectation("GetStringMapString", prop)
	v, ok := e.value.(map[string]string)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to map[string]string (this is a bug)",
			e.value))
	}
	return v
}

// GetStringMapStringSlice gets a map[string][]string
func (m *MockConfig) GetStringMapStringSlice(prop string) map[string][]string {
	e := m.mustUseExpectation("GetStringMapStringSlice", prop)
	v, ok := e.value.(map[string][]string)
	if !ok {
		panic(fmt.Errorf(
			`Unable to coerce expected value %v
			 to map[string][]string (this is a bug)`, e.value))
	}
	return v
}

// GetStringSlice gets a []string
func (m *MockConfig) GetStringSlice(prop string) []string {
	e := m.mustUseExpectation("GetStringSlice", prop)
	v, ok := e.value.([]string)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to []string (this is a bug)",
			e.value))
	}
	return v
}

// GetTime gets a time.Time
func (m *MockConfig) GetTime(prop string) time.Time {
	e := m.mustUseExpectation("GetTime", prop)
	v, ok := e.value.(time.Time)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to time.Time (this is a bug)",
			e.value))
	}
	return v
}

// IsSet returns a bool
func (m *MockConfig) IsSet(prop string) bool {
	e := m.mustUseExpectation("IsSet", prop)
	v, ok := e.value.(bool)
	if !ok {
		panic(fmt.Errorf(
			"Unable to coerce expected value %v to bool (this is a bug)", e.value))
	}
	return v
}

// ExpectGet adds an expectation for a Get call
func (m *MockConfig) ExpectGet(prop string, value interface{}) *Expectation {
	e := m.addExpectation("Get", prop)
	e.value = value
	return e
}

// ExpectGetBool adds an expectation for a GetBool call
func (m *MockConfig) ExpectGetBool(prop string, value bool) *Expectation {
	e := m.addExpectation("GetBool", prop)
	e.value = value
	return e
}

// ExpectGetDuration adds an expectation for a GetDuration call
func (m *MockConfig) ExpectGetDuration(
	prop string, value time.Duration,
) *Expectation {
	e := m.addExpectation("GetDuration", prop)
	e.value = value
	return e
}

// ExpectGetFloat64 adds an expectation for a GetFloat64 call
func (m *MockConfig) ExpectGetFloat64(prop string, value float64) *Expectation {
	e := m.addExpectation("GetFloat64", prop)
	e.value = value
	return e
}

// ExpectGetInt adds an expectation for a GetInt call
func (m *MockConfig) ExpectGetInt(prop string, value int) *Expectation {
	e := m.addExpectation("GetInt", prop)
	e.value = value
	return e
}

// ExpectGetInt64 adds an expectation for a GetInt64 call
func (m *MockConfig) ExpectGetInt64(prop string, value int64) *Expectation {
	e := m.addExpectation("GetInt64", prop)
	e.value = value
	return e
}

// ExpectGetSizeInBytes adds an expectation for a GetSizeInBytes call
func (m *MockConfig) ExpectGetSizeInBytes(
	prop string, value uint,
) *Expectation {
	e := m.addExpectation("GetSizeInBytes", prop)
	e.value = value
	return e
}

// ExpectGetString adds an expectation for a GetString call
func (m *MockConfig) ExpectGetString(prop string, value string) *Expectation {
	e := m.addExpectation("GetString", prop)
	e.value = value
	return e
}

// ExpectGetStringMap adds an expectation for a GetStringMap call
func (m *MockConfig) ExpectGetStringMap(
	prop string, value map[string]interface{},
) *Expectation {
	e := m.addExpectation("GetStringMap", prop)
	e.value = value
	return e
}

// ExpectGetStringMapString adds an expectation for a GetStringMapString call
func (m *MockConfig) ExpectGetStringMapString(
	prop string, value map[string]string,
) *Expectation {
	e := m.addExpectation("GetStringMapString", prop)
	e.value = value
	return e
}

// ExpectGetStringMapStringSlice adds an expectation for
// a GetStringMapStringSlice call
func (m *MockConfig) ExpectGetStringMapStringSlice(
	prop string, value map[string][]string,
) *Expectation {
	e := m.addExpectation("GetStringMapStringSlice", prop)
	e.value = value
	return e
}

// ExpectGetStringSlice adds an expectation for a GetStringSlice call
func (m *MockConfig) ExpectGetStringSlice(
	prop string, value []string,
) *Expectation {
	e := m.addExpectation("GetStringSlice", prop)
	e.value = value
	return e
}

// ExpectGetTime adds an expectation for a GetTime call
func (m *MockConfig) ExpectGetTime(prop string, value time.Time) *Expectation {
	e := m.addExpectation("GetTime", prop)
	e.value = value
	return e
}

// ExpectIsSet adds an expectation for a IsSet call
func (m *MockConfig) ExpectIsSet(prop string, value bool) *Expectation {
	e := m.addExpectation("GetIsSet", prop)
	e.value = value
	return e
}

// ExpectationFailure returns an error if any expectations are umet.
func (m *MockConfig) ExpectationFailure() error {
	for key, es := range m.expectations {
		for _, e := range es {
			return fmt.Errorf("Unmet expectation: %s with value %v", key, e.value)
		}
	}
	return nil
}

// MustMeetExpectations panics if the expectations are now met.
func (m *MockConfig) MustMeetExpectations() {
	err := m.ExpectationFailure()
	if err != nil {
		panic(err)
	}
}

func (m *MockConfig) addExpectation(keySegments ...string) *Expectation {
	e := &Expectation{}
	key := joinKeySegments(keySegments...)
	es, exists := m.expectations[key]
	if exists {
		m.expectations[key] = append(es, e)
	} else {
		m.expectations[key] = []*Expectation{e}
	}
	return e
}

func (m *MockConfig) useExpectation(keySegments ...string) (e *Expectation) {
	key := joinKeySegments(keySegments...)
	es, exists := m.expectations[key]
	if !exists {
		return nil
	}
	switch len(es) {
	case 0:
		return nil
	case 1:
		m.expectations[key] = []*Expectation{}
		return es[0]
	default:
		e, es = es[len(es)-1], es[:len(es)-1]
		m.expectations[key] = es
		return e
	}
}

func (m *MockConfig) mustUseExpectation(
	prop string, otherKeySegments ...string,
) *Expectation {
	allKeySegments := make([]string, len(otherKeySegments)+1)
	allKeySegments[0] = prop
	for i, v := range otherKeySegments {
		allKeySegments[i+1] = v
	}
	e := m.useExpectation(allKeySegments...)
	if e == nil {
		panic(fmt.Errorf(
			"Missing mock config read: %s(%v), add an expectation with Expect%s",
			prop, otherKeySegments, prop))
	}
	return e
}

func joinKeySegments(keySegments ...string) string {
	return strings.Join(keySegments, ".")
}
