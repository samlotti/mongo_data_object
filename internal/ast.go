package internal

import "unicode"

type IAst interface {
}

type Ast struct {
}

// AstFile - Represents a single input file
// but can have multiple classes.
type AstFile struct {
	pkg     string
	imports []string
	entity  *AstEntity
	classes []*AstClass
	enums   []*AstEnum
}

type AstEntity struct {
	name    string
	data    []*AstData
	indexes []*AstIndex
	show    []string
}

type AstClass struct {
	name string
	data []*AstData
	show []string
}

type AstEnum struct {
	name string
	data []string
}

type AstIndex struct {
	keys       []*AstIndexKeys
	unique     int
	sparse     int
	background int
}

type AstIndexKeys struct {
	dname   string
	ascDesc int
}

type AstData struct {
	dtype   string
	dname   string
	dnameAs string
	dflt    string
}

func (d *AstData) getAsName() string {
	if d.hasNameAs() {
		return d.dnameAs
	} else {
		return d.dname
	}
}

func (d *AstData) hasNameAs() bool {
	return len(d.dnameAs) > 0
}

func (d *AstData) hasDefault() bool {
	return len(d.dflt) > 0
}

func (d *AstData) getterName() string {
	rn := []rune(d.dname)
	rn[0] = unicode.ToUpper(rn[0])
	s := string(rn)

	return "get" + s
}

func (d *AstData) setterName() string {
	rn := []rune(d.dname)
	rn[0] = unicode.ToUpper(rn[0])
	s := string(rn)

	return "set" + s

}

func (d *AstData) withName() string {
	rn := []rune(d.dname)
	rn[0] = unicode.ToUpper(rn[0])
	s := string(rn)

	return "with" + s

}
