## http

Check URL attributes.

### Examples

```
--url https://example.com online,length>0,code==200
```

### Options

| Check | Description | Check type | Ops | Type |
|--------------|-----------|------------|------------|------------|
| offline | Check if URL offline | bool | None | None |
| online | Check if URL online | bool | None | None |
| code | Check URL status code | conditional | <,<=,==,>=,> | Positive Integer |
| length | Check URL response length | conditional | <,<=,==,>=,> | Positive Integer |