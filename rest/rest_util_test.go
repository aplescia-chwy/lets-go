package rest_test

import (
	"errors"
	"github.com/Chewy-Inc/lets-go/rest"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyTestStruct struct {
	ID string `jsonapi:"primary,MyTestStruct"`
	Name string `jsonapi:"attr,name"`
}

type MyBadStruct struct {
	Name string
	Code int
}

func TestSerializeAsJsonApiResponse(t *testing.T) {
	testObj := &MyTestStruct{
		ID: "10",
		Name: "bob",
	}
	output, err := rest.SerializeAsJsonApiDocument(testObj)
	t.Log(output, err)
	output, err = rest.SerializeAsJsonApiDocument(&MyBadStruct{
		Name: "",
		Code: 0,
	})
	t.Log(output,err)
	output, err = rest.SerializeAsJsonApiDocument("hello")
	assert.NotNil(t, err)
	t.Log(output, err)
}

func TestUnmarshalJsonApiDocument(t *testing.T) {
	testObj := &MyTestStruct{
		ID: "10",
		Name: "bob",
	}
	output, _ := rest.SerializeAsJsonApiDocument(testObj)
	toBeDeserialized := new(MyTestStruct)
	err := rest.UnmarshalJsonApiDocument([]byte(output), toBeDeserialized)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(toBeDeserialized)
	assert.Equal(t, toBeDeserialized, testObj)
}

func TestUnmarshalManyJsonApiDocument(t *testing.T) {
	testObj := &MyTestStruct{
		ID: "10",
		Name: "bob",
	}
	testObjTwo := &MyTestStruct{
		ID: "11",
		Name: "bobby",
	}
	var objs []*MyTestStruct
	objs = append(objs,testObj)
	objs = append(objs,testObjTwo)
	objs = append(objs,nil)
	output, err := rest.SerializeAsJsonApiDocument(objs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(output)
	res, err := rest.UnmarshalManyJsonApiDocument([]byte(output), testObjTwo)
	t.Log(len(res))
	resString, err := rest.SerializeAsJsonApiDocument(res)
	t.Log(resString,err)
}

func TestJsonApiErrorResponse(t *testing.T) {
	err := errors.New("This is my test Error")
	res := rest.JsonApiErrorResponse(500, err)
	t.Log(res)
	assert.NotEmpty(t,res)
}

func TestMarshalling(t *testing.T){
	testObj := &MyTestStruct{
		ID: "10",
		Name: "bob",
	}
	output, err := rest.MarshalAsJsonString(testObj)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(output)
	var thingy MyTestStruct
	err = rest.UnmarshalJsonString(output, &thingy)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(thingy)
}
