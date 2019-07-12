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

import (
	"strings"
)

func (t *Tree) parseLinkRefDef(line items) bool {
	_, line = line.trimLeft()
	if 1 > len(line) {
		return false
	}

	linkLabel, remains, label := t.parseLinkLabel(line)
	if nil == linkLabel {
		return false
	}

	if itemColon != remains[0].typ {
		return false
	}

	remains = remains[1:]
	whitespaces, remains := remains.trimLeft()
	newlines, _, _ := whitespaces.statWhitespace()
	if 1 < newlines {
		return false
	}
	if 1 > len(remains) {
		remains = t.nextLine()
		if remains.isBlankLine() {
			t.backupLine(remains)
			return false
		}

		_, remains = remains.trimLeft()
	}

	tokens := remains
	linkDest, remains, destination := t.parseLinkDest(tokens)
	if nil == linkDest {
		return false
	}

	link := &Link{&BaseNode{typ: NodeLink}, destination, ""}

	whitespaces, remains = remains.trimLeft()
	if nil == whitespaces {
		return false
	}
	newlines, _, _ = whitespaces.statWhitespace()
	if 1 < newlines {
		return false
	}

	lowerCaseLabel := strings.ToLower(label)

	var nextLine items
	if 1 > len(remains) {
		nextLine = t.nextLine()
		if nextLine.isBlankLine() {
			if _, ok := t.context.LinkRefDef[lowerCaseLabel]; !ok {
				t.context.LinkRefDef[lowerCaseLabel] = link
			}
			return true
		}

		_, nextLine = nextLine.trimLeft()
		remains = nextLine
	}

	tokens = remains
	_, remains, title := t.parseLinkTitle(tokens)
	if !remains.isBlankLine() {
		if nil != nextLine {
			t.backupLine(nextLine)
			if _, ok := t.context.LinkRefDef[lowerCaseLabel]; !ok {
				t.context.LinkRefDef[lowerCaseLabel] = link
			}

			return true
		}
		return false
	}

	link.Title = title
	if _, ok := t.context.LinkRefDef[lowerCaseLabel]; !ok {
		t.context.LinkRefDef[lowerCaseLabel] = link
	}

	return true
}

func (t *Tree) parseLinkText(tokens items) (ret, remains items, text string) {

	return
}

func (t *Tree) parseLinkTitle(tokens items) (ret, remains items, title string) {
	ret, remains, title = t.parseLinkTitleMatch(itemDoublequote, itemDoublequote, tokens)
	if nil == ret {
		ret, remains, title = t.parseLinkTitleMatch(itemSinglequote, itemSinglequote, tokens)
		if nil == ret {
			ret, remains, title = t.parseLinkTitleMatch(itemOpenParen, itemCloseParen, tokens)
		}
	}
	if "" != title {
		title = unescapeString(title)
	}

	return
}

func (t *Tree) parseLinkTitleMatch(opener, closer itemType, tokens items) (ret, remains items, title string) {
	remains = tokens
	length := len(tokens)
	if 2 > length {
		return
	}

	if opener != tokens[0].typ {
		return
	}

	ret = append(ret, tokens[0])
	line := tokens
	close := false
	i := 1
	for {
		token := line[i]
		ret = append(ret, token)
		title += token.val
		if closer == token.typ && !tokens.isBackslashEscape(i) {
			close = true
			title = title[:len(title)-1]
			break
		}
		if token.isNewline() {
			line = t.nextLine()
			if line.isBlankLine() {
				t.backupLine(line)
				break
			}
			i = 0
			continue
		}
		i++
	}

	if !close {
		ret = nil
		title = ""

		return
	}

	remains = tokens[i+1:]

	return
}

func (t *Tree) parseLinkDest(tokens items) (ret, remains items, destination string) {
	ret, remains, destination = t.parseLinkDest1(tokens)
	if nil == ret {
		ret, remains, destination = t.parseLinkDest2(tokens)
	}
	if nil != ret {
		destination = encodeDestination(destination)
	}

	return
}

func (t *Tree) parseLinkDest2(tokens items) (ret, remains items, destination string) {
	remains = tokens
	var leftParens, rightParens int
	i := 0
	length := len(tokens)
	for ; i < length; i++ {
		token := tokens[i]
		ret = append(ret, token)
		destination += token.val
		if itemSpace == token.typ || token.isControl() {
			destination = destination[:len(destination)-1]
			ret = ret[:len(ret)-1]
			break
		}

		if itemOpenParen == token.typ && !tokens.isBackslashEscape(i) {
			leftParens++
		}
		if itemCloseParen == token.typ && !tokens.isBackslashEscape(i) {
			rightParens++
		}
	}

	if leftParens != rightParens {
		ret = nil
		destination = ""
		return
	}

	if length <= i {
		i = length - 1
		destination = destination[:len(destination)-1]
	}

	remains = tokens[i:]

	return
}

func (t *Tree) parseLinkDest1(tokens items) (ret, remains items, destination string) {
	remains = tokens
	length := len(tokens)
	if 2 > length {
		return
	}

	if itemLess != tokens[0].typ {
		return
	}

	close := false
	i := 0
	for ; i < length; i++ {
		token := tokens[i]
		ret = append(ret, token)
		if 0 < i {
			destination += token.val
			if itemLess == token.typ && !tokens.isBackslashEscape(i) {
				ret = nil
				destination = ""
				return
			}
		}

		if itemGreater == token.typ && !tokens.isBackslashEscape(i) {
			close = true
			destination = destination[0 : len(destination)-1]
			break
		}
	}

	if !close {
		ret = nil
		destination = ""

		return
	}

	remains = tokens[i+1:]

	return
}

func (t *Tree) parseLinkLabel(tokens items) (ret, remains items, label string) {
	length := len(tokens)
	if 2 > length {
		return
	}

	if itemOpenBracket != tokens[0].typ {
		return
	}

	close := false
	i := 0
	for ; i < length; i++ {
		token := tokens[i]
		ret = append(ret, token)
		if 0 < i {
			label += token.val
		}

		if itemCloseBracket == token.typ && !tokens.isBackslashEscape(i) {
			close = true
			label = label[0 : len(label)-1]
			break
		}
	}

	if !close || "" == strings.TrimSpace(label) || 999 < len(label) {
		ret = nil
	}

	remains = tokens[i+1:]

	return
}
