package analyzer

type SyntaxNodeType string

const (
	ProgramNode  SyntaxNodeType = "program"
	EchoNode     SyntaxNodeType = "echo"
	IfNode       SyntaxNodeType = "if"
	WhileNode    SyntaxNodeType = "while"
	ForNode      SyntaxNodeType = "for"
	FunctionNode SyntaxNodeType = "function"
	VariableNode SyntaxNodeType = "variable"
	LiteralNode  SyntaxNodeType = "literal"
	BinaryOpNode SyntaxNodeType = "binary_op"
	BlockNode    SyntaxNodeType = "block"
)

type SyntaxNode struct {
	Type     SyntaxNodeType `json:"type"`
	Value    string         `json:"value,omitempty"`
	Children []SyntaxNode   `json:"children,omitempty"`
	Line     int            `json:"line,omitempty"`
}

type SyntaxTree struct {
	Root SyntaxNode `json:"root"`
}

type SemanticIssueType string

const (
	UndefinedVariable SemanticIssueType = "undefined_variable"
	UnusedVariable    SemanticIssueType = "unused_variable"
	TypeMismatch      SemanticIssueType = "type_mismatch"
)

type SemanticIssue struct {
	Type     SemanticIssueType `json:"type"`
	Message  string            `json:"message"`
	Line     int               `json:"line"`
	Severity string            `json:"severity"` // "warning" o "error"
}
