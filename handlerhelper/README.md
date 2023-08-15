# Handler Helper
Package that helps with writing http output through various process therefore this package should be called by the handler and never called directly by the services. Due to its nature, all of its functions requres httpwriter as one of the parameter. Its main validation process uses pakcage serabi (https://github.com/karincake/serabi)

## The HTTP Output
All functions return 2 values: data and error

### Data
The format being used is from package tempe's data (https://github.com/karincake/tempe/data)

### Errors
Most the the time processing data requires validation which might result errors in the process. There are 2 types of errors:
1. Non-field error, which uses single error
2. field error, which uses multiple errors

The format being used is from package tempe's error (https://github.com/karincake/tempe/error).
