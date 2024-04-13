# CodeQL
```
codeql database create erwebdb --language=csharp
codeql database analyze erwebdb --language=csharp
```
- Download
https://docs.github.com/en/code-security/codeql-cli/getting-started-with-the-codeql-cli/setting-up-the-codeql-cli#1-download-the-codeql-cli-zip-package

- Unzip + Add to PATH env
```bash
# at the same level (folder) of codeql
git clone --recursive https://github.com/github/codeql.git codeql-repo

# we need to have 2 folders at the same level
codeql
codeql-repo

# Check the configuration
codeql resolve languages
codeql resolve qlpacks

# move to the project folder and...
codeql database create erwebdb --language=csharp
codeql database analyze erwebdb --rerun --format=sarif-latest --output=result/reporte.sarif --threads=8
```