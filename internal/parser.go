package internal

import "fmt"

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
		tk := p.expectNext(CLASS, ENTITY, IMPORT, EOF)
		if tk.Type == IMPORT {
			ast.imports = append(ast.imports, p.parse_import())
		}

		if tk.Type == CLASS {
			ast.classes = append(ast.classes, p.parse_class())
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
	panic(fmt.Sprintf("invalid token: %s", tk.Literal))
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
	d.dtype = p.expectNext(IDENTIFIER).Literal
	d.dname = p.expectNext(IDENTIFIER).Literal
	tk := p.expectNext(AS, SEMI, EQUAL)
	if tk.Type == AS {
		d.dnameAs = p.expectNext(IDENTIFIER).Literal
		tk = p.expectNext(SEMI, EQUAL)
	}
	if tk.Type == EQUAL {
		// d.dflt = p.expectNext()
		tk = p.expectNext(SEMI)
	}
	return d
}

func (p *Parser) parse_entity() *AstEntity {
	c := &AstEntity{}
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