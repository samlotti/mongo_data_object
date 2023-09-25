package internal

import (
	"strings"
)

type Render struct {
	ast     *AstFile
	builder *strings.Builder
	depth   int
}

func NewRender(ast *AstFile) *Render {
	return &Render{
		ast:     ast,
		builder: nil,
		depth:   0,
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
	r.write("import java.util.Map;\n")

	r.write("import java.util.List;\n")
	r.write("import java.util.ArrayList;\n")

	r.write("import com.mongodb.client.model.IndexOptions;\n")
	r.write("import org.bson.conversions.Bson;\n")
	r.write("import com.mongodb.BasicDBObject;\n")
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

	r.renderCommonEntityClass(entity.name, entity.data, false)

	for _, cls := range r.ast.classes {
		r.renderClass(cls)
	}

	for _, enum := range r.ast.enums {
		r.renderEnum(enum)
	}

	r.renderIndexDefs(entity.name, entity.indexes)

	r.end().nl()

}

func (r *Render) renderClass(cls *AstClass) {
	r.renderCommonEntityClass(cls.name, cls.data, true)
	r.tabe().end().nl()
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
	r.tabs().write("public ").write(dta.dtype).space()

	r.write(dta.getterName()).write("()").begin().nl()

	r.tabs().write("return ").write(dta.dname).semi().nl()

	r.tabe().end().nl()
}

func (r *Render) writeSetter(dta *AstData) {
	r.tabs().write("public void ")

	r.write(dta.setterName()).write("(").w(dta.dtype).s().w("data)").begin().nl()

	r.tabs().write("this.").write(dta.dname).w(" = data").semi().nl()

	r.tabe().end().nl()
}

func (r *Render) writeBuilderCopyFunction(name string) {
	r.tabs().write("public ").write(r.builderName(name)).space().write("copy() ").begin().nl()
	r.tabs().write("return ").write(r.builderName(name)).write(".from( this );").nl()
	r.tabe().end().nl().nl()

	r.tabs().write("public static ").write(r.builderName(name)).space().write("builder() ").begin().nl()
	r.tabs().write("return new ").write(r.builderName(name)).write("();").nl()
	r.tabe().end().nl()

}

func (r *Render) builderName(name string) string {
	return name + "Builder"
}

func (r *Render) writeToString(name string, data []*AstData) {
	r.tabs().write("public String toString() ").begin().nl()
	r.tabs().write("return \"").write(name).write("{\" + ").nl()
	r.tabs().writeQ("}").semi().nl()
	r.tabe().end().nl().nl()

}

func (r *Render) writeQ(s string) *Render {
	return r.write("\"").write(s).write("\"")
}

func (r *Render) writeBuilderInnerClass(name string, data []*AstData) {
	r.tabs().w("public static class ").write(r.builderName(name)).s().begin().nl().nl()
	for _, d := range data {
		r.tabs().w("private ").w(d.dtype).s().w(d.dname)
		if d.hasDefault() {
			r.w(" = ").w(d.dflt)
		}
		r.semi().nl().nl()
	}

	r.writeFromFunction(name, data)

	r.writeBuilderMethods(name, data)

	r.nl()
	r.writeBuildMethods(name, data)

	r.tabe().end().nl()

}

func (r *Render) w(s string) *Render {
	return r.write(s)
}

func (r *Render) s() *Render {
	return r.space()
}

func (r *Render) end() *Render {
	r.depth--
	return r.w("}")
}

func (r *Render) begin() *Render {
	r.depth++
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
	r.tabs().w("public static ").w(r.builderName(name)).s().w("from(").w(name).w(" source) ").begin().nl()
	r.tabs().w("var r = new ").w(r.builderName(name)).w("();").nl()
	for _, d := range data {
		r.tabs().w("r.").w(d.dname).w(" = source.").w(d.getterName()).w("()").semi().nl()
	}
	r.tabs().w("return r;").nl()
	r.tabe().end().nl()
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
		r.tabs().w("public ").w(r.builderName(name)).s().w(d.setterName()).
			w("(").w(d.dtype).s().w(d.dname).w(") ").begin().nl()
		r.tabs().w("this.").w(d.dname).w(" = ").w(d.dname).semi().nl()
		r.tabs().w("return this;").nl()
		r.tabe().end().nl()
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
	r.tabs().w("public ").w(name).s().w("build() ").begin().nl()
	r.tabs().w("var r = new ").w(name).w("()").semi().nl()
	for _, d := range data {
		r.tabs().w("r.").w(d.dname).w(" = ").w(d.dname).semi().nl()
	}
	r.tabs().w("return r;").nl()
	r.tabe().end().nl()
}

// renderEnum
//
// public enum OrgState {
// PENDING,
// ACTIVE,
// EXPIRED,
//
// UNKNOWN
// }
func (r *Render) renderEnum(enum *AstEnum) {
	r.nl()
	r.tabs().w("public enum ").w(enum.name).s().begin().nl()
	for idx, n := range enum.data {
		if idx > 0 {
			r.w(",").nl()
		}
		r.tabs().w(n)
	}
	r.nl()
	r.tabe().end().nl()
}

func (r *Render) renderCommonEntityClass(name string, data []*AstData, inner bool) {

	pc := "public class "
	if inner {
		pc = "public static class "
	}

	r.tabs().write(pc).write(name).s().begin().nl()
	for _, dta := range data {
		r.tabs().write("public static final String").
			space().write("BSON_").write(strings.ToUpper(dta.dname)).
			space().write("=").space().qstr(dta.getAsName()).semi().nl()
	}
	r.nl()
	for _, dta := range data {
		if dta.hasNameAs() {
			r.tabs().write("@BsonProperty(").qstr(dta.getAsName()).write(")").nl()
		}
		r.tabs().write("private ").write(dta.dtype).space().write(dta.dname)
		r.semi().nl().nl()
	}

	r.nl()
	for _, dta := range data {
		r.writeGetter(dta)
		r.writeSetter(dta)
	}

	r.nl()

	r.writeBuilderCopyFunction(name)
	r.writeToString(name, data)

	r.writeBuilderInnerClass(name, data)

}

func (r *Render) tabs() *Render {
	for i := 0; i < r.depth; i++ {
		r.tab()
	}
	return r
}

// tabe - tab for end the .end() is next.
func (r *Render) tabe() *Render {
	for i := 0; i < r.depth-1; i++ {
		r.tab()
	}
	return r
}

func (r *Render) renderIndexDefs(name string, indexes []*AstIndex) {

	r.tabs().w("public static class Indexes ").begin().nl()

	r.tabs().w("public static List<Bson> ikeys = new ArrayList<>();").nl()
	r.tabs().w("public static List<IndexOptions> ioptions = new ArrayList<>();").nl()
	r.tabs().w("static ").begin().nl()
	for _, idef := range indexes {
		r.tabs().w("ikeys.add(new BasicDBObject( Map.of(").nl()
		r.depth++
		for idx, kk := range idef.keys {
			if idx > 0 {
				r.w(",").nl()
			}
			r.tabs().w(name).w(".").w("BSON_").write(strings.ToUpper(kk.dname)).w(", ")
			if kk.ascDesc == 1 {
				r.w("1")
			} else {
				r.w("-1")
			}
		}
		r.nl()
		r.depth--

		r.tabs().w(")));").nl()

		r.tabs().w("ioptions.add(new IndexOptions()")
		if idef.unique == 1 {
			r.w(".unique(true)")
		}
		if idef.sparse == 1 {
			r.w(".sparse(true)")
		}
		if idef.background == 1 {
			r.w(".background(true)")
		}

		r.w(");").nl()

	}

	r.tabe().end().nl()

	r.tabe().end().nl()

}
