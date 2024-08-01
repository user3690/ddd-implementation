# Application Layer
- Handles business flow
- Handles security and validation

## Example Flow
1. (Application) Request to import (could be a command too)
2. (Application) Validation of request / Command
3. (Domain) Start to process request
4. (Infrastructure) Sending request to supplier API, parsing response and mapping data
5. (Domain) Further processing data from supplier response
6. (Infrastructure) Mapping and saving data to mysql
7. (Application) Respond to initial request 
