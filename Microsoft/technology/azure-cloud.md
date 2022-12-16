
## Azure Web Apps (Admin)
Get detailed description
```
[URL_PUBLISHED_WEB]/Env.cshtml
```

```
[URL_WEB]/api/vfs/site/wwwroot/appsettings.json
```

## Application Insights
query filters:
```
# only local PC
| where client_IP == '0.0.0.0'

# started
| where message contains "Consumer started"

# warnings | errors
| where severityLevel > 1


#Warnings
| where severityLevel == 2
```

# Azurite
Default Storage Account [official](https://github.com/Azure/Azurite#usage-with-azure-storage-sdks-or-tools)
```
Account name: devstoreaccount1
Account key: Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==
```