package internal

import (
	"fmt"
	"strings"
)

type Parser struct {
	lex *Lexer
}

func NewParser(lex *Lexer) *Parser {
	p := &Parser{
		lex: lex,
	}
	return p
}

func (p *Parser) parse() *AstFile {
	ast := &AstFile{}

	p.expectNext(PACKAGE)
	ast.pkg = p.parse_package()

	for {
		tk := p.expectNext(CLASS, ENTITY, IMPORT, EOF, ENUM)
		if tk.Type == IMPORT {
			ast.imports = append(ast.imports, p.parse_import())
		}

		if tk.Type == CLASS {
			ast.classes = append(ast.classes, p.parse_class())
		}

		if tk.Type == ENUM {
			ast.enums = append(ast.enums, p.parse_enum())
		}

		if tk.Type == ENTITY {
			if ast.entity != nil {
				panic("multiple entities not supported!")
			}
			ast.entity = p.parse_entity()
		}

		if tk.Type == EOF {
			break
		}
	}

	return ast
}

// package com.hapticapps.amici.shared.data_models.org;
func (p *Parser) parse_package() string {
	tk := p.expectNext(IDENTIFIER)
	p.expectNext(SEMI)
	return tk.Literal
}

func (p *Parser) expectNext(ids ...TokenType) *Token {
	tk := p.lex.NextToken()
	for _, id := range ids {
		if id == tk.Type {
			return tk
		}
	}

	sb := &strings.Builder{}
	for _, id := range ids {
		sb.WriteString(string(id))
		sb.WriteString(", ")
	}

	panic(fmt.Sprintf("invalid token: '%s' expected one of: [%s] at: %d:%d", tk.Literal, sb.String(), p.lex.lineNum, p.lex.lPos))
	return nil
}

// import com.hapticapps.amici.shared.data_models.org;
func (p *Parser) parse_import() string {
	tk := p.expectNext(IDENTIFIER)
	p.expectNext(SEMI)
	return tk.Literal
}

func (p *Parser) parse_class() *AstClass {
	c := &AstClass{}
	c.name = p.expectNext(IDENTIFIER).Literal
	p.expectNext(LBRACE)

	for {
		tk := p.expectNext(RBRACE, DATA)
		if tk.Type == RBRACE {
			break
		}

		c.data = append(c.data, p.parse_data())

	}
	return c
}

// parse_data --
// ex:  String uuid as u = Utils.newUID();
func (p *Parser) parse_data() *AstData {
	d := &AstData{}
	d.dtype = p.expectNext(IDENTIFIER, CBLOCK).Literal
	d.dname = p.expectNext(IDENTIFIER).Literal
	tk := p.expectNext(AS, SEMI, EQUAL)
	if tk.Type == AS {
		d.dnameAs = p.expectNext(IDENTIFIER).Literal
		tk = p.expectNext(SEMI, EQUAL)
	}
	if tk.Type == EQUAL {
		d.dflt = p.expectNext(IDENTIFIER, CBLOCK).Literal
		tk = p.expectNext(SEMI)
	}
	return d
}

func (p *Parser) parse_entity() *AstEntity {
	c := &AstEntity{}
	c.name = p.expectNext(IDENTIFIER).Literal
	p.expectNext(LBRACE)

	for {
		tk := p.expectNext(RBRACE, DATA, INDEX, SHOW)
		if tk.Type == RBRACE {
			break
		}

		if tk.Type == INDEX {
			p.parse_index(c)
			continue
		}

		if tk.Type == SHOW {
			p.expectNext(LPAREN)
			for {
				tk := p.expectNext(IDENTIFIER)
				c.show = append(c.show, tk.Literal)
				tk = p.expectNext(COMMA, RPAREN)
				if tk.Type == RPAREN {
					break
				}
			}
			p.expectNext(SEMI)
			continue
		}

		c.data = append(c.data, p.parse_data())

	}
	return c
}

// parse_enum
//
//enum OrgState {
//PENDING,
//ACTIVE,
//EXPIRED,
//UNKNOWN
//}

func (p *Parser) parse_enum() *AstEnum {
	r := &AstEnum{}
	r.name = p.expectNext(IDENTIFIER).Literal
	p.expectNext(LBRACE)
	for {
		n := p.expectNext(IDENTIFIER)
		r.data = append(r.data, n.Literal)

		n = p.expectNext(COMMA, RBRACE)
		if n.Type == RBRACE {
			break
		}
	}
	return r
}

func (p *Parser) parse_index(c *AstEntity) {
	p.expectNext(LPAREN)

	idx := &AstIndex{
		keys:       nil,
		unique:     0,
		sparse:     0,
		background: 0,
	}
	c.indexes = append(c.indexes, idx)

	for {
		tk := p.expectNext(IDENTIFIER)

		k := &AstIndexKeys{
			dname:   tk.Literal,
			ascDesc: 1,
		}
		idx.keys = append(idx.keys, k)

		peek := p.peekToken()
		if peek.Type == ASC || peek.Type == DESC {
			peek = p.expectNext(ASC, DESC)
			if peek.Type == ASC {
				k.ascDesc = 1
			}
			if peek.Type == DESC {
				k.ascDesc = -1
			}
		}

		nxt := p.expectNext(COMMA, RPAREN)
		if nxt.Type == RPAREN {
			break
		}
	}

	for {
		tk := p.expectNext(SPARSE, BACKGROUND, UNIQUE, SEMI)
		if tk.Type == SEMI {
			break
		}

		if tk.Type == UNIQUE {
			idx.unique = 1
		}

		if tk.Type == BACKGROUND {
			idx.background = 1
		}

		if tk.Type == SPARSE {
			idx.sparse = 1
		}
	}

}

func (p *Parser) peekToken() *Token {
	return p.lex.PeekToken()
}
