// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"github.com/88250/lute/ast"
	"testing"

	"github.com/88250/lute"
)

var spinBlockDOMTests = []*parseTest{

	{"23", "<div data-node-id=\"20210510235001-cs6wzxy\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210510235007\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\"><span class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span><img src=\"assets/image-20210510235005-0g043da.png\" data-src=\"assets/image-20210510235005-0g043da.png\" alt=\"image.png\" updated=\"20210510235005\"><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\"></span></span><wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210510235001-cs6wzxy\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210510235007\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\"><span class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span><img src=\"/siyuan/0/测试笔记/assets/image-20210510235005-0g043da.png\" data-src=\"assets/image-20210510235005-0g043da.png\" alt=\"image.png\" updated=\"20210510235007\" /><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\"></span></span><wbr>\n</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"22", "<div data-node-id=\"20210509233323-qdrynbi\" data-node-index=\"1\" data-type=\"NodeTable\" class=\"table\" updated=\"20210509233434\"><div class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></div><div contenteditable=\"true\" spellcheck=\"false\"><wbr><table><thead><tr><th>col1</th><th>col2</th><th>col3</th></tr></thead><tbody><tr><td></td><td></td><td></td></tr><tr><td></td><td></td><td></td></tr></tbody></table></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210509233323-qdrynbi\" data-node-index=\"1\" data-type=\"NodeTable\" class=\"table\" updated=\"20210509233434\"><div class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></div><div contenteditable=\"true\" spellcheck=\"false\"><table><thead><tr><th>col1</th><th>col2</th><th>col3</th></tr></thead><tbody><tr><td> </td><td> </td><td> </td></tr><tr><td> </td><td> </td><td> </td></tr></tbody></table></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"21", "<div data-node-id=\"20210508095314-sbh3v64\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210508115128\"><div contenteditable=\"true\" spellcheck=\"false\">    f<wbr></div><div class=\"protyle-attr\"></div></div>", "<div data-node-id=\"20210508095314-sbh3v64\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210508115128\"><div contenteditable=\"true\" spellcheck=\"false\">f<wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"20", "<div data-node-id=\"20210507213056-5rxw090\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210507224645\" memo=\"bar\"><div contenteditable=\"true\" spellcheck=\"false\">foo<wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"><div class=\"protyle-attr--memo\"><svg><use xlink:href=\"#iconM\"></use></svg>bar</div></div></div>", "<div data-node-id=\"20210507213056-5rxw090\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" memo=\"bar\" updated=\"20210507224645\"><div contenteditable=\"true\" spellcheck=\"false\">foo<wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"><div class=\"protyle-attr--memo\"><svg><use xlink:href=\"#iconM\"></use></svg>bar</div></div></div>"},
	{"19", "<div data-node-id=\"20210506170317-sgimww6\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210506170409\"><div contenteditable=\"true\" spellcheck=\"false\">* [ ] ‸</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-subtype=\"t\" data-node-id=\"20210506170317-sgimww6\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\" updated=\"20210506170409\"><div data-marker=\"*\" data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\"><div class=\"protyle-action protyle-action--task\"><svg><use xlink:href=\"#iconUncheck\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"18", "<div data-node-id=\"20210505171041-me8noe4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210505183238\"><div contenteditable=\"true\" spellcheck=\"false\">foo <img alt=\"huaji\" class=\"emoji\" src=\"http://127.0.0.1:6806/stage/protyle/images/emoji/huaji.gif\" title=\"huaji\"><wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210505171041-me8noe4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210505183238\"><div contenteditable=\"true\" spellcheck=\"false\">foo <img alt=\"huaji\" class=\"emoji\" src=\"http://127.0.0.1:6806/stage/protyle/images/emoji/huaji.gif\" title=\"huaji\" /><wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"17", "<div data-node-id=\"20210504111639-jjqdgx5\" data-node-index=\"1\" data-type=\"NodeTable\" class=\"table\" updated=\"20210504111846\"><div class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></div><div contenteditable=\"true\" spellcheck=\"false\"><table><thead><tr><th>col1</th></tr></thead><tbody><tr><td>foo<br />bar<wbr></td></tr></tbody></table></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210504111639-jjqdgx5\" data-node-index=\"1\" data-type=\"NodeTable\" class=\"table\" updated=\"20210504111846\"><div class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></div><div contenteditable=\"true\" spellcheck=\"false\"><table><thead><tr><th>col1</th></tr></thead><tbody><tr><td>foo<br />bar<wbr></td></tr></tbody></table></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"16", "<div data-node-id=\"20210504103508-y9plmyi\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210504103508\"><div contenteditable=\"true\" spellcheck=\"false\">t<wbr><span data-type=\"a\" data-href=\"bar\" data-title=\"baz\"></span></div><div class=\"protyle-attr\"></div></div>", "<div data-node-id=\"20210504103508-y9plmyi\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210504103508\"><div contenteditable=\"true\" spellcheck=\"false\">t<wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"15", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><u>foo</u></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><u>foo</u></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"14", "<div data-node-id=\"20210502094636-x7veqk8\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210502094911\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\" style=\"display: block;\"><span class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span><img src=\"assets/image-20210502094911-rr579ht.png\" data-src=\"assets/image-20210502094911-rr579ht.png\" alt=\"image.png\" updated=\"20210502094911\" style=\"width: 544px;\"><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\"></span></span></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210502094636-x7veqk8\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210502094911\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\" style=\"display: block;\"><span class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span><img src=\"/siyuan/0/测试笔记/assets/image-20210502094911-rr579ht.png\" data-src=\"assets/image-20210502094911-rr579ht.png\" alt=\"image.png\" updated=\"20210502094911\" style=\"width: 544px;\" /><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\"></span></span>\n</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"13", "<div data-node-id=\"20210502094636-x7veqk8\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210502094911\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\"><span class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span><img src=\"assets/image-20210502094911-rr579ht.png\" data-src=\"assets/image-20210502094911-rr579ht.png\" alt=\"image.png\" updated=\"20210502094911\" style=\"width: 867px;\"><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\"></span></span></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210502094636-x7veqk8\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210502094911\"><div contenteditable=\"true\" spellcheck=\"false\"><span contenteditable=\"false\" data-type=\"img\" class=\"img\"><span class=\"protyle-action\"><svg class=\"svg\"><use xlink:href=\"#iconMore\"></use></svg></span><img src=\"/siyuan/0/测试笔记/assets/image-20210502094911-rr579ht.png\" data-src=\"assets/image-20210502094911-rr579ht.png\" alt=\"image.png\" updated=\"20210502094911\" style=\"width: 867px;\" /><span class=\"protyle-action__drag\"></span><span class=\"protyle-action__title\"></span></span>\n</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"12", "<div data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\"><div data-marker=\"*\" data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li protyle-task--done\"><div class=\"protyle-action protyle-action--task\"><svg><use xlink:href=\"#iconCheck\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\"><div data-marker=\"*\" data-subtype=\"t\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li protyle-task--done\"><div class=\"protyle-action protyle-action--task\"><svg><use xlink:href=\"#iconCheck\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"11", "<div data-subtype=\"o\" data-node-id=\"20210430192929-txskdcn\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\" updated=\"20210430192930\"><div data-marker=\"1.\" data-subtype=\"o\" data-node-id=\"20210430192930-83abtjf\" data-type=\"NodeListItem\" class=\"li\"><div class=\"protyle-action protyle-action--order\">1.</div><div data-node-id=\"20210430192930-4rx7cjp\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-subtype=\"o\" data-node-id=\"20210430192929-txskdcn\" data-node-index=\"1\" data-type=\"NodeList\" class=\"list\" updated=\"20210430192930\"><div data-marker=\"1.\" data-subtype=\"o\" data-node-id=\"20210430192930-83abtjf\" data-type=\"NodeListItem\" class=\"li\"><div class=\"protyle-action protyle-action--order\" contenteditable=\"false\">1.</div><div data-node-id=\"20210430192930-4rx7cjp\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"10", "<div data-subtype=\"h2\" data-node-id=\"20210430171307-joyl2la\" data-node-index=\"1\" data-type=\"NodeHeading\" class=\"h2\" updated=\"20210430171401\"><div contenteditable=\"true\" spellcheck=\"false\"></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-subtype=\"h2\" data-node-id=\"20210430171307-joyl2la\" data-node-index=\"1\" data-type=\"NodeHeading\" class=\"h2\" updated=\"20210430171401\"><div contenteditable=\"true\" spellcheck=\"false\"></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	//{"9", "<div data-node-id=\"20210430093711-eyx7gyw\" data-node-index=\"1\" data-type=\"NodeCodeBlock\" class=\"code-block\" updated=\"20210430094009\"><div class=\"protyle-action\"><div class=\"protyle-action__language\">js</div><div class=\"protyle-action__copy\"><span aria-label=\"复制\" onmouseover=\"this.setAttribute('aria-label', '复制')\" class=\"b3-tooltips b3-tooltips__n\"><svg class=\"svg\"><use xlink:href=\"#iconCopy\"></use></svg></span></div></div><div contenteditable=\"true\" spellcheck=\"false\" class=\"hljs nginx protyle-linenumber\" data-render=\"true\"><span class=\"hljs-attribute\">foo</span>\n\n```<wbr></div><span contenteditable=\"false\" class=\"protyle-linenumber__rows\"><span style=\"height:20px\"></span><span style=\"height:20px\"></span><span style=\"height:20px\"></span></span><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210430093711-eyx7gyw\" data-node-index=\"1\" data-type=\"NodeCodeBlock\" class=\"code-block\" updated=\"20210430094009\"><div class=\"protyle-action\"><div class=\"protyle-action__language\">js</div><div class=\"protyle-action__copy\"></div></div><div contenteditable=\"true\" spellcheck=\"false\">foo\n\n\u200d```<wbr>\n</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	//{"8", "<div data-node-id=\"20210429212952-th2sdjq\" data-node-index=\"1\" data-type=\"NodeCodeBlock\" class=\"code-block\" updated=\"20210429213602\"><div class=\"protyle-action\"><div class=\"protyle-action__language\">js</div><div class=\"protyle-action__copy\"><span aria-label=\"复制\" onmouseover=\"this.setAttribute('aria-label', '复制')\" class=\"b3-tooltips b3-tooltips__n\"><svg class=\"svg\"><use xlink:href=\"#iconCopy\"></use></svg></span></div></div><div contenteditable=\"true\" spellcheck=\"false\" class=\"hljs js protyle-linenumber\" data-render=\"true\"><span class=\"hljs-attribute\">foo<wbr></span>\n</div><span contenteditable=\"false\" class=\"protyle-linenumber__rows\"><span style=\"height:16px\"></span></span><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210429212952-th2sdjq\" data-node-index=\"1\" data-type=\"NodeCodeBlock\" class=\"code-block\" updated=\"20210429213602\"><div class=\"protyle-action\"><div class=\"protyle-action__language\">js</div><div class=\"protyle-action__copy\"></div></div><div contenteditable=\"true\" spellcheck=\"false\">foo<wbr>\n</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"7", "<div data-node-id=\"20210428163425-gd63njj\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210428163507\"><div contenteditable=\"true\" spellcheck=\"false\">{{foo}}<wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-content=\"foo\" data-node-id=\"20210428163425-gd63njj\" data-node-index=\"1\" data-type=\"NodeBlockQueryEmbed\" class=\"render-node\" updated=\"20210428163507\"></div>"},
	{"6", "<div data-node-id=\"20210428155259-1j2zqx0\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210428155312\"><div contenteditable=\"true\" spellcheck=\"false\">foo\n\nb<wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210428155259-1j2zqx0\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210428155312\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\">b<wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"5", "<div data-node-id=\"20210428094047-w9di4p3\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210428094048\"><div contenteditable=\"true\" spellcheck=\"false\"># <wbr></div><div class=\"protyle-attr\"></div></div>", "<div data-subtype=\"h1\" data-node-id=\"20210428094047-w9di4p3\" data-node-index=\"1\" data-type=\"NodeHeading\" class=\"h1\" updated=\"20210428094048\"><div contenteditable=\"true\" spellcheck=\"false\"><wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"4", "<div data-node-id=\"20210428094047-w9di4p3\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210428094048\"><div contenteditable=\"true\" spellcheck=\"false\">#<wbr></div><div class=\"protyle-attr\"></div></div>", "<div data-node-id=\"20210428094047-w9di4p3\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210428094048\"><div contenteditable=\"true\" spellcheck=\"false\">#<wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"3", "<div data-content=\"name:foo\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeBlockQueryEmbed\" class=\"render-node\"></div>", "<div data-content=\"name:foo\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeBlockQueryEmbed\" class=\"render-node\"></div>"},
	{"2", "<div data-node-id=\"20210501114345-dc8wcm0\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210501114428\"><div contenteditable=\"true\" spellcheck=\"false\">foo\n* <wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210501114345-dc8wcm0\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210501114428\"><div contenteditable=\"true\" spellcheck=\"false\">foo</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"2\" data-type=\"NodeList\" class=\"list\"><div data-marker=\"*\" data-subtype=\"u\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeListItem\" class=\"li\"><div class=\"protyle-action\"><svg><use xlink:href=\"#iconDot\"></use></svg></div><div data-node-id=\"20060102150405-1a2b3c4\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><wbr></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"1", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><kbd>foo</kbd></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20060102150405-1a2b3c4\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\"><div contenteditable=\"true\" spellcheck=\"false\"><kbd>foo</kbd></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
	{"0", "<div data-node-id=\"20210426094859-uataalw\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210426101601\"><div contenteditable=\"true\" spellcheck=\"false\">[[<wbr><wbr><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span>\n</div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>", "<div data-node-id=\"20210426094859-uataalw\" data-node-index=\"1\" data-type=\"NodeParagraph\" class=\"p\" updated=\"20210426101601\"><div contenteditable=\"true\" spellcheck=\"false\">[[<wbr><wbr><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span><span data-type=\"block-ref\" data-id=\"20210426091959-npvs57l\" data-anchor=\"\" contenteditable=\"false\"></span></div><div class=\"protyle-attr\" contenteditable=\"false\"></div></div>"},
}

func TestSpinBlockDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.SetProtyleWYSIWYG(true)
	luteEngine.ParseOptions.Mark = true
	luteEngine.ParseOptions.BlockRef = true
	luteEngine.SetKramdownIAL(true)
	luteEngine.ParseOptions.SuperBlock = true
	luteEngine.SetLinkBase("/siyuan/0/测试笔记/")
	luteEngine.SetAutoSpace(false)
	luteEngine.SetSub(true)
	luteEngine.SetSup(true)
	luteEngine.SetGitConflict(true)
	luteEngine.SetIndentCodeBlock(false)
	luteEngine.SetEmojiSite("http://127.0.0.1:6806/stage/protyle/images/emoji")

	ast.Testing = true
	for _, test := range spinBlockDOMTests {
		html := luteEngine.SpinBlockDOM(test.from)

		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
	ast.Testing = false
}
