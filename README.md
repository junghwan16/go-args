```go
// go-args highly inspired by the coding dojo kata `Args`
// see: https://codingdojo.org/kata/Args/

spec := args.New().
	Boolean("l", args.WithHelp("enable logging")).
	Int("p", args.WithDefault(0), args.WithHelp("port number")).
	String("d", args.WithHelp("directory")).
	FloatSlice("nums", args.WithHelp("comma separated"))


res, err := spec.Parse([]string{"-l","-p","8080","-d","/var","-nums","1,2,3"})
if err != nil { /* handle */ }

fmt.Println(res.IsValid())				// true
fmt.Println(res.Bool("l"))        // true
fmt.Println(res.Int("p"))         // 8080
fmt.Println(res.String("d"))      // /var
fmt.Println(res.FloatSlice("nums")) // [1 2 3]
```
