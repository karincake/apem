# Handler Helper
Package that helps with writing http output through various processes, therefore this package should be called by the handler and never called directly by the services. Due to its nature, all of its functions requres httpwriter as one of the parameter, and whenever error occured it writes error directly from it.

Note that many of its output depends on the package tempe data and error (https://github.com/karincake/tempe)

## Validation
There are many validation available in the pacakage. Some of its validation process utilizes package serabi (https://github.com/karincake/serabi)


## Errors Output
There are 2 primary types of errors:
1. Data field error, that can have multiple errors at once and should always uses unprocessable entity error, uses te.XErrors format
2. Non data field error, that only has single error, uses te.XError format
