# Azure Cli
basic
```
az login --use-device-code   // not local browser neccesarilly

az login --allow-no-subscriptions  //login with tenant level access
az login --tenant [tenantId] (aveces no se ve el tenanId, pero tenantId = DirectoryId ... tons ir al Portal: from Portal Settings / Directory + Subscriptions)
```
basic (to Rider)
```
# to validate current login
az account show

# see the default tenant
az account list

# set the default tennant (Rider needs the correct Tenant as default)
az account set --subscription [tenantId]
```
Miscelaneos
```
az SOMERESOURCE -h     // help for SOMERESOURCE
```
# resources
## Select
```
az group list
az group list --query "[].name"
az group list --query "[].{resourceGroupName:name,location:location}"
az group list --query "[].{nombre:name,ubicacion:location}"
```
## Filters
```
az group list --query "[?location=='eastus']"
az group list --query "[?location=='eastus'].name"      // filter + select 
az group list --query "[?location=='eastus']" -o table  // filter + format as table
```

## Formats
```
az group list -o tsv        # csv
az group list -o jsonc      # json colorized
az group list -o table      # table
```

### examples:
```
# names of all resource groups
az group list --query "[?type=='Microsoft.Resources/resourceGroups'].name"
```
# ServiceBus
```
az servicebus topic list --namespace-name NAME --resource-group ID  --query "[].createdAt"
```
# Extensions

## Resource Graph 
[Add resource graph extension](https://docs.microsoft.com/en-us/azure/governance/resource-graph/first-query-azurecli#add-the-resource-graph-extension)
```
az extension add --name resource-graph
az graph query -h                           # help
```

#### Use

```
# Top 5 first resources
az graph query -q 'Resources | project name, type | limit 5'

# Run Azure Resource Graph query with 'order by'
az graph query -q 'Resources | project name, type | limit 5 | order by name asc'

# Run Azure Resource Graph query with `order by` first, then with `limit`
az graph query -q 'Resources | project name, type | order by name asc | limit 5'

# WHERE
az graph query -q "project id, type, location | where type =~ 'Microsoft.Compute/virtualMachines' | summarize count() by location | top 3 by count_"

# 
az graph query -q "project id, type, location | where type =~ 'Microsoft.Compute/virtualMachines' | summarize count() by location | top 3 by count_"
```
## Application insight
extension installed automatically
```
az monitor app-insights component show --app [resourceName] -g [groupName]

# gets Insight's resource filtered by InstrumentationKey 
az monitor app-insights component show --query "[?instrumentationKey=='IDDDD']"
```