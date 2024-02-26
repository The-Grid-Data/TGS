// Profiles represent organizations, groups, or projects within the system.
Table profiles {
  id varchar [primary key, note: 'Unique identifier for the profile']
  type varchar [note: 'Type of profile e.g., DAO, Company']
  name varchar [note: 'Full name of the project']
  status profile_status [note: 'Current status of the project e.g., Active, Closed']
  logo varchar [note: 'URL to the project logo, preferably in PNG or SVG format']
  descriptionShort text [note: 'Short description of the project']
  descriptionLong longtext [note: 'Detailed description of the project']
  urlMain varchar [note: 'Main URL to the project website, must be https']
  urlDocumentation varchar [note: 'URL to the project documentation, must be https']
  productId varchar [ref: > products.id, note: 'Associated product ID']
  assetId varchar [ref: > assets.id, note: 'Associated asset ID']
  socials varchar [note: 'JSON or array storing social media links']
  industry varchar [note: 'Main industry of the project e.g., Finance']
  useCase varchar [note: 'Specific use cases of the project, stored as JSON or array']
  foundingDate date [note: 'Date when the project was founded or went live']
  targetAudience varchar [note: 'Intended users of the project e.g., Retail Traders']
  blogLinks varchar [note: 'URL to the project’s blog, must be https']
}

enum profile_status {
  Announced
  Active
  Inactive
  Closed
}

// Socials for projects, including types and URLs.
Table socials {
  id varchar [primary key, note: 'Unique identifier for social media link']
  url varchar [note: 'URL to the social media page, must be https']
  type socialType [note: 'Type of social media e.g., twitter, discord']
}

enum socialType {
  twitter
  farcaster
  discord
}

// Assets represent digital or physical assets associated with profiles.
Table assets {
  id varchar [primary key, note: 'Unique identifier for the asset']
  assetTicker varchar [note: 'Ticker symbol of the asset']
  assetName varchar [note: 'Full name of the asset']
  assetTypeId varchar [ref: > assetType.id, note: 'Type of asset']
  urlToAssetDocs varchar [note: 'URL to asset documentation, must be https']
  address varchar [note: 'Smart contract address for digital assets']
  assetStandard varchar [ref: > assetStandardSupport.id, note: 'Token standard followed by the asset']
  deployedOn varchar [ref: > products.id, note: 'Product or platform where the asset is deployed']
}

// Types of assets, similar to an ENUM but with additional information.
Table assetType {
  id varchar [primary key, note: 'Unique identifier for the asset type']
  typeName varchar [note: 'Name of the asset type']
  description varchar [note: 'Description of the asset type']
}

// Support for different asset standards (e.g., ERC-20).
Table assetStandardSupport {
  id varchar [primary key, note: 'Unique identifier for the asset standard']
  standard varchar [note: 'Name of the standard e.g., ERC-20']
  description varchar [note: 'Description of the standard']
}

// Products associated with profiles, detailing their specifics.
Table products {
  id varchar [primary key, note: 'Unique identifier for the product']
  productName varchar [note: 'Name of the product']
  urlToProduct varchar [note: 'URL to the product, must be https']
  descriptionShort text [note: 'Short description of the product']
  productAddress varchar [note: 'Address or location of the product, for digital products this could be a smart contract address']
  productTypeId varchar [ref: > productTypes.id, note: 'Type of product']
  productStatusId varchar [ref: > productStatus.id, note: 'Current status of the product']
  assetSupport varchar [ref: > assets.id, note: 'Associated asset ID']
  mainProduct bool [note: 'Indicates if this is the main product of the profile']
}

// Types of products, providing a structured classification.
Table productTypes {
  id varchar [primary key, note: 'Unique identifier for the product type']
  typeName varchar [note: 'Name of the product type']
  description varchar [note: 'Description of the product type']
}

// Current status of products, useful for lifecycle management.
Table productStatus {
  id varchar [primary key, note: 'Unique identifier for the product status']
  productStatus varchar [note: 'Status of the product e.g., Active, In Development']
  description varchar [note: 'Description of the product status']
}

// Entities represent legal or organizational entities associated with profiles.
Table entities {
  id varchar [primary key, note: 'Unique identifier for the entity']
  entityName varchar [note: 'Legal name of the entity']
  entityTradeName varchar [note: 'Trade name of the entity, if different from legal name']
  urlToEntity varchar [note: 'URL to the entity’s website, if available']
  profileID varchar [ref: < profiles.id, note: 'Associated profile ID']
  childEntities varchar [ref: < entities.id, note: 'Child entities, if any']
  localRegistrationNumber varchar [note: 'Local registration number of the entity']
  country varchar [note: 'Country where the entity is registered']
  leiNumber varchar [note: 'Legal Entity Identifier (LEI) number']
}

// Add tables for Funding, Regulatory, Token, Integrations, and Partnerships as needed based on CSV structure.
