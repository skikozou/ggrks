# How to Use

## 1. Install
```
go install "github.com/skikozou/ggrks"
```

## 2. Initial settings
Please check [this](https://github.com/skikozou/ggrks?tab=readme-ov-file#initial-settings)

## 3. Fun!

# Initial settings

## 1. Create a Google Cloud Project
1. Go to the [Google Cloud Console](https://console.cloud.google.com/).
2. Log in with your Google account.
3. Click on the **"Select a project"** dropdown at the top and choose **"New Project"**.
4. Enter a **project name** (e.g., "Custom Search API Project") and click **"Create"**.

## 2. Enable the Custom Search JSON API
1. In the Google Cloud Console, go to **APIs & Services** > **Library**.
2. Search for **"Custom Search JSON API"** in the library search bar.
3. Click on the API and then click **"Enable"**.

## 3. Obtain Your API Key
1. Navigate to **APIs & Services** > **Credentials** in the Google Cloud Console.
2. Click **"Create Credentials"** > **"API Key"**.
3. Google will generate your API Key. Copy and save it securely; you’ll need it to authenticate your requests.

## 4. Create a Custom Search Engine
1. Go to the [Google Programmable Search Engine](https://programmablesearchengine.google.com/about/).
2. Click **"Get started"** and log in if necessary.
3. Click **"Create"**, leaving the site URL field blank to allow searches across the entire web.

## 5. Retrieve Your Search Engine ID
1. Once your Custom Search Engine is created, click on its name in the [Custom Search Engine Dashboard](https://programmablesearchengine.google.com/cse/all).
2. In the **"Basics"** tab, you’ll see the **Search engine ID**.
3. Copy this ID, as it is required to use the API for making queries.

## 6. Write config
1. Start your console.
2. Run "ggrks".
3. Write API keys.