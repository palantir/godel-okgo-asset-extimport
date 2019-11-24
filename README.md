DEPRECATED
==========
The extimport check was originally developed to ensure the consistency of the vendor directory for projects. This was
necessary when the vendor directory was managed manually or using basic tools such as [govendor](https://github.com/kardianos/govendor).
However, tools such as [dep](https://github.com/golang/dep) and the support for [modules in Go](https://blog.golang.org/using-go-modules) 
have made the functionality provided by this tool unnecessary. As such, active development on this project has ended.

godel-okgo-asset-extimport
==========================
godel-okgo-asset-extimport is an asset for the gödel [okgo plugin](https://github.com/palantir/okgo). It provides the
functionality of the [go-extimport](https://github.com/palantir/go-extimport) check.

This check verifies that a project does not import any packages that are external to the project (not part of the
project or its vendor directories).
