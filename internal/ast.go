package internal

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
}

type AstEntity struct {
	name string
	data []*AstData
}

type AstClass struct {
	name string
	data []*AstData
}

type AstData struct {
	dtype   string
	dname   string
	dnameAs string
	dflt    string
}
