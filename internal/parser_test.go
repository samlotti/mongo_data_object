package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser1(t *testing.T) {
	sample := `
package com.hapticapps.amici.shared.data_models.org;

import com.hapticapps.amici.shared.utils.Utils;
import org.bson.codecs.pojo.annotations.BsonProperty;

entity OrgPerson {

    data ~String~ uuid as u = ~new Id()~;

    data String orgId as o;

    data String name as n;

    data String email as e;
	
	data Address address as a1;

	data OrgState state as st;

}


class Address {

    data String street as s;

    data String city as c;

    data String st as st;

    data String zip as z;

}

enum OrgState {
        PENDING,
        ACTIVE,
        EXPIRED,
        UNKNOWN
}


`

	// package com.hapticapps.amici.shared.data_models.org;
	p := NewParser(NewLexer(sample, "TestLexer1"))
	ast := p.parse()
	assert.Equal(t, "com.hapticapps.amici.shared.data_models.org", ast.pkg)
	assert.Equal(t, 2, len(ast.imports))
	assert.Equal(t, "org.bson.codecs.pojo.annotations.BsonProperty", ast.imports[1])
	assert.Equal(t, 1, len(ast.classes))
	assert.Equal(t, "Address", ast.classes[0].name)
	assert.Equal(t, "OrgPerson", ast.entity.name)
	assert.Equal(t, "OrgState", ast.enums[0].name)
	assert.Equal(t, 4, len(ast.enums[0].data))

}
