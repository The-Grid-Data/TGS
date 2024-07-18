# Profile Lens

The Profile lens captures comprehensive information about an entity’s offerings and fundamental characteristics. This includes:

- Services or products offered
- Business activities
- Supported currencies
- Real-world identifiers

## Parameters

### Name
- **Parameter ID:** `profiles.name`
- **Type:** Short Text
- **Description:** The full official name of the Profile, as used in their branding, documentation, and across online platforms.
- **Validation Steps:**
  1. Check the "About" or "Company" page on the Profile's official website.
  2. Look for consistency in the name used across their website, social media profiles, and any legal documentation.
  3. If there are variations of the name (e.g. full name vs. acronym), use the most frequently used full name.
- **In DBD:** true
- **Notes:** We use the term “Profile” because not every entity will be or have a project, product, or organization. Those features of a Profile will be defined in other fields.

### Type
- **Parameter ID:** `profiles.profileType`
- **Type:** Link to another Table
- **Specifications:** Profile - Sector / Types
- **Description:** The primary operational category or business model of the Profile (e.g. DAO, project, investor, government). Refer to standardized list of Profile types.
- **Validation Steps:**
  1. Review the Profile's website and documentation to identify how they describe their operational structure and purpose.
  2. Check if the Profile fits into one of the predefined categories in the standardized Profile types list.
  3. If the Profile spans multiple categories, select the one that best represents their core operation or purpose.
- **In DBD:** true
- **Notes:** Acts as a high-level overview of the Profile, for example, investor.

### Sector
- **Parameter ID:** `profiles.profileSector`
- **Type:** Link to another Table
- **Specifications:** Profile - Sector / Types
- **Description:** The primary industry, use case, or problem space the Profile operates in or serves. Refer to standardized list of sectors and subsectors.
- **Validation Steps:**
  1. Review the Profile's website, documentation, and marketing content to identify the key industry or use case they target.
  2. Check which standardized sector and subsector best align with the Profile's focus area.
  3. Consider the types of events, conferences, or communities the Profile engages with to validate the sector choice.
- **In DBD:** true
- **Notes:** Could be use case VS industry. Example: is it an ecosystem surrounding an original smart contract, or is it a proper blockchain where participants can build upon it?

### Current Status
- **Parameter ID:** `profiles.profileStatus`
- **Type:** Link to another Table
- **Specifications:** Profile - Sector / Types
- **Description:** The current operational state of the Profile (e.g. Active, Inactive, Closed). Refer to the standardized list of status options.
- **Validation Steps:**
  1. Check the Profile's website and social media for any announcements regarding their operational status and if they have a live product.
  2. Attempt to interact with the Profile's products or services to confirm if they are functional and actively maintained.
  3. Ensure the selected status aligns with the standardized list of status options.
- **In DBD:** true

### Founding Date
- **Parameter ID:** `profiles.foundingDate`
- **Type:** Date
- **Description:** The date when the Profile was first established, registered, or publicly announced. Use ISO 8601 format (YYYY-MM-DD) when possible.
- **Validation Steps:**
  1. Check the Profile's website, about page, or documentation for information on when they were founded or launched.
  2. Look for official announcements, press releases, or blog posts that mention the Profile's founding date.
  3. If only a month and year are available, use the first day of that month (e.g. 2022-01-01 for January 2022).
- **In DBD:** true

### Logo
- **Parameter ID:** `profiles.logo`
- **Type:** Image
- **Specifications:** URL
- **Description:** A square PNG or SVG image file representing the Profile's logo, with a transparent background and a minimum resolution of 256x256 pixels.
- **Validation Steps:**
  1. Check the Profile's website, media kit, or brand guidelines for its official logo.
  2. Confirm the image is in PNG or SVG format, has a transparent background, and is square in shape.
  3. Verify the image has a minimum resolution of 256x256 pixels and is visually clear and recognizable.
- **In DBD:** true
- **Notes:** Consider the creation of 2 logos, one the logo without the name (icon) and one with the name, full logo.

