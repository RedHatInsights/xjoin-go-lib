package avro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeWrapperMarshalJSONEmpty(t *testing.T) {
	tw := TypeWrapper{}
	b, err := tw.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "", string(b))
}

func TestTypeWrapperMarshalJSONSingleEmpty(t *testing.T) {
	tw := TypeWrapper{Type{}}
	b, err := tw.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "\"\"", string(b))
}

func TestTypeWrapperJSONMarshalingSingle(t *testing.T) {
	tw := TypeWrapper{
		Type{
			Type:             "testtype",
			Name:             "testname",
			Items:            TypeWrapper{Type{}},
			XJoinType:        "testxjointype",
			ConnectVersion:   1,
			ConnectName:      "testconnectname",
			XJoinCase:        "testxjoincase",
			XJoinEnumeration: true,
			XJoinPrimaryKey:  false,
		},
	}
	b, err := tw.MarshalJSON()
	assert.Nil(t, err)

	jsonString := string(b)
	assert.Contains(t, jsonString, "testtype")
	assert.Contains(t, jsonString, "testname")
	assert.Contains(t, jsonString, "testxjointype")
	assert.Contains(t, jsonString, "testconnectname")
	assert.Contains(t, jsonString, "testxjoincase")
	assert.Contains(t, jsonString, "true")

	ntw := TypeWrapper{}
	err = ntw.UnmarshalJSON(b)
	assert.Nil(t, err)
	assert.Equal(t, tw, ntw)
}

func TestTypeWrapperUnmarshalJSONStringType(t *testing.T) {
	tw := TypeWrapper{}
	err := tw.UnmarshalJSON([]byte("[\"null\"]"))
	assert.Nil(t, err)
	assert.Equal(t, 1, len(tw))
	assert.Equal(t, "null", tw[0].Type)
}

func TestTypeWrapperJSONMarshalingMore(t *testing.T) {
	tw := TypeWrapper{
		Type{
			Type:             "testtype",
			Name:             "testname",
			XJoinType:        "testxjointype",
			ConnectVersion:   1,
			ConnectName:      "testconnectname",
			XJoinCase:        "testxjoincase",
			XJoinEnumeration: true,
			XJoinPrimaryKey:  false,
		},
		Type{
			Type: "test2type",
			Name: "test2name",
			Items: TypeWrapper{Type{
				Name:      "abcd",
				XJoinType: "test2subxjointype",
			}},
			XJoinType:        "test2xjointype",
			ConnectVersion:   1,
			ConnectName:      "test2connectname",
			XJoinCase:        "test2xjoincase",
			XJoinEnumeration: true,
			XJoinPrimaryKey:  false,
		},
	}
	b, err := tw.MarshalJSON()
	assert.Nil(t, err)

	jsonString := string(b)
	assert.Contains(t, jsonString, "testtype")
	assert.Contains(t, jsonString, "testname")
	assert.Contains(t, jsonString, "testxjointype")
	assert.Contains(t, jsonString, "testconnectname")
	assert.Contains(t, jsonString, "testxjoincase")
	assert.Contains(t, jsonString, "true")

	assert.Contains(t, jsonString, "test2type")
	assert.Contains(t, jsonString, "test2name")
	assert.Contains(t, jsonString, "test2xjoincase")
	assert.Contains(t, jsonString, "test2connectname")
	assert.Contains(t, jsonString, "test2connectname")
	assert.Contains(t, jsonString, "test2subxjointype")
	assert.Contains(t, jsonString, "abcd")

	ntw := TypeWrapper{}
	err = ntw.UnmarshalJSON(b)
	assert.Nil(t, err)
	assert.Equal(t, len(tw), len(ntw))
	assert.Equal(t, tw, ntw)
}

func TestSchemaGetFieldByName(t *testing.T) {
	s := Schema{
		Fields: []Field{
			{
				Name: "a",
				Type: []Type{
					{
						Fields: []Field{
							{
								Name: "b",
								Type: []Type{
									{
										Fields: []Field{
											{
												Name: "c",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	cField, err := s.GetFieldByName("a.b.c")
	assert.Nil(t, err)
	assert.Equal(t, Field{Name: "c"}, cField)

	expBField := Field{
		Name: "b",
		Type: []Type{
			{
				Fields: []Field{
					{
						Name: "c",
					},
				},
			},
		},
	}
	bField, err := s.GetFieldByName("a.b")
	assert.Nil(t, err)
	assert.Equal(t, expBField, bField)

	unkField, err := s.GetFieldByName("unk")
	assert.NotNil(t, err)
	assert.Equal(t, Field{}, unkField)
}

func TestSchemaAddField(t *testing.T) {
	var err error
	s := Schema{
		Fields: []Field{
			{
				Name: "a",
				Type: []Type{
					{
						Fields: []Field{
							{
								Name: "b",
								// Type: []Type{
								// 	{
								// 		Fields: []Field{
								// 			{
								// 				Name: "c",
								// 			},
								// 		},
								// 	},
								// },
							},
						},
					},
				},
			},
		},
	}
	err = s.AddField("a.b", Field{Name: "c"})

	assert.Nil(t, err)

	assert.Equal(t, 1, len(s.Fields))
	assert.Equal(t, "a", s.Fields[0].Name)

	assert.Equal(t, 1, len(s.Fields[0].Type))
	assert.Equal(t, 2, len(s.Fields[0].Type[0].Fields))
	assert.Equal(t, "b", s.Fields[0].Type[0].Fields[0].Name)
	assert.Equal(t, "c", s.Fields[0].Type[0].Fields[1].Name)
}
