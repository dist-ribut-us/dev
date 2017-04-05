## Generator
Simple generator for handling boilerplate code. Reads in a json file and outputs
a go file, all on stdin/stdout. This is not meant to be a general purpose code
generation tool, but a tools specific to the needs of the dist.ribut.us project.
Any generated code should still be checked in.

### Generator File
A generator file is json, it has two top level fields: Package and Imports.
Package is the package name and expects a string. Imports is a list of imports.
The generator will reduce the imports if something is repeated and code
generation tools may add other imports.

### Threadsafe Maps
In the generator file the field is TSMaps and it expects a list of ThreadSafe
Map objects. Each object should have Key, Value and Name as strings. Key and
Value are types and name is the name of the struct that will be generated. and
optionally Export can be set to true which will make the new func and all
methods exported.

A threadsafe map combines a map with RWMutex. It has Get, Set and Delete
methods. I considered adding a way to iterate over the map, but there is no
good way to handle that for all cases. The mutex is embeded.