package analyzer

import "encoding/json"

type AnalysisResult struct {
	Lexical  []Token         `json:"lexical"`
	Syntax   SyntaxTree      `json:"syntax"`
	Semantic []SemanticIssue `json:"semantic"`
}

func (r AnalysisResult) ToJSON() string {
	jsonData, _ := json.MarshalIndent(r, "", "  ")
	return string(jsonData)
}

func AnalyzePHP(code string) AnalysisResult {
	tokens := LexicalAnalysis(code)
	syntaxTree := SyntaxAnalysis(tokens)
	semanticIssues := SemanticAnalysis(syntaxTree)

	return AnalysisResult{
		Lexical:  tokens,
		Syntax:   syntaxTree,
		Semantic: semanticIssues,
	}
}
