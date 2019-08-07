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
	"fmt"
	"strings"
)

// NewHTMLRenderer 创建一个 HTML 渲染器。
func NewHTMLRenderer() (ret *Renderer) {
	ret = &Renderer{rendererFuncs: map[int]RendererFunc{}}

	// 注册渲染器函数

	ret.rendererFuncs[NodeRoot] = ret.renderRoot
	ret.rendererFuncs[NodeParagraph] = ret.renderParagraph
	ret.rendererFuncs[NodeText] = ret.renderText
	ret.rendererFuncs[NodeCodeSpan] = ret.renderCodeSpan
	ret.rendererFuncs[NodeCodeBlock] = ret.renderCodeBlock
	ret.rendererFuncs[NodeEmphasis] = ret.renderEmphasis
	ret.rendererFuncs[NodeStrong] = ret.renderStrong
	ret.rendererFuncs[NodeBlockquote] = ret.renderBlockquote
	ret.rendererFuncs[NodeHeading] = ret.renderHeading
	ret.rendererFuncs[NodeList] = ret.renderList
	ret.rendererFuncs[NodeListItem] = ret.renderListItem
	ret.rendererFuncs[NodeThematicBreak] = ret.renderThematicBreak
	ret.rendererFuncs[NodeHardBreak] = ret.renderHardBreak
	ret.rendererFuncs[NodeSoftBreak] = ret.renderSoftBreak
	ret.rendererFuncs[NodeHTMLBlock] = ret.renderHTML
	ret.rendererFuncs[NodeInlineHTML] = ret.renderInlineHTML
	ret.rendererFuncs[NodeLink] = ret.renderLink
	ret.rendererFuncs[NodeImage] = ret.renderImage

	return
}

const (

)

func (r *Renderer) renderImage(node Node, entering bool) (WalkStatus, error) {
	n := node.(*Image)
	var out items
	if entering {
		if 0 == r.disableTags {
			out = toItems("<img src=\"")
			r.Write(out)
			out = escapeHTML(toItems(n.Destination))
			r.Write(out)
			out = toItems("\" alt=\"")
			r.Write(out)
		}
		r.disableTags++
		return WalkContinue, nil
	}

	r.disableTags--
	if 0 == r.disableTags {
		out = toItems("\"")
		if "" != n.Title {
			out = append(out, toItems(" title=\"")...)
			out = append(out, escapeHTML(toItems(n.Title))...)
			out = append(out, toItems("\"")...)
		}
		out = append(out, toItems(" />")...)
		r.Write(out)
	}
	return WalkContinue, nil
}

func (r *Renderer) renderLink(node Node, entering bool) (WalkStatus, error) {
	if entering {
		n := node.(*Link)
		attrs := [][]string{{"href", fromItems(escapeHTML(toItems(n.Destination)))}}
		if "" != n.Title {
			attrs = append(attrs, []string{"title", fromItems(escapeHTML(toItems(n.Title)))})
		}
		r.tag("a", attrs, false)

		return WalkContinue, nil
	}

	r.tag("/a", nil, false)

	return WalkContinue, nil
}

func (r *Renderer) renderHTML(node Node, entering bool) (WalkStatus, error) {
	if !entering {
		return WalkContinue, nil
	}

	r.Newline()
	r.Write(node.Value())
	r.Newline()

	return WalkContinue, nil
}

func (r *Renderer) renderInlineHTML(node Node, entering bool) (WalkStatus, error) {
	if !entering {
		return WalkContinue, nil
	}

	r.Write(node.Value())

	return WalkContinue, nil
}

func (r *Renderer) renderRoot(node Node, entering bool) (WalkStatus, error) {
	return WalkContinue, nil
}

func (r *Renderer) renderParagraph(node Node, entering bool) (WalkStatus, error) {
	if grandparent := node.Parent().Parent(); nil != grandparent && NodeList == grandparent.Type() {
		if grandparent.(*List).tight {
			return WalkContinue, nil
		}
	}

	if entering {
		r.Newline()
		r.tag("p", nil, false)
	} else {
		r.tag("/p", nil, false)
		r.Newline()
	}

	return WalkContinue, nil
}

