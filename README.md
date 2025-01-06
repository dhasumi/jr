# jr

This is a custom Jira CLI wrapper command for personal use.

## Prerequisites

- [jira-cli](https://github.com/ankitpokhrel/jira-cli) (Ensure `jira` command is available and properly set up)
- golang (Only if you build it from source code)

## Installation

Please download a compiled binary from [Release Page](https://github.com/dhasumi/jr/releases), or build from source code with the below command.

```
go install github.com/dhasumi/jr@latest
```

## Usage

The basic usage of the `jr` command with `create` subcommand is as follows. In the example below, the string specified is treated as the summary of a Task-type ticket using the `jira` command. The numeric value specified with the `--sp` option is reflected to the Story Points field. If not explicitly specified, the ticket will be assigned to the user themselves and placed in the current sprint.

```
jr create "summary string" --sp 4
```

### Optional parameters

Other optional parameters are available as follows:

`-t STRING`, `--type STRING`:
Allows the creation of ticket types other than `Task`. Specify the desired type using this option.

`-b STRING`, `--body STRING`:
Specifies a string to be entered into the ticket's Description field.

`-e STRING`, `--epic STRING`:
Specifies an Epic ID. Tickets created with this option will be associated with the given Epic. The Epic ID needs to be identified using the `jira` command. (Example: `ABC-123`)

`-a STRING`, `--assign STRING`:
Assigns the created ticket to a user other than yourself. The specified string must be a valid email address or username recognizable by `jira` command. The default assigner is the user themselves.

`-r STRING`, `--report STRING`:
Assigns the created ticket to a user other than yourself as a ticke reporter. The specified string must be a valid email address or username recognizable by `jira` command. The default reporter is the user themselves.

`-y STRING`, `--priority STRING`:
Sets the priority of the ticket. e.g. {Highest, High, Medium (default), Low, Lowest}

`-p STRING`, `--project STRING`:
Sets the project name you want to create ticket on.

`-l STRING`, `--label STRING`:
Adds the specified string as a label to the ticket. If you want to add multiple labels, separate them with commas (,) in a single option.

`-c STRING`, `--component STRING`:
Adds the specified string as a component to the ticket. If you want to add multiple component, separate them with commas (,) in a single option.

`-s NUM`, `--sprint NUM`:
Specifies a sprint number. Tickets will be created in the sprint corresponding to the provided number (SPRINT NAME). The sprint number must correspond to upcoming sprint, and the sprint must already exist.

`--next-sprint:`
Places the ticket in the next sprint following the current one. If used in conjunction with `-e` option, `-e` option takes precedence.

`--future-sprint NUM`:
Places the ticket in the sprint occurring a specified number of sprints after the current one. For example, `--future-sprint 1` is equivalent to `--next-sprint option`.

`--backlog`:
Specifies created ticket will be located on backlog.

`--template FILE_PATH`:
Specifies a template file, similar to the functionality of the `jira` command. If used together with `-b` option, this option takes precedence.
