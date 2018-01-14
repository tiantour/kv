# kv
kv is a key/value store base on badger db

# How to use

### write

List

```
package main

import (
	"log"
	"time"

	"github.com/tiantour/kv"
)

func main() {
	args := map[string][]byte{
		"key1": []byte("value1"),
		"key2": []byte("value2"),
		"key3": []byte("value3"),
	}
	err := kv.NewWrite().List(args)
	if err != nil {
		log.Fatal(err)
	}
}
```
Item

```
package main

import (
	"log"
	"time"

	"github.com/tiantour/kv"
)

func main() {
	err := kv.NewWrite().Item("key1", []byte("value1"))
	if err != nil {
		log.Fatal(err)
	}
}
```
### read

List

```
package main

import (
	"fmt"

	"github.com/tiantour/kv"
)

func main() {
	result := kv.NewRead().List("key1", "key2", "key3")
	for k, v := range result {
		fmt.Printf("key = %s, value = %s\n", k, v)
	}
}
```
Item
```
package main

import (
	"fmt"
	"log"

	"github.com/tiantour/kv"
)

func main() {
	result, err := kv.NewRead().Item("key1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("value = %s\n", result)
}
```
### key

All
```
package main

import (
	"fmt"

	"github.com/tiantour/kv"
)

func main() {
	size := 100
	result := kv.NewKey().All(size)
	for k, v := range result {
		fmt.Printf("key = %s, value = %s\n", k, v)
	}
}
```

Keys
```
package main

import (
	"fmt"

	"github.com/tiantour/kv"
)

func main() {
	size := 100
	result := kv.NewKey().Keys(size)
	for k, v := range result {
		fmt.Printf("index = %d, key = %s\n", k+1, *v)
	}
}
```

Prefix

```
package main

import (
	"fmt"

	"github.com/tiantour/kv"
)

func main() {
	size := 10
	result := kv.NewKey().Prefix("key", size)
	for k, v := range result {
		fmt.Printf("key = %s, value = %s\n", k, v)
	}
}
```

Delete

```
package main

import (
	"log"

	"github.com/tiantour/kv"
)

func main() {
	err := kv.NewKey().Delete("key1", "key2", "key3")
	if err != nil {
		log.Fatal(err)
	}
}
```

Unique

```
package main

import (
	"fmt"
	"log"

	"github.com/tiantour/kv"
)

func main() {
	result, err := kv.NewKey().Unique("x", 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
```
