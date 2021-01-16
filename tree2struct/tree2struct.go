package tree2struct

type SequenceOfCharactersCall_t struct {
	SequenceOfCharactersKeyword string
	Argument                    string
}

type SetOfCharactersCall_t struct {
	SetOfCharactersKeyword string
	Argument               string
}

type SetOfNotCharactersCall_t struct {
	SetOfNotCharactersKeyword string
	Argument                  string
}

type ExpressionStatement_t struct {
	ExpressionName     string
	ExpressionArgument *ExpressionArgument_t
}

type VarStatement_t struct {
	Varkeyword  string
	Expressions []*ExpressionStatement_t
}

type PackageStatement_t struct {
	PackageKeyword string
	PackageName    string
}

type ImportStatement_t struct {
	ImportKeyword        string
	ImportAccessOperator string
	ImportURL            string
}

type GogexFile_t struct {
	PackageStatement *PackageStatement_t
	ImportStatement  *ImportStatement_t
	VarStatement     *VarStatement_t
}

type ExpressionArgument_t struct {
	ExpressionName           string
	SequenceCall             *SequenceCall_t
	SequenceOfCharactersCall *SequenceOfCharactersCall_t
	SetOfCharactersCall      *SetOfCharactersCall_t
	SetOfNotCharactersCall   *SetOfNotCharactersCall_t
	SetCall                  *SetCall_t
	RangeCall                *RangeCall_t
	LabelCall                *LabelCall_t
}

type LabelCall_t struct {
	LabelKeyword       string
	ExpressionArgument *ExpressionArgument_t
	Labels             []string
}

type RangeCall_t struct {
	RangeKeyword       string
	ExpressionArgument *ExpressionArgument_t
	MinValue           int
	MaxValue           int
}

type SetCall_t struct {
	SetKeyword          string
	ExpressionArguments []*ExpressionArgument_t
}

type SequenceCall_t struct {
	SequenceKeyword     string
	ExpressionArguments []*ExpressionArgument_t
}
