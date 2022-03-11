```
--fs '/tmp/file' file,exists,mode=7xx,owner=user,group=group,size>10Mb
```

### Options

| exists | Check if file glob is not empty | bool | None | None |
| regular | Check if match is a regulare file | bool | None | None |

| count | Check file glob size | conditional | <,<=,==,>=,> | Positive Integer |
| size | Check size of each found file | conditional | <,<=,==,>=,> | Bytes |
