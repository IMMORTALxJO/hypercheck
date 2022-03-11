## fs

Check file attributes.

### Examples

```
--fs '/var/log/app' dir,exists,owner=user,group=group,size<10Mb
```

### Options

| Check | Description | Check type | Ops | Type |
|--------------|-----------|------------|------------|------------|
| exists | Check if file glob is not empty | bool | None | None |
| regular | Check if match is a regulare file | bool | None | None |
| dir | Check if match is a directories | bool | None | None |
| count | Check file glob count | conditional | <,<=,==,>=,> | Positive Integer |
| size | Check size of each found file | conditional | <,<=,==,>=,> | Bytes |
| totalsize | Check total size of all files | conditional | <,<=,==,>=,> | Bytes |
| uid | Check owner user id | conditional | <,<=,==,>=,> | Positive Integer |
| gid | Check owner group id | conditional | <,<=,==,>=,> | Positive Integer |