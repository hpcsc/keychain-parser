# keychain-parser

Parse output from MacOS `security` command and output as JSON

### Installation

Download binary from github release and move to a location in your PATH

### Example Usage

```shell
# add a generic password to default keychain
security add-generic-password -a test-account -s test-service -G test-attribute -l test-label -j test-comment

# find above generic password in keychain using service name, pipe to `keychain-parser`
security find-generic-password -s test-service | keychain-parser 

# output:
# [
#   {
#     "class": "genp",
#     "service": "test-service",
#     "account": "test-account",
#     "attribute": "test-attribute",
#     "label": "test-label",
#     "comment": "test-comment"
#   }
#] 

# Dump all items in keychain
security dump-keychain | keychain-parser | jq '...'
```

### Note

This tool only parses a few fields that are useful for me like:

| Column	    | Attribute	                      | Description |
| -------     | ---------                       | ----------- |
| `icmt`	    | `kSecAttrComment`	              | User editable comment for the item |
| `acct`	    | `kSecAttrAccount`	              | Account key (such as user id) |
| `svce`	    | `kSecAttrService`	              | Service name (such as Application identifier) |
| `gena`	    | `kSecAttrGeneric`	              | User defined attribute |

#### Reference for other fields

https://gist.github.com/0xmachos/5bcf2ad0085e09f3b553a88bb0e0574d

| Column	    | Attribute	                      | Description |
| -------     | ---------                       | ----------- |
| `cdat`	    | `kSecAttrCreationDate`	        | Item creation date in Unix epoch time format |
| `mdat`	    | `kSecAttrModificationDate`	    | Item modification date in Unix epoch time format |
| `desc`	    | `kSecAttrDescription`	          | User visible string that describes the item |
| `icmt`	    | `kSecAttrComment`	              | User editable comment for the item |
| `crtr`	    | `kSecAttrCreator`	              | Application created (4 char) code |
| `type`	    | `kSecAttrType`	                | Item type |
| `scrp`	    | `kSecAttrScriptCode`	          | String script code (such as encoding type) |
| `labl`	    | `kSecAttrLabel`                 | Label to be displayed to the user (print name) |
| `alis`	    | `kSecAttrAlias`	                | Item alias |
| `invi`	    | `kSecAttrIsInvisible`	          | Invisible |
| `nega`	    | `kSecAttrIsNegative`	          | Invalid item |
| `cusi`	    | `kSecAttrHasCustomIcon`	        | Existence of application specific icon (Boolean) |
| `prot`	    | `kSecProtectedDataItemAttr`     | ?	Item’s data is protected (Boolean) |
| `acct`	    | `kSecAttrAccount`	              | Account key (such as user id) |
| `svce`	    | `kSecAttrService`	              | Service name (such as Application identifier) |
| `gena`	    | `kSecAttrGeneric`	              | User defined attribute |
| `data`	    | `kSecValueData`                 | Actual data (such as password, crypto key…) |
| `agrp`	    | `kSecAttrAccessGroup`	          | Keychain access group |
| `pdmn`	    | `kSecAttrAccessible`	          | Access restrictions (Data protection classes) |
