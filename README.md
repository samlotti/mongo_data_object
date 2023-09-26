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

go install github.com/samlotti/mongo_data_object@latest

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
class OrgPerson {

    data String uuid as u = Utils.newUID();

    data String orgId as o;

    data String name as n;

    data String email as e;

}

```

The output is: [output java](internal/sample_output.txt)

```
    Usage of the output class:
    
    var org = OrgPerson.builder().setOrgId("45").setName("test").build();
    var orgUpdate = org.copy().setName("newName").build();

```

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

# Contributing

All contributions are welcome – if you find a bug please report it.

# Contributors

* Sam Lotti (support@hapticappsllc.com)
