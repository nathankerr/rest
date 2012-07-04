/*
A RESTful HTTP client and server.

rest.go uses the standard http package by adding resource routes. Add
a new route by:

	rest.Resource("resourcepath", resourcevariable)

and then use http as normal.

A resource is an object that may have any of the following methods which
respond to the specified HTTP requests:

	GET /resource/ => Index(http.ResponseWriter)
	GET /resource/id => Find(http.ResponseWriter, id string)
	POST /resource/ => Create(http.ResponseWriter, *http.Request)
	PUT /resource/id => Update(http.ResponseWriter, id string, *http.Request)
	DELETE /resource/id => Delete(http.ResponseWriter, id string)
	OPTIONS /resource/ => Options(http.ResponseWriter, id string)
	OPTIONS /resource/id => Options(http.ResponseWriter, id string)

The server will then route HTTP requests to the appropriate method call.

The snips example provides a full example of both a client and server.
*/
package rest
