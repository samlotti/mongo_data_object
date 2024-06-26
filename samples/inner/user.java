package com.hapticapps.amici.shared.data_models.org;
// Generated by mdo do not edit this file, see the .mdo file 

import com.hapticapps.amici.shared.utils.Utils;
import org.bson.codecs.pojo.annotations.BsonProperty;
import java.util.Map;
import java.util.List;
import java.util.ArrayList;
import com.mongodb.client.model.IndexOptions;
import org.bson.conversions.Bson;
import com.mongodb.BasicDBObject;

public class User {
	public static final String BSON_UUID = "u";
	public static final String BSON_NAME = "n";
	public static final String BSON_EMAIL = "e";

	@BsonProperty("u")
	private String uuid = Utils.newUID();

	@BsonProperty("n")
	private String name;

	@BsonProperty("e")
	private String email;


	public String getUuid(){
		return uuid;
	}
	public void setUuid(String data){
		this.uuid = data;
	}
	public String getName(){
		return name;
	}
	public void setName(String data){
		this.name = data;
	}
	public String getEmail(){
		return email;
	}
	public void setEmail(String data){
		this.email = data;
	}

	public UserBuilder copy() {
		return UserBuilder.from( this );
	}

	public static UserBuilder builder() {
		return new UserBuilder();
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

		public UserBuilder withUuid(String uuid) {
			this.uuid = uuid;
			return this;
		}

		public UserBuilder withName(String name) {
			this.name = name;
			return this;
		}

		public UserBuilder withEmail(String email) {
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
	public static class Indexes {
		public static List<Bson> ikeys = new ArrayList<>();
		public static List<IndexOptions> ioptions = new ArrayList<>();
		static {
		}
	}
}

