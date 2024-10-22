

# OrganizationUserDataSourceFilters


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tags** | **Object** |          Tags to filter by. Supports logical AND and OR operations. Input should be like below:         {             \&quot;OR\&quot;: [                 {                 \&quot;key\&quot;: \&quot;subject\&quot;,                 \&quot;value\&quot;: \&quot;holy-bible\&quot;,                 \&quot;negate\&quot;: false                 },                 {                     \&quot;key\&quot;: \&quot;person-of-interest\&quot;,                     \&quot;value\&quot;: \&quot;jesus christ\&quot;,                     \&quot;negate\&quot;: false                 },                 {                     \&quot;key\&quot;: \&quot;genre\&quot;,                     \&quot;value\&quot;: \&quot;fiction\&quot;,                     \&quot;negate\&quot;: true                 }                 {                     \&quot;AND\&quot;: [                         {                             \&quot;key\&quot;: \&quot;subject\&quot;,                             \&quot;value\&quot;: \&quot;tao-te-ching\&quot;,                             \&quot;negate\&quot;: true                         },                         {                             \&quot;key\&quot;: \&quot;author\&quot;,                             \&quot;value\&quot;: \&quot;lao-tzu\&quot;,                             \&quot;negate\&quot;: false                         }                     ]                 }             ]         }         For a single filter, the filter block can be placed within either an \&quot;AND\&quot; or \&quot;OR\&quot; block.          |  [optional] |
|**source** | **DataSourceTypeNullable** |  |  [optional] |
|**ids** | **List&lt;Integer&gt;** |  |  [optional] |
|**revokedAccess** | **Boolean** |  |  [optional] |



