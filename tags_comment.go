package pongo2

import (
	"bytes"
)

type tagCommentNode struct{}

func (node *tagCommentNode) Execute(ctx *ExecutionContext, buffer *bytes.Buffer) error {
	return nil
}

func tagCommentParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	comment_node := &tagCommentNode{}

	// TODO: Process the endtag's arguments (see django 'comment'-tag documentation)
	_, _, err := doc.WrapUntilTag("endcomment")
	if err != nil {
		return nil, err
	}

	if arguments.Count() != 0 {
		return nil, arguments.Error("Tag 'comment' does not take any argument.", nil)
	}

	return comment_node, nil
}

func init() {
	RegisterTag("comment", tagCommentParser)
}
