package com.hapticapps.amici.shared.data_models.org;

import com.hapticapps.amici.shared.utils.Utils;
import org.bson.codecs.pojo.annotations.BsonProperty;

public class OrgPerson {
	public static final String BSON_UUID = "u";
	public static final String BSON_ORGID = "o";
	public static final String BSON_NAME = "n";
	public static final String BSON_EMAIL = "e";

	@BsonProperty("u")
	private String uuid;

	@BsonProperty("o")
	private String orgId;

	@BsonProperty("n")
	private String name;

	@BsonProperty("e")
	private String email;


	public String getUuid(){
		return uuid;
	}
	public String getOrgId(){
		return orgId;
	}
	public String getName(){
		return name;
	}
	public String getEmail(){
		return email;
	}

	public OrgPersonBuilder copy() {
		return OrgPersonBuilder.from( this );
	}
	public String toString() {
		return "OrgPerson{" + 
			"}";
	}

	public static class OrgPersonBuilder {

		private String uuid = Utils.newUID();

		private String orgId;

		private String name;

		private String email;

		public static OrgPersonBuilder from(OrgPerson source) {
			var r = new OrgPersonBuilder();
			r.uuid = source.getUuid();
			r.orgId = source.getOrgId();
			r.name = source.getName();
			r.email = source.getEmail();
			return r;
		}

		public OrgPersonBuilder setUuid(String uuid) {
			this.uuid = uuid;
			return this;
		}

		public OrgPersonBuilder setOrgId(String orgId) {
			this.orgId = orgId;
			return this;
		}

		public OrgPersonBuilder setName(String name) {
			this.name = name;
			return this;
		}

		public OrgPersonBuilder setEmail(String email) {
			this.email = email;
			return this;
		}

		public OrgPerson build() {
			var r = new OrgPerson();
			r.uuid = uuid;
			r.orgId = orgId;
			r.name = name;
			r.email = email;
			return r;
		}
	}
}

