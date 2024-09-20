

- https://sdkman.io/install
- copy 7zip.exe and rename as 'zip.exe' to:
  C:\Program Files\Git\mingw64\bin
- open gitBash

sdk list java  (press 'space bar' as 'next page')
sdk install java 22.0.2-oracle
sdk install java 21.0.4-oracle
sdk install java 17.0.12-oracle

sdk use java 21.0.4-oracle
sdk default java 17.0.12-oracle

# show current version
sdk current java

# show path downloaded
sdk home java 17.0.12-oracle
