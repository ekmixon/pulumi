
### Improvements
  
- [codegen/dotnet] - Emit dynamic config-getters.
  [#7549](https://github.com/pulumi/pulumi/pull/7549)

- [sdk/python] - Support for authoring resource methods in Python.
  [#7555](https://github.com/pulumi/pulumi/pull/7555)

### Bug Fixes

- [cli] - `pulumi stack ls` now returns all accessible stacks (removing
  earlier cap imposed by the httpstate backend).
  [#3620](https://github.com/pulumi/pulumi/issues/3620)
- [sdk/go] - Fix target and replace options for the Automation API
  [#7426](https://github.com/pulumi/pulumi/pull/7426)
