package analyzer

func SyntaxAnalysis(tokens []Token) SyntaxTree {
	if len(tokens) == 0 {
		return SyntaxTree{Root: SyntaxNode{Type: ProgramNode}}
	}

	current := 0

	var parseBlock func() SyntaxNode
	var parseStatement func() SyntaxNode
	var parseExpression func() SyntaxNode

	parseBlock = func() SyntaxNode {
		var block SyntaxNode
		block.Type = BlockNode

		// Saltar la llave de apertura si existe
		if current < len(tokens) && tokens[current].Value == "{" {
			current++
		}

		// Parsear todas las declaraciones dentro del bloque
		for current < len(tokens) && tokens[current].Value != "}" {
			stmt := parseStatement()
			if stmt.Type != "" { // Ignorar nodos vacíos
				block.Children = append(block.Children, stmt)
			}

			// Saltar punto y coma si existe
			if current < len(tokens) && tokens[current].Value == ";" {
				current++
			}
		}

		// Saltar la llave de cierre si existe
		if current < len(tokens) && tokens[current].Value == "}" {
			current++
		}

		return block
	}

	parseStatement = func() SyntaxNode {
		if current >= len(tokens) {
			return SyntaxNode{} // Nodo vacío
		}

		token := tokens[current]

		switch token.Type {
		case Keyword:
			switch token.Value {
			case "echo":
				current++
				expr := parseExpression()
				return SyntaxNode{
					Type:     EchoNode,
					Children: []SyntaxNode{expr},
					Line:     token.Line,
				}
			case "if":
				current++
				cond := parseExpression()
				thenBlock := parseBlock()

				var elseBlock SyntaxNode
				if current < len(tokens) && tokens[current].Value == "else" {
					current++
					elseBlock = parseBlock()
				}

				return SyntaxNode{
					Type:     IfNode,
					Children: []SyntaxNode{cond, thenBlock, elseBlock},
					Line:     token.Line,
				}
			default:
				current++
				return SyntaxNode{
					Type: SyntaxNodeType(token.Value + "_stmt"),
					Line: token.Line,
				}
			}
		default:
			return parseExpression()
		}
	}

	parseExpression = func() SyntaxNode {
		if current >= len(tokens) {
			return SyntaxNode{} // Nodo vacío
		}

		token := tokens[current]
		current++

		switch token.Type {
		case Identifier:
			return SyntaxNode{
				Type:  VariableNode,
				Value: token.Value,
				Line:  token.Line,
			}
		case Literal:
			return SyntaxNode{
				Type:  LiteralNode,
				Value: token.Value,
				Line:  token.Line,
			}
		case Operator:
			left := parseExpression()
			right := parseExpression()
			return SyntaxNode{
				Type:     BinaryOpNode,
				Value:    token.Value,
				Children: []SyntaxNode{left, right},
				Line:     token.Line,
			}
		default:
			return SyntaxNode{
				Type:  SyntaxNodeType("unknown"),
				Value: token.Value,
				Line:  token.Line,
			}
		}
	}

	// Construir el árbol de sintaxis
	root := SyntaxNode{Type: ProgramNode}
	for current < len(tokens) {
		stmt := parseStatement()
		if stmt.Type != "" { // Ignorar nodos vacíos
			root.Children = append(root.Children, stmt)
		}

		// Saltar punto y coma si existe
		if current < len(tokens) && tokens[current].Value == ";" {
			current++
		}
	}

	return SyntaxTree{Root: root}
}
