package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRenderOrgPerson(t *testing.T) {
	sample := `
package com.hapticapps.amici.shared.data_models.org;

import com.hapticapps.amici.shared.utils.Utils;

/**
    The org person
**/
entity OrgPerson {
    index (uuid asc) unique;
    index (orgId asc) unique;
    index (email asc) sparse;
    index (state asc, name desc) sparse;


    data String uuid as u = ~Utils.newUID()~;

    data String orgId as o;

    data String name as n;

    data String email as e;

	data OrgState state as os;

	show (name, email, state);

}

enum OrgState {
        PENDING,
        ACTIVE,
        EXPIRED,
        UNKNOWN
}

/**
Products owned by the organization with expiration date
**/
class Product {
    data String productId as pid;
    data long expData as ed;
    data String license as lic;
}





`

	// package com.hapticapps.amici.shared.data_models.org;
	p := NewParser(NewLexer(sample, "TestLexer1"))
	r := NewRender(p.parse())
	str := r.Render()
	assert.True(t, strings.Contains(str,
		`// Generated by mdo do not edit this file, see the .mdo file`))

	fmt.Println(str)

	assert.True(t, strings.Contains(str, `public static final String BSON_UUID = "u";`))
	assert.True(t, strings.Contains(str, `public OrgPersonBuilder withUuid(String uuid) {`))
	assert.True(t, strings.Contains(str, `public enum OrgState {`))

}
