package helper

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestGetFieldBytes(t *testing.T) {
	json := []byte(`{"name":"John","age":25}`)
	path := "name"

	value, err := GetFieldBytes(json, path)
	assert.NoError(t, err)
	assert.Equal(t, "John", value)

	// Test with non-existing path
	nonExistingPath := "address"
	_, err = GetFieldBytes(json, nonExistingPath)
	assert.Error(t, err)
	assert.EqualError(t, err, "path not exist : address # raw json is : {\"name\":\"John\",\"age\":25}")
}

func TestGetField(t *testing.T) {
	json := `{"name":"John","age":25}`
	path := "name"

	value, err := GetField(json, path)
	assert.NoError(t, err)
	assert.Equal(t, "John", value)

	// Test with non-existing path
	nonExistingPath := "address"
	_, err = GetField(json, nonExistingPath)
	assert.Error(t, err)
	assert.EqualError(t, err, "path not exist : address # raw json is : {\"name\":\"John\",\"age\":25}")
}

func TestGetResultBytes(t *testing.T) {
	json := []byte(`{"name":"John","age":25}`)
	path := "name"

	result, err := GetResultBytes(json, path)
	assert.NoError(t, err)
	assert.Equal(t, "John", result.String())

	// Test with non-existing path
	nonExistingPath := "address"
	_, err = GetResultBytes(json, nonExistingPath)
	assert.Error(t, err)
	assert.EqualError(t, err, "path not exist : address # raw json is : {\"name\":\"John\",\"age\":25}")
}

func TestGetResult(t *testing.T) {
	json := `{"name":"John","age":25}`
	path := "name"

	result, err := GetResult(json, path)
	assert.NoError(t, err)
	assert.Equal(t, "John", result.String())

	// Test with non-existing path
	nonExistingPath := "address"
	_, err = GetResult(json, nonExistingPath)
	assert.Error(t, err)
	assert.EqualError(t, err, "path not exist : address # raw json is : {\"name\":\"John\",\"age\":25}")
}

// Test structures
type Person struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type ComplexStruct struct {
	Person  Person    `json:"person"`
	Address Address   `json:"address"`
	Tags    []string  `json:"tags"`
	Scores  []int     `json:"scores"`
	Active  bool      `json:"active"`
	Data    SimpleMap `json:"data"`
}

type SimpleMap map[string]interface{}

func TestStructToJSON(t *testing.T) {
	// Create test time
	testTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		input   interface{}
		want    string
		wantErr bool
	}{
		{
			name: "Simple struct",
			input: Person{
				Name:      "John Doe",
				Age:       30,
				Email:     "john@example.com",
				CreatedAt: testTime,
			},
			want:    `{"name":"John Doe","age":30,"email":"john@example.com","created_at":"2024-01-01T00:00:00Z"}`,
			wantErr: false,
		},
		{
			name: "Complex struct",
			input: ComplexStruct{
				Person: Person{
					Name:      "Jane Doe",
					Age:       25,
					Email:     "jane@example.com",
					CreatedAt: testTime,
				},
				Address: Address{
					Street:  "123 Main St",
					City:    "New York",
					Country: "USA",
				},
				Tags:   []string{"tag1", "tag2"},
				Scores: []int{85, 90, 95},
				Active: true,
				Data: SimpleMap{
					"key1": "value1",
					"key2": 123,
				},
			},
			want:    `{"person":{"name":"Jane Doe","age":25,"email":"jane@example.com","created_at":"2024-01-01T00:00:00Z"},"address":{"street":"123 Main St","city":"New York","country":"USA"},"tags":["tag1","tag2"],"scores":[85,90,95],"active":true,"data":{"key1":"value1","key2":123}}`,
			wantErr: false,
		},
		{
			name:    "Nil input",
			input:   nil,
			want:    "null",
			wantErr: false,
		},
		{
			name:    "Empty struct",
			input:   struct{}{},
			want:    "{}",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructToJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StructToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructToPrettyJSON(t *testing.T) {
	testTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	person := Person{
		Name:      "John Doe",
		Age:       30,
		Email:     "john@example.com",
		CreatedAt: testTime,
	}

	got, err := StructToPrettyJSON(person)
	if err != nil {
		t.Errorf("StructToPrettyJSON() error = %v", err)
		return
	}

	// Check if the output is properly formatted
	lines := strings.Split(got, "\n")
	if len(lines) != 6 { // Opening brace, 4 fields, closing brace
		t.Errorf("StructToPrettyJSON() wrong number of lines = %v, want 6", len(lines))
	}

	// Verify the content is still valid JSON
	var decoded Person
	if err := json.Unmarshal([]byte(got), &decoded); err != nil {
		t.Errorf("StructToPrettyJSON() output is not valid JSON: %v", err)
	}
}

func TestJSONToStruct(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		want    Person
		wantErr bool
	}{
		{
			name: "Valid JSON",
			json: `{"name":"John Doe","age":30,"email":"john@example.com","created_at":"2024-01-01T00:00:00Z"}`,
			want: Person{
				Name:      "John Doe",
				Age:       30,
				Email:     "john@example.com",
				CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name:    "Invalid JSON",
			json:    `{"name":"John Doe"`, // Incomplete JSON
			want:    Person{},
			wantErr: true,
		},
		{
			name:    "Empty JSON",
			json:    "{}",
			want:    Person{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Person
			err := JSONToStruct(tt.json, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (got != tt.want) {
				t.Errorf("JSONToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
