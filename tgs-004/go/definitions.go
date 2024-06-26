package tgs_004

var TGS_004 = map[string]ParamDefinition{
	`profiles.id`: {
		ID:                   `profiles.id`,
		Description:          `Unique identifier for the Profile, auto-increment`,
		Validation:           ``,
		IsLinkToAnotherTable: true,
	},
	`profiles.name`: {
		ID:                   `profiles.name`,
		Description:          `Full name of the Profile`,
		Validation:           `The full name of the Profile, as is used in their logo, across social media and documentation.`,
		IsLinkToAnotherTable: true,
	},
	`profiles.profileType`: {
		ID:                   `profiles.profileType`,
		Description:          `What is the type of this Profile? Often things like DAO, project, investor, government or other. `,
		Validation:           `What would a user most likely refer to this Profile as. `,
		IsLinkToAnotherTable: true,
	},
	`profiles.profileSector`: {
		ID:                   `profiles.profileSector`,
		Description:          `The main industry or client type that this Profile serves. Acts as a high level category of a Profile. `,
		Validation:           `Sector may be determined by the type of conferences or events the Profile attends. Sector may also be determined by the regulations the Profile is subject to follow.`,
		IsLinkToAnotherTable: true,
	},
	`profiles.profileSubsector`: {
		ID:                   `profiles.profileSubsector`,
		Description:          `Depending on which Sector is selected, the Sub Sector provides further context as to the service it provides, the industry it serves, or how the Profile interacts with its customers.`,
		Validation:           `Information about the Sub Sector may be determined by the services page on the Profile’s website.`,
		IsLinkToAnotherTable: true,
	},
	`profiles.profileStatus`: {
		ID:                   `profiles.profileStatus`,
		Description:          `What is the Profile’s current stage?`,
		Validation:           `If a person can only interact one way with the profile (e.g. sign up for newsletter, follow on Twitter), then the profile is Announced. If a person can successfully have a two-way interaction with the profile’s offering (e.g. use the product or service, build on the blockchain), then the profile is Active. If a person has frequent unsuccessful interactions with the profile’s offering (e.g. product does not function as intended, the platform’s data is out of date), then the profile is Inactive. If a person can no longer interact with the profile’s offering, or if the entity has announced the profile’s closure, then the profile is Closed.`,
		IsLinkToAnotherTable: true,
	},
	`profiles.foundingDate`: {
		ID:                   `profiles.foundingDate`,
		Description:          `Date when Profile publicised itself or registered with a regulatory body. Can usually be shortened to Month and Year. `,
		Validation:           `Making the best effort to figure out when this Profile launched, can often be found on their social media bio (e.g. when they setup the account) or in their first white paper. The Founding Date may also be when the Profile published their first piece of content (e.g. their website, their first Tweet, their first LinkedIn post). Always put in a date where possible but month and year is acceptable.`,
		IsLinkToAnotherTable: true,
	},
	`profiles.logo`: {
		ID:                   `profiles.logo`,
		Description:          `URL to Standard size image of the visual identity of the Profile with a transparent background, use PNG or SVG`,
		Validation:           `This image should be recognizable as the profile’s visual identity in both color and black-and-white. For example, if the Profile modifies their Logo to celebrate holidays or special occasions, a person should still be able to identify the visual identity of the brand within the graphic.`,
		IsLinkToAnotherTable: true,
	},
	`profiles.descriptionShort`: {
		ID:          `profiles.descriptionShort`,
		Description: `Explains in one sentence (no more than 140 characters) what this Profile does. It may also describe whom the Profile serves, and how the Profile is used.`,
		Validation: `Description is written in third person singular. Description must answer what the Profile is in a complete sentence. Description may answer whom the Profile serves and how the Profile is used. Description cannot have any content that is subject to change frequently or in a timely manner. Description cannot make any claims subject to opinion. Does not include links. Does not define common industry terms within the description.
Uniswap
Yes: Uniswap is a DeFi protocol centered around its DEX that sources liquidity for swapping crypto assets. (101 characters. This description explains what the Profile is in objective terms.)
No: Uniswap is the largest decentralized exchange (or DEX) operating on the Ethereum blockchain. (”largest” is subjective. DEX does not need to be defined.)

Solana
Yes: Solana is a L1 blockchain that supports smart contracts and facilitates the creation of dApps. (94 characters. This description explains what Solana is and how people interact with the profile.)
No: Solana is a Layer 1 blockchain that offers users fast speeds and affordable costs. (”fast speeds and affordable costs” are subjective. They are considered perks or advantages of using the Profile’s offering, and not an explanation of what Solana is, who the Profile serves, or how the Profile is used.)`,
		IsLinkToAnotherTable: true,
	},
	`profiles.descriptionLong`: {
		ID:                   `profiles.descriptionLong`,
		Description:          `In 3 to 5 sentences, briefly expands on what this profile does and for whom, and can also briefly explain why. It does not need to explain how.`,
		Validation:           `Description is written in third person singular. Description may give more detail about what the profile is. Description may give more detail about whom the profile serves. Description may include a mission statement or goals. Description cannot have any content that is subject to change frequently or in a timely manner. Description cannot make any claims subject to opinion. Does not include links or calls to action.`,
		IsLinkToAnotherTable: true,
	},
	`profiles.tagLine`: {
		ID:                   `profiles.tagLine`,
		Description:          `A slogan or phrase that the Profile commonly uses in marketing materials.`,
		Validation:           `This text may be found in the Profile’s social media account descriptions and on the Profile’s website.`,
		IsLinkToAnotherTable: false,
	},
	`profiles.urlMain`: {
		ID:                   `profiles.urlMain`,
		Description:          `Links to the main website of the Profile. At the moment we use only domain name (like bitcoin.org) for technical reasons.`,
		Validation:           `Check social media and Google to confirm that this the main URL of the profile. Confirm that it is not a ‘similar’ URL. Domain name (without https or www). Must be unique. `,
		IsLinkToAnotherTable: true,
	},
	`profiles.urlBlog`: {
		ID:                   `profiles.urlBlog`,
		Description:          `Link to the root of the Profile’s page for content`,
		Validation:           `The URL is for the website with the most published content that does not focus on the Profile’s services or operations. The text or content on the website will often have opinions and announcements concerning the Profile, its Sector, and its Subsector. If the Profile appears to have more than one blog (e.g. a news subpage on their website and a Medium page), choose the blog that has the most recently published content.`,
		IsLinkToAnotherTable: true,
	},
	`profiles.urlDocumentation`: {
		ID:                   `profiles.urlDocumentation`,
		Description:          `Link to the root of the documentation of a profile`,
		Validation:           `The URL will likely have “docs” or “documentation” in the address. The site will likely have sections with titles such as “Getting Started”, “Governance”, “Wallets”, “Transactions”, “Explorer”, “Glossary” and “Terminology”. The site may also have assets and sources from gitbook.com . `,
		IsLinkToAnotherTable: true,
	},
	`profiles.urlWhitepaper`: {
		ID:                   `profiles.urlWhitepaper`,
		Description:          `Link to the most recent copy of the whitepaper`,
		Validation:           `The whitepaper is often a PDF, but it may instead be in Notion or Gitbook. It provides written details of the problem, solution, and possible technical architecture that the Profile provides.`,
		IsLinkToAnotherTable: true,
	},
	`products.id`: {
		ID:                   `products.id`,
		Description:          `Unique identifier, int64, auto-increment`,
		Validation:           ``,
		IsLinkToAnotherTable: true,
	},
	`products.name`: {
		ID:                   `products.name`,
		Description:          `What this service or offering is called`,
		Validation:           `A Product is something that a customer of the Profile can interact with (e.g. an app). It is not an Asset.`,
		IsLinkToAnotherTable: true,
	},
	`products.productType`: {
		ID:                   `products.productType`,
		Description:          `What kind of service or offering is this? For example is this an exchange or is it a wallet? `,
		Validation:           `Type can be determined by how customers and other ecosystem actors interact with the Product.`,
		IsLinkToAnotherTable: true,
	},
	`products.descriptionShort`: {
		ID:                   `products.descriptionShort`,
		Description:          `Explains what the service or offering does (no more than 140 characters)`,
		Validation:           `Description is written in third person. Text includes information about what the Product is. Text may include how the service or offering is used. Description does not include opinions or subjective information.`,
		IsLinkToAnotherTable: true,
	},
	`products.urlToProduct`: {
		ID:                   `products.urlToProduct`,
		Description:          `Direct URL to where the service or offering can be accessed, either subdomain or folder`,
		Validation:           `Domain name (without https or www). Must be unique. Is not the same as the Profile URL, but may be a subpage or subdomain of it. Should not be a link to a third party app store (e.g. Google Play).`,
		IsLinkToAnotherTable: true,
	},
	`products.productStatus`: {
		ID:                   `products.productStatus`,
		Description:          `Describes who has permission to interact with the service or offering.`,
		Validation:           `If the Product is not accessible to any user, it is in Development. If the Profile or Entity is strictly limiting the number or type of people who may use the Product, it is in Early Access. If the Profile has publicly invited anyone to use the Product under the disclaimer that it is still in Development, it is in Open Beta. If the Product can be used by anyone, it is Live. The Product may also be considered Live if it has a price or fee for usage, or if it is available through third party platforms. A Product may be considered in Testnet if the Profile announces that it is in Testnet. It may also be considered in Testnet if the Product is a blockchain. A Product may be considered in Mainnet if the Profile announces that it is in Mainnet. It may also be considered in Mainnet if the Product is a blockchain. `,
		IsLinkToAnotherTable: true,
	},
	`products.productAddress`: {
		ID:                   `products.productAddress`,
		Description:          `What is the main smart contract address of this product, if any?`,
		Validation:           `This text is most likely a sequence of numbers and letters, with no spaces in between.`,
		IsLinkToAnotherTable: true,
	},
	`products.isMainProduct`: {
		ID:                   `products.isMainProduct`,
		Description:          `Is this the primary service or offering of the Profile that it is a part of?`,
		Validation:           `This Product may generate the majority of revenue for the Profile. It may have the most DAU or downloads of the Profile’s suite of offerings.`,
		IsLinkToAnotherTable: true,
	},
	`products.launchDate`: {
		ID:                   `products.launchDate`,
		Description:          `When this Product was originally released`,
		Validation:           `The date may be found on the Profile’s social media posts, or in press releases and announcements on the Blog. Must include year, and preferably the month. May include the day.`,
		IsLinkToAnotherTable: true,
	},
	`products.deployedOn`: {
		ID:                   `products.deployedOn`,
		Description:          `On which blockchain does this Product operate, if applicable? This is not the hosting provider (e.g. AWS, Google Cloud) `,
		Validation:           `This information can be confirmed with the Address. This information may also be confirmed with the Product’s documentation.`,
		IsLinkToAnotherTable: true,
	},
	`products.assetSupport`: {
		ID:                   `products.assetSupport`,
		Description:          `Which stores of value does this product support, if applicable? It might be general support or specific to that application. `,
		Validation:           `This information is often stored within the Product itself, but it may also be found on the Product’s website. It may be called Coin Support or Token Support.`,
		IsLinkToAnotherTable: true,
	},
	`products.profileId`: {
		ID:                   `products.profileId`,
		Description:          `What Profile this product is associated with.`,
		Validation:           ``,
		IsLinkToAnotherTable: false,
	},
	`entities.id`: {
		ID:                   `entities.id`,
		Description:          `Unique identifier, int64, auto-increment`,
		Validation:           ``,
		IsLinkToAnotherTable: true,
	},
	`entities.name`: {
		ID:                   `entities.name`,
		Description:          `Name that is registered with the regulatory or governmental body`,
		Validation:           `This Name can be confirmed with the regulatory body in the jurisdiction where the Entity is registered. It may also be confirmed in the Terms and Conditions on the Entity’s website.`,
		IsLinkToAnotherTable: true,
	},
	`entities.tradeName`: {
		ID:                   `entities.tradeName`,
		Description:          `Other names this Entity commonly goes by. Also called Doing Business As (DBA)`,
		Validation:           `This name may be frequently used on the same website where the Terms and Conditions can be found. This name may also match the name of the linked Profile.`,
		IsLinkToAnotherTable: true,
	},
	`entities.entityType`: {
		ID:                   `entities.entityType`,
		Description:          `How is this Entity registered with the regulatory or governmental body?`,
		Validation:           `This information may be found with the Entity’s regulatory body, such as the Chamber of Commerce or Secretary of State. This information may also be found in the Entity’s or Profile’s Terms and Conditions.`,
		IsLinkToAnotherTable: true,
	},
	`entities.dateOfIncorporation`: {
		ID:                   `entities.dateOfIncorporation`,
		Description:          `When this entity was started, in a standard date format`,
		Validation:           `This information may be found with the Entity’s regulatory body, such as the Chamber of Commerce or Secretary of State. This information may also be found in the Entity’s or Profile’s Terms and Conditions.`,
		IsLinkToAnotherTable: true,
	},
	`entities.urlToEntity`: {
		ID:                   `entities.urlToEntity`,
		Description:          `Link to the website for this Entity. Sometimes might be different from main URL for the Profile`,
		Validation:           `Domain name (without https or www). Must be unique. This website may have more information about the Entity’s operations and less marketing material for their target audience.`,
		IsLinkToAnotherTable: true,
	},
	`entities.profile`: {
		ID:                   `entities.profile`,
		Description:          `Link to the Profile this Entity is associated with. For example - For Uniswap, Universal Navigation Inc the Entity is associated with Uniswap the Profile`,
		Validation:           ``,
		IsLinkToAnotherTable: true,
	},
	`entities.parentEntity`: {
		ID:                   `entities.parentEntity`,
		Description:          `Name of the business, organisation, or body that owns or manages this Entity (if applicable).`,
		Validation:           `This information may be found with the Entity’s regulatory body, such as the Chamber of Commerce or Secretary of State. This information may also be found in the Entity’s or Profile’s Terms and Conditions.`,
		IsLinkToAnotherTable: true,
	},
	`entities.localRegistrationNumber`: {
		ID:                   `entities.localRegistrationNumber`,
		Description:          `The unique identifier assigned to the Entity by the regulatory or government body`,
		Validation:           `This information may be found with the Entity’s regulatory body, such as the Chamber of Commerce or Secretary of State. This information may also be found in the Entity’s or Profile’s Terms and Conditions.`,
		IsLinkToAnotherTable: true,
	},
	`entities.taxIdentificationNumber`: {
		ID:                   `entities.taxIdentificationNumber`,
		Description:          `The unique identifier assigned to the Entity by the tax authority`,
		Validation:           `This information may be found with the tax authority where the Entity is registered (e.g. Belastingdienst, IRS). This information may also be found in the Entity’s or Profile’s Terms and Conditions.`,
		IsLinkToAnotherTable: true,
	},
	`entities.address`: {
		ID:                   `entities.address`,
		Description:          `Physical location where the Entity is registered`,
		Validation:           `The address includes the street name and street number. The address may include additional information such as the unit, suite, or floor. The address includes the city and postal code. The address may include the state or province if applicable.`,
		IsLinkToAnotherTable: true,
	},
	`entities.country`: {
		ID:                   `entities.country`,
		Description:          `Name of the Jurisdiction the Entity is registered in`,
		Validation:           `This information may be found with the Entity’s regulatory body, such as the Chamber of Commerce or Secretary of State. This information may also be found in the Entity’s or Profile’s Terms and Conditions.`,
		IsLinkToAnotherTable: true,
	},
	`entities.leiNumber`: {
		ID:                   `entities.leiNumber`,
		Description:          `What is the LEI code linked to their GLEIF profile`,
		Validation:           `This information can be found on the GLEIF website`,
		IsLinkToAnotherTable: true,
	},
	`assets.id`: {
		ID:                   `assets.id`,
		Description:          `Unique identifier, int64, auto-increment`,
		Validation:           ``,
		IsLinkToAnotherTable: true,
	},
	`assets.name`: {
		ID:                   `assets.name`,
		Description:          `What to call the item or token that represents value, and can be traded or sold`,
		Validation:           `The name may be different from the Ticker symbol (e.g. Ether is the name, ETH is the ticker).`,
		IsLinkToAnotherTable: true,
	},
	`assets.ticker`: {
		ID:                   `assets.ticker`,
		Description:          `How the Asset is identified and listed on exchanges`,
		Validation:           `Do not include any symbols. Only use numbers and capital letters.`,
		IsLinkToAnotherTable: true,
	},
	`assets.icon`: {
		ID:                   `assets.icon`,
		Description:          `Standard square Image of the Logo of the Token (if different from the Profile) with a transparent background, use PNG or SVG`,
		Validation:           ``,
		IsLinkToAnotherTable: true,
	},
	`assets.shortDescription`: {
		ID:                   `assets.shortDescription`,
		Description:          `Short explanation of what the Asset is, and possibly include the value it represents and how the Asset holder uses it.`,
		Validation:           `The description may include details about what the Asset represents (e.g. proof of membership, equity in a company, monetary value). It does not include price or details about where the Asset can be acquired (e.g. does not list exchanges or marketplaces).`,
		IsLinkToAnotherTable: false,
	},
	`assets.assetType`: {
		ID:                   `assets.assetType`,
		Description:          `What type of representation of value is this?`,
		Validation:           `If the asset is primarily used for payment or investment purposes, then it is a Currency. If the asset is a collection of unique tokens (usually representing a form of media such as art, images, music or writing), then it is an NFT. If it is a token required to make decisions about a protocol or within an entity, then it is a Governance asset. If the token is required to take action within a certain system (such as redeeming services or claiming benefits), then it is a Utility asset. If the asset is often pegged to a fiat currency or traditional commodity, then it is a Stablecoin. If the token is clearly marked to represent a fungible financial instrument (but not a fiat currency or traditional commodity), it is a Security. If the asset has little to no utility and has “fun” as part of its purpose, then it is a Meme.`,
		IsLinkToAnotherTable: true,
	},
	`assets.urlToAssetDocs`: {
		ID:                   `assets.urlToAssetDocs`,
		Description:          `What is the best official source for asset docs? `,
		Validation:           `Domain name (without https or www). Must be unique. `,
		IsLinkToAnotherTable: true,
	},
	`assets.address`: {
		ID:                   `assets.address`,
		Description:          `What is the smart contract address of the asset issuer? Aka where did this asset come from?`,
		Validation:           `This information should be found via block explorer and ideally is checked against the documentation from the project. Assets should only be deployed on one address per blockchain. `,
		IsLinkToAnotherTable: true,
	},
	`assets.assetStandard`: {
		ID:                   `assets.assetStandard`,
		Description:          `Does this token follow an existing standard like ERC721, BRC20, ERC4337`,
		Validation:           `Most commonly focused in the documentation of the project, it could also be figured out on chain by seeing which functions it has. `,
		IsLinkToAnotherTable: true,
	},
	`assets.deployedOn`: {
		ID:                   `assets.deployedOn`,
		Description:          `On which blockchain is the asset built?`,
		Validation:           `This information is often found in the Documentation (Docs) or the About section of the Asset or Profile.`,
		IsLinkToAnotherTable: true,
	},
	`assets.profileId`: {
		ID:                   `assets.profileId`,
		Description:          `What Profile this Asset is associated with.`,
		Validation:           ``,
		IsLinkToAnotherTable: false,
	},
	`socials.id`: {
		ID:                   `socials.id`,
		Description:          `Unique identifier, int64, auto-increment`,
		Validation:           ``,
		IsLinkToAnotherTable: true,
	},
	`socials.url`: {
		ID:                   `socials.url`,
		Description:          `What is the URL of this social media link? This is a confirmed link. `,
		Validation:           ``,
		IsLinkToAnotherTable: true,
	},
	`socials.socialType`: {
		ID:                   `socials.socialType`,
		Description:          `This is what type of social media link is this: Discord, Farcaster, X, Facebook, Bluesky, Lens, Linkedin`,
		Validation:           ``,
		IsLinkToAnotherTable: true,
	},
	`socials.name`: {
		ID:                   `socials.name`,
		Description:          `What is the name of the account in social media?`,
		Validation:           ``,
		IsLinkToAnotherTable: false,
	},
}
