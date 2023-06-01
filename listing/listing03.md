Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Вывод:
<nil>
false

В первом случае через происходит вывод его содержимого err(то есть nil указателя)

В втором же случае мы сравниваем err с nil. 
Так как в err хранится информация о типе внутри него, 
он не может считаться nil, и выводится false

Интерфейс можно представить как контракт. 
Любой тип, имеющий все методы какого-то интерфейса, 
будет соответсвовать этому интерфейсу.

Пустой интерфейс не имею никаких методов, и это означает что ему соответсвует любой тип.

```