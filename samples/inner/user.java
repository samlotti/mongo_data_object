package com.hapticapps.amici.shared.data_models.org;

import com.hapticapps.amici.shared.utils.Utils;
import org.bson.codecs.pojo.annotations.BsonProperty;

public class User {
	public static final String BSON_UUID = "u";
	public static final String BSON_NAME = "n";
	public static final String BSON_EMAIL = "e";

	@BsonProperty("u")
	private String uuid;

	@BsonProperty("n")
	private String name;

	@BsonProperty("e")
	private String email;


	public String getUuid(){
		return uuid;
	}
	public String getName(){
		return name;
	}
	public String getEmail(){
		return email;
	}

	public UserBuilder copy() {
		return UserBuilder.from( this );
	}
	public String toString() {
		return "User{" + 
			"}";
	}

	public static class UserBuilder {

		private String uuid = Utils.newUID();

		private String name;

		private String email;

		public static UserBuilder from(User source) {
			var r = new UserBuilder();
			r.uuid = source.getUuid();
			r.name = source.getName();
			r.email = source.getEmail();
			return r;
		}

		public UserBuilder setUuid(String uuid) {
			this.uuid = uuid;
			return this;
		}

		public UserBuilder setName(String name) {
			this.name = name;
			return this;
		}

		public UserBuilder setEmail(String email) {
			this.email = email;
			return this;
		}

		public User build() {
			var r = new User();
			r.uuid = uuid;
			r.name = name;
			r.email = email;
			return r;
		}
	}
}

