# BOFHchan

## Desarrollo

Para probar BOFHchan en local, basta con ejecutar:

```bash
go run cmd/bofhchan/main.go
```

BOFHchan está diseñado para ser autocontenido. Por lo que depende de pkger para "embeber" todos sus ficheros estáticos y plantillas, y es invocado por go generate.

En desarrollo (modo "debug") los ficheros estáticos y plantillas se leen de disco siempre, ya que ejecutar el comando anterior hace que se ignoren los demás fuentes en el directorio `cmd/bofhchan`.

### URLs

Es importante tener en cuenta que las URLs tienen precedencia. Los dos siguientes ejemplos no se comportan de la misma manera:

_`/submit` invoca `publish`, el resto de URLs van a `get`._

```go
    router.Get("/submit", publish)
    router.Get("/:id", get)
```

_Todas las URLs van a `get`, `/submit` nunca es invoca._

```go
    router.Get("/:id", get)
    router.Get("/submit", publish)
```

## Producción

Para compilar, deben ejecutarse los siguientes pasos:

```bash
go generate ./...
go build -o bofhchan ./cmd/bofhchan/...
```

La compilación con los comandos anteriores, o cambiando el `go build ...` por `go install ./...`, activa el modo "release" y los ficheros embebidos con pkger.
