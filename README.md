# tranform
transform loaded oai with AI


To convert OAI xml to CSV, Zek is required
* https://github.com/miku/zek


Run Zek on a metha- OAI source to generate the golang datastructure

Save OAI to a single file using `metha-cat`

```
metha-cat -$URL -set $SET_NAME > input_xml/$OAI_CODE.xml
```

Run Zek on the file to generate the Go data structure- used for custom conversion per oai set

```
zek input_xml/$OAI_CODE.xml
```
