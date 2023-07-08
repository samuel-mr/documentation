# DataGrip

| Shortcut                              | Description                             |
|---------------------------------------|-----------------------------------------|
| Ctrl + Q (in explorer)                | Show metadata inline                    |
| F1 (over table in sql statement)      | Show metadata + select top 10           |
| Alt + Enter (over * in sql statement) | Show refactor to expand the column list |
| Alt + Enter (over table name)         | Introduce aliase                        |
| Ctrl + Alt + t                        | Surround with...                        |
| Shift + F6                            | Rename...                               |

| Shortcut                              | Description                             |
|---------------------------------------|-----------------------------------------|
| ins                | Live template Insert                    |
| (select columns in conflic) + Alt + Enter               | Add Alias                     |

# Config
## When add connection config
For some servers, In `General / URL` add:
```
;trustServerCertificate=true
```
## Code completion
```
Settings | Editor | General | Code Completion | SQL
 - Sugguest object from
    - All available schemas
 - Qualify objects with:
    - Table : Always
```