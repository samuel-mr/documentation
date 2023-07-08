# Global variables
```
process             # diferent information of Node environment
    .versions   #diferent lib and tools used

os                  # os information
    .userInfo() 
    

# current running process of: dir & file 
__dirname       # >W:\Learn\Web\Node\faztcourse
__filename      # >W:\Learn\Web\Node\faztcourse\globals.js

# allow separte app in multiple files
module          # variable 
    .export     # array with all exported files

# allow import all separated files
require         # function
```
```
# run and exit
node -e "console.log(10+10)"
```
# Global functions
```
setInterval( [func] , [mls])        # recurrent
setTimer( [func] , [mls])           # once

clearInterval                       # to stop setInterval
```

# Import
```
require('./module/myModuleName')        # 
const somename = require('......')      # with alias
```