package com.hapticapps.amici.shared.data_models.org;

import com.hapticapps.amici.shared.utils.Utils;
import org.bson.codecs.pojo.annotations.BsonProperty;

public class Org {
	public static final String BSON_UUID = "u";
	public static final String BSON_NAME = "n";
	public static final String BSON_CITY = "c";
	public static final String BSON_STATE = "state";
	public static final String BSON_ZIP = "z";
	public static final String BSON_MEMBERS = "m";
	public static final String BSON_EMAIL = "e";

	@BsonProperty("u")
	private String uuid;

	@BsonProperty("n")
	private String name;

	@BsonProperty("c")
	private String city;

	private String state;

	@BsonProperty("z")
	private String zip;

	@BsonProperty("m")
	private List<MemberId> members;

	@BsonProperty("e")
	private String email;


	public String getUuid(){
		return uuid;
	}
	public String getName(){
		return name;
	}
	public String getCity(){
		return city;
	}
	public String getState(){
		return state;
	}
	public String getZip(){
		return zip;
	}
	public List<MemberId> getMembers(){
		return members;
	}
	public String getEmail(){
		return email;
	}

	public OrgBuilder copy() {
		return OrgBuilder.from( this );
	}
	public String toString() {
		return "Org{" + 
			"}";
	}

	public static class OrgBuilder {

		private String uuid = Utils.newUID();

		private String name;

		private String city;

		private String state = "";

		private String zip;

		private List<MemberId> members = new ArrayList<>();

		private String email;

		public static OrgBuilder from(Org source) {
			var r = new OrgBuilder();
			r.uuid = source.getUuid();
			r.name = source.getName();
			r.city = source.getCity();
			r.state = source.getState();
			r.zip = source.getZip();
			r.members = source.getMembers();
			r.email = source.getEmail();
			return r;
		}

		public OrgBuilder setUuid(String uuid) {
			this.uuid = uuid;
			return this;
		}

		public OrgBuilder setName(String name) {
			this.name = name;
			return this;
		}

		public OrgBuilder setCity(String city) {
			this.city = city;
			return this;
		}

		public OrgBuilder setState(String state) {
			this.state = state;
			return this;
		}

		public OrgBuilder setZip(String zip) {
			this.zip = zip;
			return this;
		}

		public OrgBuilder setMembers(List<MemberId> members) {
			this.members = members;
			return this;
		}

		public OrgBuilder setEmail(String email) {
			this.email = email;
			return this;
		}

		public Org build() {
			var r = new Org();
			r.uuid = uuid;
			r.name = name;
			r.city = city;
			r.state = state;
			r.zip = zip;
			r.members = members;
			r.email = email;
			return r;
		}
	}
}

