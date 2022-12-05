# Moonlight
A new way to write native applications

## Get Tokens from a TypeScript file
To create a Token list of a  TypeScript (*.ts) file, you can use the next command in your terminal:

```
moonlight get-tokens FilePath
```

Or if you want to get tokens from multiple files:

```
moonlight get-tokens FilesPath...
```

It will return a JSON object as string:

```
// The output will be like this:
//Note: The types discribed here are JS Types but in golang: [String] is of type string and [Number] is of type int

{ FilePath [String] : 
 { FileNumber [String] :
   { LineNumber [String] : [
     { "Type": [String],
        "Value":[String | Number],
        "Loc":{
          "Start":{
            "Line": [Number],
            "Column": [Number],
           },
           "End":{
            "Line": [Number],
            "Column": [Number],
           },
         },
       },
     ],
   },
 },
}

```
