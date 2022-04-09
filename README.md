# study-generics-perf
Study on how generics may affect overall performance of real life applications

```go
package main

import(
    "log"
    "time"
)

func[T any]f(a T){
    log.Printf("%#v",a)
}

func main(){
    f(int64(1))
    f([]byte("ASD"))
    f(time.Now())
}
```