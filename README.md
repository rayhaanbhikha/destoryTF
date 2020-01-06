# How to use

The cli tool is a wrapper around the terraform cli which will destroy aws resources in a given terraform workspace.

Resources _**must**_ include the following tags.

```js
{
  "Environment": "dev",
  "Owner": "DL-TheUnit-Leeds@dazn.com",
  "Project": "acc-audit",
  "Workspace": workspace, // passed as arg to cli
  "Component": "api", // generic name of component (must be the same as the terraform plan)
  "Type": "branch-builds"
}
```

Currently this _**requires**_ the `Type` tag but this will eventually be removed to enable generic usage.


## Example usage
For instance if you had the following terraform directory with an `api` and `lambda` tf components.
```sh
# ~/project/some-project
|-- terraform
|   |-- api
|   |-- lambda
```

Running the following command would destroy the above terraform resources in the `some-workspace` workspace (`-a` flag means auto approve).
```sh
destroyTF -d ~/projects/some-project/terraform -w some-workspace -a
```