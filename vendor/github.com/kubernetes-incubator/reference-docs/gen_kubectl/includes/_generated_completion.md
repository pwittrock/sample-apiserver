------------

# completion

>bdocs-tab:example Install bash completion on a Mac using homebrew

```bdocs-tab:example_shell
brew install bash-completion
printf "\n# Bash completion support\nsource $(brew --prefix)/etc/bash_completion\n" >> $HOME/.bash_profile
source $HOME/.bash_profile
```

>bdocs-tab:example Load the kubectl completion code for bash into the current shell

```bdocs-tab:example_shell
source <(kubectl completion bash)
```

>bdocs-tab:example Write bash completion code to a file and source if from .bash_profile

```bdocs-tab:example_shell
kubectl completion bash > ~/.kube/completion.bash.inc
printf "\n# Kubectl shell completion\nsource '$HOME/.kube/completion.bash.inc'\n" >> $HOME/.bash_profile
source $HOME/.bash_profile
```

>bdocs-tab:example Load the kubectl completion code for zsh[1] into the current shell

```bdocs-tab:example_shell
source <(kubectl completion zsh)
```


Output shell completion code for the specified shell (bash or zsh). The shell code must be evalutated to provide interactive completion of kubectl commands.  This can be done by sourcing it from the .bash _profile. 

Note: this requires the bash-completion framework, which is not installed by default on Mac.  This can be installed by using homebrew: 

  $ brew install bash-completion
  
Once installed, bash completion must be evaluated.  This can be done by adding the following line to the .bash profile 

  $ source $(brew --prefix)/etc/bash_completion
  
Note for zsh users: [1] zsh completions are only supported in versions of zsh >= 5.2

### Usage

`$ completion SHELL`



