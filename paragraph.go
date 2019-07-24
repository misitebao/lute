// Lute - A structured markdown engine.
// Copyright (C) 2019-present, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lute

type Paragraph struct {
	*BaseNode

	OpenTag, CloseTag string
}

func (p *Paragraph) Continuation(tokens items) int {
	if tokens.isBlankLine() {
		return 1
	}

	return 0
}

func (p *Paragraph) AcceptLines() bool {
	return true
}

func (p *Paragraph) CanContain(node Node) bool {
	return false
}

func (t *Tree) parseParagraph(tokens items) (ret Node) {
	tokens = tokens.trim()
	p := &Paragraph{&BaseNode{typ: NodeParagraph, tokens: tokens}, "<p>", "</p>"}
	ret = p

	return
}
