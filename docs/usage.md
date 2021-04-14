# Usage

## Profiling redis configuration

```shell
$ insecticide --redis-config test/data/redis.conf --redis-version 6
Parameter: requirepass
[CRITICAL]
Advice: Parameter requirepass is not set. Index: 0
Reason: Set variable for Param requirepass and Index 0

Parameter: appendonly
[WARNING]
Advice: Read this page and make decision: https://redis.io/topics/persistence
Reason: Persistence disabled. Your data just exists as long as the server is running

Parameter: timeout
[WARNING]
Advice: Set timeout
Reason: If timeout is 0, clients connections won't be closed. They will be in idle status.

Parameter: loglevel
[WARNING]
Advice: In production use a less aggressive logging policy (notice or warning)
Reason: Many rarely useful info, but not a mess like the debug level

```

## Help
For help do following command
```shell
$ insecticide --help
```
