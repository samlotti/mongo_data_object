# MDO - Mongo data objects

MDO is a java code generator to create immutable data objects with enhancements for supporting mongo db.

The generated code includes builders and copy support of the objects.  

Features:

* Small easy to read data declarations.
* Support immutable programming.
* Auto code generation.
* Builder for the object type.
* Copy to builder for creating modified versions.
* Index declaration.
* Mongo field name support for reduced field names in the database to save storage.
* Inner enums and inner data objects.

## Installing

Note that mdo is written in go, so go will need to be installed.

Install the command line tool

    go install github.com/samlotti/mongo_data_object

    or if cloning the repository

    go install mdo.go

## Running
In the root directory of your java project, or the root of where your mongo entities will reside.
Type
>  mdo

This will search the current directory and **child** directories for files ending in .mdo

It will read the mdo file and create java files of the same name in the same directory.

>   mdo -help
> 
>   mdo -rebuild

The default is to only rebuild java files that are older than the mdo file. If you want to rebuild all then pass -rebuild.

Example OrgPerson.mdo will create OrgPerson.java in the same directory of the .mdo.
Note the package name, in this case the .mdo is in the directory: com/hapticapps/amici/shared/data_models/org  

```
package com.hapticapps.amici.shared.data_models.org;

import com.hapticapps.amici.shared.utils.Utils;
import org.bson.codecs.pojo.annotations.BsonProperty;

/**
    The org person
**/
entity OrgPerson {

    index (uuid asc) unique;
    index (name asc) unique;
    index (email asc) unique sparse;
    index (friendIds asc) sparse;

    data String uuid as u = ~Utils.newUID()~;

    data String orgId as o;

    data String name as n;

    data String email as e;

    data Status status as st = ~Status.PENDING~;

    data ~List<String>~ friendIds as f = ~new ArrayList<>()~;

    data ~List<Address>~ addresses as a = ~new ArrayList<>()~;
    
    ...

```

The output is: [output java](samples/sample_input.java)

```
    Usage of the output class:
    
    var org = OrgPerson.builder().setOrgId("45").setName("test").build();
    
    // Update the org.  Note the original is not modified.
    var orgUpdate = org.copy().setName("newName").build();

    
```

# Examples

See the [samples](samples) folder.
* [org.mdo](samples/acct.mdo)
* [generated java class](samples/acct.java)

# Mongo

This project only contains support for data objects that can be persisted in mongo.

The generated class should fit in your current mongo project, ex:
```java
    
    collection.insertOne( org );
    
    collection.findOne(eq(OrgPerson.BSON_NAME, "test"));
    
```



These are to be treated as POJO for the mongo services, so the codec registry should be configured:

```
    ConnectionString connectionString = new ConnectionString(config.url);

    CodecRegistry pojoCodecRegistry = fromProviders(PojoCodecProvider.builder().automatic(true).build());
    CodecRegistry codecRegistry = fromRegistries(MongoClientSettings.getDefaultCodecRegistry(),
            pojoCodecRegistry);


    MongoClientSettings clientSettings = MongoClientSettings.builder()
            .applyConnectionString(connectionString)
            .codecRegistry(codecRegistry)
            .build();
```

The mdo grammar.
```
    package package.name;
    
    /**
        Comments
    **/
    import ... ;
    
    entity ClassName {
        
           index ( fieldName asc/desc [, fieldName asc/desc ] ) unique sparse background;
           .. index ..
           
           /**
           Data represents a field.
           Type is the java type (can be a class type)
           ~string~ are identifiers but can have special characters.  For example an
           array.  
           
           data type field [as bsonName] [= initialValue];
           
           */
           data ~Type~ uuid as u = ~new Id()~;
           data ~List<Product>~ products as p = ~new ArrayList<>()~;
    
    } 
    
    enum ClassName {
        Val, Val ...
    }
    
    class ClassName {
        /**
            data fields like the entity.
        */
    }
    
    [enums, classes]

```

# Why setters?

Mongo needs setters and getters for reading and writing the data.  At this time I have not found a way around this.  
If you want immutable, just use the builder and copy methods to update the data.

# Contributing

All contributions are welcome – if you find a bug please report it.

# Contributors

* Sam Lotti (support@hapticappsllc.com)
