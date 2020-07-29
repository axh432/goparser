package goparser

/*
import (
	"fmt"
	. "github.com/axh432/gogex"
	"strings"
)

var primatives = Set(
	Label(word, "word"),
	Label(whitespaceNoNewLineBlock, "whitespaceNoNewLineBlock"),
	Label(whitespaceAtLeastOneNewLineBlock, "whitespaceAtLeastOneNewLineBlock"),
	Label(underscore, "underscore"),
	Label(dot, "dot"),
	Label(comma, "comma"),
	Label(openBracket, "openBracket"),
	Label(closedBracket, "closedBracket"),
	Label(quote, `quote`),
	Label(String, "String"))

func ComposeStatesNoOptionals(variants []string) (string, *MatchTree) {

	exp := Range(primatives, 1, -1)

	patterns, tree := variantsToGrammarPattern(variants, exp)

	if tree != nil {
		return "", tree
	}

	return grammarPatternsToStateExpression(patterns), nil
}

func ComposeList(variants []string) (string, *MatchTree) {
	exp := Range(primatives, 1, -1)

	patterns, tree := variantsToGrammarPattern(variants, exp)

	if tree != nil {
		return "", tree
	}

	sb := strings.Builder{}

	baseItem := patterns[0]

	inbetween := grammarPatternSubtract(patterns[1], baseItem)

	sb.WriteString(fmt.Sprintf("Sequence(Range(Sequence(%s,%s), 1, -1), %s)", baseItem, inbetween, baseItem))

	return sb.String(), nil
}

func grammarPatternSubtract(left, right string) string {
	left = strings.ReplaceAll(left, right + ",", "")
	left = strings.ReplaceAll(left, "," + right, "")
	return left
}

func ComposeStatesOptionals(variants []string) ([]string, *MatchTree) {

	exp := Range(primatives, 1, -1)

	patterns, tree := variantsToGrammarPattern(variants, exp)

	if tree != nil {
		return nil, tree
	}

	return patterns, nil
}

func grammarIntersect(left, right string) string {
	splitLeft := strings.Split(grammarDeduplicate(left), ",")
	splitRight := strings.Split(grammarDeduplicate(right), ",")
	sb := strings.Builder{}
	for _, lstr := range splitLeft {
		for _, rstr := range splitRight {
			if lstr == rstr {
				sb.WriteString(lstr)
				sb.WriteString(",")
			}
		}
	}
	result := sb.String()
	if len(result) > 0 {
		result = result[0:len(result)-1]
	}
	return result
}




func grammarDeduplicate(grammar string) string {
	split := strings.Split(grammar, ",")
	count := map[string]int{}
	sb := strings.Builder{}
	for _, str := range split {
		count[str] = 0
	}
	finalIndex := len(split) - 1
	for i, str := range split {
		sb.WriteString(fmt.Sprintf("%s%d", str, count[str]))
		count[str]++
		if i != finalIndex {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func grammarReduplicate(grammar string) string {
	split := strings.Split(grammar, ",")
	sb := strings.Builder{}
	finalIndex := len(split) - 1
	for i, str := range split {
		sb.WriteString(str[0:len(str)-1])
		if i != finalIndex {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func variantsToGrammarPattern(variants []string, exp Expression) ([]string, *MatchTree) {
	patterns := []string{}
	for _, variant := range variants {
		gp, tree := variantToGrammarPattern(variant, exp)
		if tree != nil {
			return nil, tree
		}
		patterns = append(patterns, gp)
	}
	return patterns, nil
}

func variantToGrammarPattern(variant string, exp Expression) (string, *MatchTree) {
	sb := strings.Builder{}
	tree := Match(variant, exp)

	if tree.IsValid == false {
		return "", &tree
	}

	visitor := func(mt *MatchTree) {
		if mt.Label != "" {
			sb.WriteString(mt.Label)
			sb.WriteString(",")
		}
	}
	tree.AcceptVisitor(visitor)
	str := sb.String()
	return str[0 : len(str)-1], nil //trim the stupid extra comma
}

func grammarPatternsToStateExpression(gps []string) string {
	sb := strings.Builder{}

	sb.WriteString("Set(")
	finalIndex := len(gps) - 1

	for i, pattern := range gps {
		sb.WriteString(grammarPatternToExpression(pattern))
		if i != finalIndex {
			sb.WriteString(",")
		}
	}

	sb.WriteString(")")
	return sb.String()

}

func grammarPatternToExpression(pattern string) string {
	splitPattern := strings.Split(pattern, ",")

	sb := strings.Builder{}

	sb.WriteString("Sequence(")

	finalIndex := len(splitPattern) - 1

	for i, part := range splitPattern {
		sb.WriteString(part)
		if i != finalIndex {
			sb.WriteString(",")
		}
	}

	sb.WriteString(")")
	return sb.String()
}*/
