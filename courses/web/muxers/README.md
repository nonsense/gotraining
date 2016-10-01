## Ultimate Web - Alternative Muxers
The basic muxer in Go has gotten us a long way by this point, but it has it’s limitations. Let’s tour three very different types muxers/routers.

### Pat ([github.com/gorilla/pat](github.com/gorilla/pat))

*(compatible with the standard library)*

[RESTful Server](../../../topics/web/muxers/example1/main.go)

#### Exercise

Implement the `PUT`, `PATCH`, and `DELETE` responses in the [RESTful Pat Server](../../../topics/web/muxers/example1/main.go).

### HttpRouter

*(non-compatible with the standard library)*

[RESTful Server](../../../topics/web/muxers/example2/main.go)

#### Exercise

Implement the `PUT`, `PATCH`, and `DELETE` responses in the [RESTful Httprouter Server](../../../topics/web/muxers/example2/main.go).

### Echo

*(a micro-framework)*

[RESTful Server](../../../topics/web/muxers/example3/main.go)

#### Exercise

Implement the `PUT`, `PATCH`, and `DELETE` responses in the [RESTful Echo Server](../../../topics/web/muxers/example3/main.go).

*Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).