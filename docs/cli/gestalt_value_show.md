## gestalt value show

show a single value in a schema

### Synopsis


Show a value specifed by `{schema name}` and `{value name}`. In successful cases, the value will be printed as JSON. This command corresponds to `GET /v1/schemas/{schema name}/values/{value name}`.

```
gestalt value show {schema name} {value name}
```

### Options inherited from parent commands

```
      --host="localhost:3000": API hostname
      --pretty[=true]: pretty-print responses
      --scheme="http": set the request scheme
      --timeout=20s: set the request timeout
```

### SEE ALSO
* [gestalt value](gestalt_value.md)	 - work with values

###### Auto generated by spf13/cobra on 22-Dec-2015
