package typescript

var (
	TypeBoolean = &Type{
		Name: "boolean",
		Kind: KindNative,
	}
	TypeNumber = &Type{
		Name: "number",
		Kind: KindNative,
	}
	TypeString = &Type{
		Name: "string",
		Kind: KindNative,
	}
	TypeNull = &Type{
		Name: "null",
		Kind: KindNative,
	}
	TypeUndefined = &Type{
		Name: "undefined",
		Kind: KindNative,
	}
	TypeDate = &Type{
		Name: "Date",
		Kind: KindNative,
	}
)
