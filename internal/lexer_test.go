package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLexer1(t *testing.T) {
	sample := `
package com.hapticapps.amici.shared.data_models.org;

import com.hapticapps.amici.shared.utils.Utils;
import ~org.bson.codecs.pojo.annotations.BsonProperty~;

/**
A Comment
**/

class OrgPerson {

    data String uuid as u = Utils.newUID();

    data String orgId as o;

    data String name as n;

    data String email as e;

}

`

	// package com.hapticapps.amici.shared.data_models.org;
	lex := NewLexer(sample, "TestLexer1")
	tkn := lex.NextToken()
	assert.Equal(t, PACKAGE, string(tkn.Type))
	assert.Equal(t, "package", tkn.Literal)
	tkn = lex.NextToken()
	assert.Equal(t, IDENTIFIER, string(tkn.Type))
	assert.Equal(t, "com.hapticapps.amici.shared.data_models.org", tkn.Literal)
	tkn = lex.NextToken()
	assert.Equal(t, SEMI, string(tkn.Type))

	// import com.hapticapps.amici.shared.utils.Utils;
	tkn = lex.NextToken()
	assert.Equal(t, IMPORT, string(tkn.Type))
	tkn = lex.NextToken()
	assert.Equal(t, IDENTIFIER, string(tkn.Type))
	tkn = lex.NextToken()
	assert.Equal(t, SEMI, string(tkn.Type))

	// import org.bson.codecs.pojo.annotations.BsonProperty;
	tkn = lex.NextToken()
	assert.Equal(t, IMPORT, string(tkn.Type))
	tkn = lex.NextToken()
	assert.Equal(t, CBLOCK, string(tkn.Type))
	tkn = lex.NextToken()
	assert.Equal(t, SEMI, string(tkn.Type))

	// class OrgPerson {
	tkn = lex.NextToken()
	assert.Equal(t, CLASS, string(tkn.Type))

}
