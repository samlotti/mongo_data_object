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

    data String uuid as u = ~Utils.newUID()~;

    data String orgId as o;

    data String name as n;

    data String email as e;

}



`

	// package com.hapticapps.amici.shared.data_models.org;
	p := NewParser(NewLexer(sample, "TestLexer1"))
	r := NewRender(p.parse())
	str := r.Render()
	assert.True(t, strings.Contains(str,
		`package com.hapticapps.amici.shared.data_models.org;

import com.hapticapps.amici.shared.utils.Utils;
import org.bson.codecs.pojo.annotations.BsonProperty;

`))
	assert.True(t, strings.Contains(str, `public static final String BSON_UUID = "u";`))
	assert.True(t, strings.Contains(str, `public OrgPersonBuilder setUuid(String uuid) {`))

	fmt.Println(str)

}