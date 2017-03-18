------------

# auth



Inspect authorization

### Usage

`$ auth`



------------

## <em>can-i</em>

>bdocs-tab:example Check to see if I can create pods in any namespace

```bdocs-tab:example_shell
kubectl auth can-i create pods --all-namespaces
```

>bdocs-tab:example Check to see if I can list deployments in my current namespace

```bdocs-tab:example_shell
kubectl auth can-i list deployments.extensions
```

>bdocs-tab:example Check to see if I can get the job named "bar" in namespace "foo"

```bdocs-tab:example_shell
kubectl auth can-i list jobs.batch/bar -n foo
```


Check whether an action is allowed. 

VERB is a logical Kubernetes API verb like 'get', 'list', 'watch', 'delete', etc. TYPE is a Kubernetes resource.  Shortcuts and groups will be resolved. NAME is the name of a particular Kubernetes resource.

### Usage

`$ can-i VERB [TYPE | TYPE/NAME]`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
all-namespaces |  | false | If true, check the specified action in all namespaces. 
quiet | q | false | If true, suppress output and just return the exit code. 



