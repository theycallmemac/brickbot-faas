# brickbot-faas

A collection of function configs and code that gives functionality to [brickbot](https://github.com/redbrick/brickbot). These are run using OpenFaaS. 

I've written a tool called `faasup` which actually abstracts a lot of the work in maintaining OpenFaaS away from the user. You can run batch build and deploys for all functions rather than maintaining each individual one.

### faasup

#### Installing

faasup is built in Go. There's a Makefile in here, so to install faasup just run `make`. This should grab all the dependencies you need too. You can view the help guide by running `./faasup --help`. You should always run faasup from this location.

#### Usage 

You can view the help guide by running `./faasup --help`. You should always run faasup from this location and should always use the tool as the root user.

To list the current functions run `faasup list`. 

To build a specific fuction just run `faasup build -f $functionName`. If you want to build all functions just run `faasup build -a`.

To deploy a specific fuction just run `faasup deploy -f $functionName`. If you want to deploy all functions just run `faasup build -a`.

By default the build and deploy subcommands dont log output verbosely. To change this just supply the `--verbose` flag at the end of eithe rof those commands.


### Contibuting

If you want to add some configs and code in here feel free to make a PR! OpenFaaS supports languages and thus many ways to build your function, so go crazy!

### Help Guide

```
manage the brickbot openfaas functions

Usage:
  faasup [command]

Available Commands:
  build       build a specified function
  deploy      deploy a specified function
  help        Help about any command
  info        display info for a specified function
  list        list usable functions
  make        make a new function

Flags:
  -h, --help   help for faasup

Use "faasup [command] --help" for more information about a command.


```
