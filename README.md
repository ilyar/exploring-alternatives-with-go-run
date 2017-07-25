Example code for https://pocketgophers.com/exploring-alternatives-with-go-run/

Get the code with:

```
go get -d pocketgophers.com/exploring-alternatives-with-go-run
```

Run the examples with:

```
go run main.go messy.go |tee messy.txt
```

and

```
go run main.go clean.go |tee clean.txt
```

then compare the results with:

```
benchstat messy.txt clean.txt
```
