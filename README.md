Avoid repeated three liner err nil checkers.

Following can be replaced 
```
if err != nil {
    log.Fatalf(msg, err)
}
```
with simple 
```
tryerr.HandleErrWithLogFatal(msg, err)
```

Advantage:
```
Clean and clear code
```