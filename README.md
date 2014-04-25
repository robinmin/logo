logo
======

A lightweight log in golang with the capability to extend as standard log package.

What's an idea log package? In my view, the idea log package should be:

-  Object free : No any burden to to create object or initialize the logger;
-  Compact with the standard log interface, so that any one with io.Writer interface
	can be leveraged;
-  Level based : So that end user can customize the output at a certain level;
-  Multiple loggers integrated;
-  Simple enough : Have no chance to create bug;

logo is a implementation base on this philosophy; You can use it as showing below:

```go
import (
	"bytes"
	log "github.com/robinmin/logo"
)

func main() {
	// you can just use the log package like, the system will output the log to stdout by default
	log.Debug("Balabala,bala......")

	// If you have a customized logger, you can use it like the following. The system will output
	//  the log into both stdout and buf
	var buf = bytes.NewBuffer(nil)
	log.AddLogger("buffer", buf, log.ALL)
	defer log.ReleaseLogger("buffer")
	str := "Buffer one"
	log.Debug(str)
}

```
