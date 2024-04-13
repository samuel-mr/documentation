# Assemblies

- Basic unit of deployment. Defines a boundary for type resolution
- Container of:
  - compiled types (IL)
  - runtime resources
  - versioning and referencing other assemblies
- .dll extension

when build and executable app:
| .Net Framework | .Net 6|
| ------------ | --------------------- |
| Portable executable (PE) assembly | an assembly (.dll)|
| Ctrl + m + m | an executable launcher (.exe) for some platform|