### Short Description
- **Parameter ID:** `profiles.descriptionShort`
- **Type:** Long text
- **Description:** A concise, one-sentence summary of the Profile's primary purpose, product, or service. Limit to 140 characters and avoid subjective claims.
- **Validation Steps:**
  1. Ensure the description clearly and objectively states the Profile's main function or offering.
  2. Confirm the description is written in the third person and does not include subjective claims or opinions.
  3. Verify the description is within the 140-character limit.
- **In DBD:** true
- **Notes:** Why so short: The shorter the description, the less we have to validate.

### Long Description
- **Parameter ID:** `profiles.descriptionLong`
- **Type:** Long text
- **Description:** A more detailed overview of the Profile, including its mission, target audience, and key features or offerings. Limit to 500 characters and avoid subjective claims.
- **Validation Steps:**
  1. Ensure the description provides a clear and objective overview of the Profile's purpose and offerings.
  2. Confirm the description is written in the third person and does not include subjective claims or opinions.
  3. Verify the description is within the 500-character limit.
- **In DBD:** true
- **Notes:** Why so short, even as a “long” description: The shorter the description, the less we have to validate.

### Tag Line
- **Parameter ID:** `profiles.tagLine`
- **Type:** Long text
- **Description:** A brief, memorable phrase or slogan used by the Profile to convey its mission, value proposition, or unique selling point. Limit to 100 characters.
- **Validation Steps:**
  1. Check the Profile's website, marketing materials, or social media accounts for a consistently used tagline or slogan.
  2. Confirm the tagline is concise, memorable, and effectively communicates the Profile's key message.
  3. Verify the tagline is within the 100-character limit.
- **In DBD:** true

### Main URL
- **Parameter ID:** `profiles.urlMain`
- **Type:** URL
- **Description:** The primary official website URL for the Profile. Use the root domain without "www" or "https://".
- **Validation Steps:**
  1. Check the Profile's social media accounts, press releases, and other official sources to confirm their main website URL.
  2. Ensure the URL is the root domain and does not include "www" or "https://".
  3. Validate that the URL leads to an active website related to the Profile.
- **In DBD:** true

### Blog URL
- **Parameter ID:** `profiles.urlBlog`
- **Type:** URL
- **Description:** The URL of the Profile's primary blog or news section. Use the root domain without "www" or "https://".
- **Validation Steps:**
  1. Check the Profile's website for a dedicated blog, news, or insights section.
  2. Confirm the URL is the root domain and does not include "www" or "https://".
  3. Validate that the URL leads to an active blog or news page related to the Profile.
- **In DBD:** true
- **Notes:** The purpose of this field is for people to learn more about the content that a Profile is creating. Some companies will have a blog page on their website, but they may publish more frequently somewhere else for distribution purposes. Many companies also have a “news” page, which consolidates content such as articles, podcasts, press releases, and product announcements.

### Documentation URL
- **Parameter ID:** `profiles.urlDocumentation`
- **Type:** URL
- **Description:** The URL of the Profile's primary technical documentation or knowledge base. Use the root domain without "www" or "https://".
- **Validation Steps:**
  1. Check the Profile's website for a dedicated documentation, guides, or support section.
  2. Confirm the URL is the root domain and does not include "www" or "https://".
  3. Validate that the URL leads to an active documentation page related to the Profile's products or services.
- **In DBD:** true

### Whitepaper URL
- **Parameter ID:** `profiles.urlWhitepaper`
- **Type:** URL
- **Description:** The URL of the Profile's most recent whitepaper or technical document outlining its technology, roadmap, and token economics. Use the root domain without "www" or "https://".
- **Validation Steps:**
  1. Check the Profile's website, blog, or documentation for links to its whitepaper or technical documents.
  2. Confirm the URL is the root domain and does not include "www" or "https://".
  3. Validate that the URL leads to a downloadable or web-hosted document in a standard format (e.g., PDF, HTML).
- **In DBD:** true
- **Notes:** This could be a copy that we host? It may also be versioned!

### Slug
- **Parameter ID:** `profiles.slug`
- **Type:** Short Text
- **In DBD:** false
