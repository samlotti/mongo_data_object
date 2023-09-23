package internal

import (
	"strings"
)

type Render struct {
	ast     *AstFile
	builder *strings.Builder
}

func NewRender(ast *AstFile) *Render {
	return &Render{
		ast:     ast,
		builder: nil,
	}
}

func (r *Render) Render() string {
	r.builder = &strings.Builder{}
	r.write("package ").write(r.ast.pkg).semi().nl()
	r.nl()

	for _, imprt := range r.ast.imports {
		r.write("import ").write(imprt).semi().nl()
	}
	r.write("import org.bson.codecs.pojo.annotations.BsonProperty;\n")
	r.nl()

	r.renderEntity(r.ast.entity)

	r.nl()
	return r.builder.String()
}

func (r *Render) write(s string) *Render {
	r.builder.WriteString(s)
	return r
}

func (r *Render) nl() *Render {
	r.builder.WriteString("\n")
	return r
}

func (r *Render) renderEntity(entity *AstEntity) {
	r.write("public class ").write(entity.name).s().begin().nl()
	for _, dta := range entity.data {
		r.tab().write("public static final String").
			space().write("BSON_").write(strings.ToUpper(dta.dname)).
			space().write("=").space().qstr(dta.getAsName()).semi().nl()
	}
	r.nl()
	for _, dta := range entity.data {
		if dta.hasNameAs() {

			r.tab().write("@BsonProperty(").qstr(dta.getAsName()).write(")").nl()
		}
		r.tab().write("private ").write(dta.dtype).space().write(dta.dname)
		//if dta.hasDefault() {
		//	r.write(" = ").write(dta.dflt)
		//}
		r.semi().nl().nl()
	}

	r.nl()
	for _, dta := range entity.data {
		r.writeGetter(dta)
	}

	r.nl()

	r.writeBuilderCopyFunction(entity.name)
	r.writeToString(entity.name, entity.data)

	r.writeBuilderInnerClass(entity.name, entity.data)

	r.end().nl()

}

func (r *Render) tab() *Render {
	r.write("\t")
	return r

}

func (r *Render) space() *Render {
	r.write(" ")
	return r
}

func (r *Render) qstr(name string) *Render {
	r.write("\"").write(name).write("\"")
	return r
}

func (r *Render) semi() *Render {
	return r.write(";")
}

func (r *Render) writeGetter(dta *AstData) {
	r.tab().write("public ").write(dta.dtype).space()

	r.write(dta.getterName()).write("()").begin().nl()

	r.tab2().write("return ").write(dta.dname).semi().nl()

	r.tab().end().nl()
}

func (r *Render) writeBuilderCopyFunction(name string) {
	r.tab().write("public ").write(r.builderName(name)).space().write("copy() ").begin().nl()
	r.tab2().write("return ").write(r.builderName(name)).write(".from( this );").nl()
	r.tab().end().nl()
}

func (r *Render) builderName(name string) string {
	return name + "Builder"
}

func (r *Render) writeToString(name string, data []*AstData) {
	r.tab().write("public String toString() ").begin().nl()
	r.tab2().write("return \"").write(name).write("{\" + ").nl()
	r.tab3().writeQ("}").semi().nl()
	r.tab().end().nl().nl()

}

func (r *Render) writeQ(s string) *Render {
	return r.write("\"").write(s).write("\"")
}

func (r *Render) writeBuilderInnerClass(name string, data []*AstData) {
	r.tab().w("public static class ").write(r.builderName(name)).s().begin().nl().nl()
	for _, d := range data {
		r.tab2().w("private ").w(d.dtype).s().w(d.dname)
		if d.hasDefault() {
			r.w(" = ").w(d.dflt)
		}
		r.semi().nl().nl()
	}

	r.writeFromFunction(name, data)

	r.writeBuilderMethods(name, data)

	r.nl()
	r.writeBuildMethods(name, data)

	r.tab().end().nl()

}

func (r *Render) tab2() *Render {
	return r.tab().tab()
}

func (r *Render) tab3() *Render {
	return r.tab2().tab()
}

func (r *Render) tab4() *Render {
	return r.tab2().tab2()
}

func (r *Render) w(s string) *Render {
	return r.write(s)
}

func (r *Render) s() *Render {
	return r.space()
}

func (r *Render) end() *Render {
	return r.w("}")
}

func (r *Render) begin() *Render {
	return r.w("{")
}

// writeFromFunction
//
//	 public static OrgPersonBuilder from(OrgPerson source) {
//	    var r = new OrgPersonBuilder();
//	    r.uuid = source.getUuid();
//	    r.name = source.getName();
//	    r.orgId = source.getOrgId();
//	    r.email = source.getEmail();
//	    return r;
//	}
func (r *Render) writeFromFunction(name string, data []*AstData) {
	r.tab2().w("public static ").w(r.builderName(name)).s().w("from(").w(name).w(" source) ").begin().nl()
	r.tab3().w("var r = new ").w(r.builderName(name)).w("();").nl()
	for _, d := range data {
		r.tab3().w("r.").w(d.dname).w(" = source.").w(d.getterName()).w("()").semi().nl()
	}
	r.tab3().w("return r;").nl()
	r.tab2().end().nl()
}

// writeBuilderMethods
//
//	 public OrgPersonBuilder setEmail(String email) {
//	    this.email = email;
//	    return this;
//	}
func (r *Render) writeBuilderMethods(name string, data []*AstData) {
	for _, d := range data {
		r.nl()
		r.tab2().w("public ").w(r.builderName(name)).s().w(d.setterName()).
			w("(").w(d.dtype).s().w(d.dname).w(") ").begin().nl()
		r.tab3().w("this.").w(d.dname).w(" = ").w(d.dname).semi().nl()
		r.tab3().w("return this;").nl()
		r.tab2().end().nl()
	}
}

// writeBuildMethods -
//
//	 public OrgPerson build() {
//	    var r = new OrgPerson();
//	    r.uuid = uuid;
//	    r.orgId = orgId;
//	    r.email = email;
//	    r.name = name;
//	    return r;
//	}
func (r *Render) writeBuildMethods(name string, data []*AstData) {
	r.tab2().w("public ").w(name).s().w("build() ").begin().nl()
	r.tab3().w("var r = new ").w(name).w("()").semi().nl()
	for _, d := range data {
		r.tab3().w("r.").w(d.dname).w(" = ").w(d.dname).semi().nl()
	}
	r.tab3().w("return r;").nl()
	r.tab2().end().nl()
}