func (r *Renderer) renderText(node Node, entering bool) (WalkStatus, error) {
	if !entering {
		return WalkContinue, nil
	}

	r.Write(escapeHTML(node.Value()))

	return WalkContinue, nil
}

func (r *Renderer) renderCodeSpan(node Node, entering bool) (WalkStatus, error) {
	if entering {
		r.Write(toItems("<code>"))
		r.Write(escapeHTML(node.Value()))
		return WalkSkipChildren, nil
	}
	r.Write(toItems("</code>"))
	return WalkContinue, nil
}

func (r *Renderer) renderCodeBlock(node Node, entering bool) (WalkStatus, error) {
	if entering {
		r.Newline()
		n := node.(*CodeBlock)
		if "" != n.info {
			infoWords := strings.Fields(n.info)
			r.Write(toItems("<pre><code class=\"language-" + infoWords[0] + "\">"))
			r.Write(escapeHTML(n.value))
		} else {
			r.Write(toItems("<pre><code>"))
			r.Write(escapeHTML(n.value))
		}
		return WalkSkipChildren, nil
	}
	r.Write(toItems("</code></pre>"))
	r.Newline()
	return WalkContinue, nil
}

func (r *Renderer) renderEmphasis(node Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("em", nil, false)
	} else {
		r.tag("/em", nil, false)
	}
	return WalkContinue, nil
}

func (r *Renderer) renderStrong(node Node, entering bool) (WalkStatus, error) {
	if entering {
		r.Write(toItems("<strong>"))
		r.Write(node.Value())
	} else {
		r.Write(toItems("</strong>"))
	}
	return WalkContinue, nil
}

func (r *Renderer) renderBlockquote(n Node, entering bool) (WalkStatus, error) {
	if entering {
		r.Newline()
		r.Write(toItems("<blockquote>"))
		r.Newline()
	} else {
		r.Newline()
		r.Write(toItems("</blockquote>"))
		r.Newline()
	}
	return WalkContinue, nil
}

func (r *Renderer) renderHeading(node Node, entering bool) (WalkStatus, error) {
	n := node.(*Heading)
	if entering {
		r.Newline()
		r.Write(toItems("<h" + " 123456"[n.Level:n.Level+1] + ">"))
	} else {
		r.Write(toItems("</h" + " 123456"[n.Level:n.Level+1] + ">"))
		r.Newline()
	}
	return WalkContinue, nil
}

func (r *Renderer) renderList(node Node, entering bool) (WalkStatus, error) {
	n := node.(*List)
	tag := "ul"
	if nil == n.bulletChar {
		tag = "ol"
	}
	if entering {
		r.Newline()
		attrs := [][]string{{"start", fmt.Sprintf("%d", n.start)}}
		if nil == n.bulletChar && 1 != n.start {
			r.tag(tag, attrs, false)
		} else {
			r.tag(tag, nil, false)
		}
		r.Newline()
	} else {
		r.Newline()
		r.tag("/"+tag, nil, false)
		r.Newline()
	}
	return WalkContinue, nil
}

func (r *Renderer) renderListItem(node Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("li", nil, false)
	} else {
		r.tag("/li", nil, false)
		r.Newline()
	}
	return WalkContinue, nil
}

func (r *Renderer) renderThematicBreak(node Node, entering bool) (WalkStatus, error) {
	if entering {
		r.Newline()
		r.tag("hr", nil, true)
		r.Newline()
	}

	return WalkContinue, nil
}

func (r *Renderer) renderHardBreak(node Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("br", nil, true)
		r.Newline()
	}

	return WalkContinue, nil
}

func (r *Renderer) renderSoftBreak(node Node, entering bool) (WalkStatus, error) {
	if entering {
		r.Newline()
	}

	return WalkContinue, nil
}

func (r *Renderer) tag(name string, attrs [][]string, selfclosing bool) {
	if r.disableTags > 0 {
		return
	}

	r.Write(toItems("<" + name))
	if 0 < len(attrs) {
		for _, attr := range attrs {
			r.Write(toItems(" " + attr[0] + "=\"" + attr[1] + "\""))
		}
	}
	if selfclosing {
		r.Write(toItems(" /"))
	}
	r.Write(toItems(">"))
}
