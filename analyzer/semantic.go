package analyzer

func SemanticAnalysis(tree SyntaxTree) []SemanticIssue {
	var issues []SemanticIssue

	// Analizar variables no definidas
	variables := make(map[string]bool)

	var analyzeNode func(node SyntaxNode)
	analyzeNode = func(node SyntaxNode) {
		switch node.Type {
		case VariableNode:
			if !variables[node.Value] {
				issues = append(issues, SemanticIssue{
					Type:     UndefinedVariable,
					Message:  "Variable no definida: " + node.Value,
					Line:     node.Line,
					Severity: "error",
				})
			}
		case FunctionNode:
			// Analizar parámetros y cuerpo de función
			// (implementación pendiente)
		}

		// Analizar hijos recursivamente
		for _, child := range node.Children {
			analyzeNode(child)
		}
	}

	analyzeNode(tree.Root)

	return issues
}
