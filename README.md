martini-data is a martini middleware wrapper around
github.com/albrow/go-data-parser. It automatically parses
data from the request body and the url query parameters,
provides helper methods for converting data to other types,
and supports validations.

Installation and Usage
----------------------

Install like you would any other package:
```
go get github.com/albrow/martini-data
```

Don't forget to import:
```
import "github.com/albrow/martini-data"
```

Then add to your martini middleware stack:
```
m.Use(data.Parser())
```

Now you can access a Data object in your handler functions:
```
func MyHandler(data data.Data, res http.ResponseWriter) {
	// ...
}
```

Note: the Data type comes from the github.com/albrow/go-data-parser
package, so you will need to import it as well.
```
import "github.com/albrow/go-data-parser"
```

See github.com/albrow/go-data-parser for more information about Data
and Validator.