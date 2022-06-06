package types

type nullableDescriber interface {
	Nullable() (bool, bool) // Compatible with sql.ColumnType
}

type databaseTypeNameDescriber interface {
	DatabaseTypeName() string
}

type anyDataTypeProvider interface {
	DataType() interface{}
}
