# MDO - Mongo data objects

## Installing

go install mdo.go

## Running
In the root directory of your java project, or the root of where your mongo entities will reside.
Type
>  mdo

This will search the current directory and child directories for files ending in .mdo

It will read the mdo file and create java files of the same name in the same directory.

>   mdo -help
> 
>   mdo -rebuild

The default is to only rebuild java files that are older than the mdo file. If you want to rebuild all then pass -rebuild.




