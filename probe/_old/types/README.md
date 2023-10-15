### Probe
Two basic types:
```
type Input struct {
	Key        string
	Aggregator string
	Operator   string
	Value      string
}

type Probe interface {
	Up(*Input) (bool, string)
	GetType() string
	GetDescription() string
}
```

But that simple logic can create huge things :-)

### Primitives
All primitives follow *Probe* interface, the difference in *Input* processing logic:
- Bool - boolean logic ( True or False )
- Number - number logic ( comparison with numbers )
- String - string logic ( substring, length, etc. )
- Map - contains *Probe*-s as map values, apply *Input* to *Input.key* key value
- List - *Input* applied to every item of the list
- Parametrized - *Probe* with predefined *Input* ( ignores any other *Input*)
- Generator - execute function which return *Probe* and apply *Input* to the returned function ( lazy loading )

Basicaly, each probe is one of the primitives above. If probe has attributes ( online, offline, etc. ) - it's *Map* with attribute name as a key and primitive as a value.
