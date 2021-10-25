pub contract NFTContract {

    pub var issuedSchemaId : UInt64 
    pub var ownedSchema : {UInt64 : Schema}
    pub var totalSupply: UInt64

    //enum for schema 
    pub enum schemaType: UInt8 {
        pub case String
        pub case Int
        pub case Fix64
        pub case Number
        pub case Bool
        pub case Address
    }

    // Structure of Schema
    pub struct Schema{
        pub let schemaName: String
        pub let schemaId: UInt64
        pub let format: [{String:schemaType}]

        init(name : String , format:[{String:schemaType}]){
            pre{
            name.length> 0 : "Could not create schema: name is required"
            }
        self.schemaName = name      
        self.schemaId = NFTContract.issuedSchemaId
        self.format = format
        }
    } 
  pub resource interface Minter {
        pub fun createSchema(name : String , format: [{String:schemaType}])

  }

  pub resource TemplateCollection: Minter {
            pub fun createSchema(name : String , format: [{String:schemaType}]){
           pre{
                NFTContract.ownedSchema.containsKey(NFTContract.issuedSchemaId) != true : "Schema Id is already exits"
            }
        
            for data in format {
                for values in data.values {
            // value should be string or integer or float or url
                if values == schemaType.String || values == schemaType.Int || values == schemaType.Fix64 || values == schemaType.Number ||  values == schemaType.Bool ||  values == schemaType.Address   {
                    continue
                } else {
                    panic("Must contain valid DataType")
                } 
                }
            }

            var newSchema = Schema(name:name, format : format)
            NFTContract.ownedSchema[NFTContract.issuedSchemaId] = newSchema
            NFTContract.issuedSchemaId =   NFTContract.issuedSchemaId + (1 as UInt64)
        }

  }

  pub fun createEmptyCollection(): @TemplateCollection{
        return <- create TemplateCollection()
    }

  init() {
        self.issuedSchemaId = 0
        self.totalSupply = 0
        self.ownedSchema = {}

        self.account.save(<- self.createEmptyCollection(),to:/storage/Collectionv1)
        self.account.link<&{Minter}>(/public/CollectionReceiver, target: /storage/Collectionv1)
    }
}