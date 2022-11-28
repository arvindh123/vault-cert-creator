# Prequest before running the program    
### STEP 1 :    
Run command :  ```vault server -dev```     
After running this command you will ```Root Token``` printed in CLI , Please copy and paste in the program in Line 12 TOKEN=\<<Root_Token\>>      

### STEP 2:    
Run command :  ```export VAULT_ADDR=http://127.0.0.1:8200```      

### STEP 3
Run command :  ```vault secrets enable pki```    

### STEP 4:   v
Run command: ```vault secrets tune -max-lease-ttl=87600h pki```    

### STEP 5:
Run command: ```vault write pki/root/generate/internal common_name=myvault.com ttl=87600h```     

### STEP 4:
Run command: ```vault write pki/config/urls issuing_certificates="http://vault.example.com:8200/v1/pki/ca" crl_distribution_points="http://vault.example.com:8200/v1/pki/crl"```      

### STEP 5:
Run command: ```vault write pki/roles/example-dot-com allowed_domains=example.com allow_subdomains=true max_ttl=72h```     

## What does this program does ?
This program will generate and revoke the certificate continuously         
The below two steps are automated in this go program       
Certificate generate/issue Command :    ```vault write pki/issue/example-dot-com common_name=blah.example.com```       
Certificate Revoke Command         :    ```vault write pki/revoke serial_number=5e:7e:32:6c:66:e2:e4:04:33:b8:bd:f6:ea:5d:1f:97:01:d0:7e:aa```     


### Realted to Vault Limits 
https://developer.hashicorp.com/vault/docs/internals/limits


### Related to vault tidy 
https://www.hashicorp.com/blog/certificate-management-with-vault#:~:text=Vault%20will%20maintain%20expired%20certificates%20for%20a%20certain%20buffer%20period.%20To%20optimize%20Vault%E2%80%99s%20storage%20backend%20and%20CRL%2C%20use%20the%20tidy%20endpoint%20to%20remove%20expired%20certificates%20from%20Vault.

https://developer.hashicorp.com/vault/api-docs/secret/pki#tidy