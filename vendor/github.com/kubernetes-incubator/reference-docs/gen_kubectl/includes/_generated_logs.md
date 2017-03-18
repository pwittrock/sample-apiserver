------------

# logs

>bdocs-tab:example Return snapshot logs from pod nginx with only one container

```bdocs-tab:example_shell
kubectl logs nginx
```

>bdocs-tab:example Return snapshot logs for the pods defined by label app=nginx

```bdocs-tab:example_shell
kubectl logs -lapp=nginx
```

>bdocs-tab:example Return snapshot of previous terminated ruby container logs from pod web-1

```bdocs-tab:example_shell
kubectl logs -p -c ruby web-1
```

>bdocs-tab:example Begin streaming the logs of the ruby container in pod web-1

```bdocs-tab:example_shell
kubectl logs -f -c ruby web-1
```

>bdocs-tab:example Display only the most recent 20 lines of output in pod nginx

```bdocs-tab:example_shell
kubectl logs --tail=20 nginx
```

>bdocs-tab:example Show all logs from pod nginx written in the last hour

```bdocs-tab:example_shell
kubectl logs --since=1h nginx
```

>bdocs-tab:example Return snapshot logs from first container of a job named hello

```bdocs-tab:example_shell
kubectl logs job/hello
```

>bdocs-tab:example Return snapshot logs from container nginx-1 of a deployment named nginx

```bdocs-tab:example_shell
kubectl logs deployment/nginx -c nginx-1
```


Print the logs for a container in a pod or specified resource. If the pod has only one container, the container name is optional.

### Usage

`$ logs [-f] [-p] (POD | TYPE/NAME) [-c CONTAINER]`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
container | c |  | Print the logs of this container 
follow | f | false | Specify if the logs should be streamed. 
include-extended-apis |  | true | If true, include definitions of new APIs via calls to the API server. [default true] 
interactive |  | false | If true, prompt the user for input when required. 
limit-bytes |  | 0 | Maximum bytes of logs to return. Defaults to no limit. 
previous | p | false | If true, print the logs for the previous instance of the container in a pod if it exists. 
selector | l |  | Selector (label query) to filter on. 
since |  | 0s | Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used. 
since-time |  |  | Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used. 
tail |  | -1 | Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided. 
timestamps |  | false | Include timestamps on each line in the log output 


