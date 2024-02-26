## TGS V0

```
// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table profiles { //consider putting in orginisations (groups of people doing something)
  id varchar [primary key]
  type varchar 
  name varchar
  // profileTypeId varchar
  // mainProductType varchar
  status profile_status 
  logo varchar
  descriptionShort text
  descriptionLong longtext
  urlMain varchar
  urlDocumentation varchar 
  productId varchar [ref: > products.id]
  assetId varchar [ref: > assets.id]
  socials varchar [ref: > socials.id]
  
}

enum profile_status {
  Announced
  Active
  Inactive
  Closed
}

Table socials {
  id varchar
  url varchar 
  type socialType
}

enum socialType {
  twitter
  farcaster
  discord
}


Table assets {
  id varchar [primary key]
  assetTicker varchar [note:' This is the ticker, it COULD be without the $ as we can append on the front end']
  assetName varchar 
  assetTypeId varchar [ref: > assetType.id]
  urlToAssetDocs varchar
  address varchar [note: 'This is the smart contract address']
  assetStandard varchar [ref: > assetStandardSupport.id, note: 'Does this asset follow a token standard?']
  deployedOn varchar [ref: > products.id, note: 'Where is this asset deployed']
}

Table assetType [note: 'Acts like an ENUM but has alot of extra info']{
  id varchar [primary key]
  typeName varchar 
  description varchar [note: 'We may want to store this externally' ]
}

Table assetStandardSupport [note: 'Acts like an ENUM']{
  id varchar [primary key]
  standard varchar [note: 'EG ERC-20' ]
  description varchar [note: 'We may want to store this externally' ]
}

Table products {
  id varchar [primary key]
  productName varchar
  urlToProduct varchar
  descriptionShort text
  productAddress varchar
  productTypeId varchar [ref: > productTypes.id] 
  productStatusId varchar [ref: > productStatus.id]
  assetSupport varchar [ref: > assets.id]
  mainProduct bool
}

Table productTypes [note: 'Acts like an ENUM']{
  id varchar [primary key]
  typeName varchar 
  description varchar 
}

Table productStatus {
  id varchar [primary key]
  productStatus varchar
  description varchar [note: 'We may want to store this externally, but could be helpful to include as a source of trurth' ]
}

Table entities {
  id varchar [primary key]
  entityName varchar 
  entityTradeName varchar
  urlToEntity varchar
  profileID varchar [ref: < profiles.id]
  childEntities varchar [ref: < entities.id]
  localRegestrationNumber varchar
  country varchar 
  leiNumber varchar
}

```

## Profiles - Consider putting in organizations (groups of people doing something)
- **id** - Primary key.
- **type** - No specific description provided.
- **name** - No specific description provided.
- **status** - Enum for profile status (Announced, Active, Inactive, Closed).
- **logo** - No specific description provided.
- **descriptionShort** - No specific description provided.
- **descriptionLong** - No specific description provided.
- **urlMain** - No specific description provided.
- **urlDocumentation** - No specific description provided.
- **productId** - Foreign key reference to `products.id`.
- **assetId** - Foreign key reference to `assets.id`.
- **socials** - Foreign key reference to `socials.id`.

## Socials
- **id** - No specific description provided.
- **url** - No specific description provided.
- **type** - Enum for social type (twitter, farcaster, discord).

## Assets
- **id** - Primary key.
- **assetTicker** - This is the ticker, it could be without the $ as we can append on the front end.
- **assetName** - No specific description provided.
- **assetTypeId** - Foreign key reference to `assetType.id`.
- **urlToAssetDocs** - No specific description provided.
- **address** - This is the smart contract address.
- **assetStandard** - Does this asset follow a token standard? Foreign key reference to `assetStandardSupport.id`.
- **deployedOn** - Where is this asset deployed. Foreign key reference to `products.id`.

## AssetType - Acts like an ENUM but has a lot of extra info.
- **id** - Primary key.
- **typeName** - No specific description provided.
- **description** - We may want to store this externally.

## AssetStandardSupport - Acts like an ENUM.
- **id** - Primary key.
- **standard** - E.g., ERC-20.
- **description** - We may want to store this externally.

## Products
- **id** - Primary key.
- **productName** - No specific description provided.
- **urlToProduct** - No specific description provided.
- **descriptionShort** - No specific description provided.
- **productAddress** - No specific description provided.
- **productTypeId** - Foreign key reference to `productTypes.id`.
- **productStatusId** - Foreign key reference to `productStatus.id`.
- **assetSupport** - Foreign key reference to `assets.id`.
- **mainProduct** - Boolean indicating if this is the main product.

## ProductTypes - Acts like an ENUM.
- **id** - Primary key.
- **typeName** - No specific description provided.
- **description** - No specific description provided.

## ProductStatus
- **id** - Primary key.
- **productStatus** - No specific description provided.
- **description** - We may want to store this externally, but could be helpful to include as a source of truth.

## Entities
- **id** - Primary key.
- **entityName** - No specific description provided.
- **entityTradeName** - No specific description provided.
- **urlToEntity** - No specific description provided.
- **profileID** - Foreign key reference to `profiles.id`.
- **childEntities** - Foreign key reference to `entities.id` (self-reference for hierarchical structures).
- **localRegestrationNumber** - No specific description provided.
- **country** - No specific description provided.
- **leiNumber** - No specific description provided.
