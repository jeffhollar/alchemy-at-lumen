# OpenAI Text Transformation Demo

This project includes two LangChain application which demonstrate different approaches to using OpenAI's language models for text transformation tasks. It showcases both direct OpenAI API usage and LangChain integration for handling prompts and responses.

A Streamlit-based web application has been included which helps businesses translate and transform customer reviews into polite, professional messages in different languages.

## Features [ intro-chat-prompts.py ]

- Text translation and tone modification using OpenAI's GPT models
- Two different implementation approaches:
  1. Direct OpenAI API calls
  2. LangChain framework with prompt templates
- Environment variable management for API key security

## Features [ streamlit_app.py ]

- **Polite Tone Transformation**: Automatically rewrites customer reviews in a more professional and respectful tone
- **Multi-language Support**: Translates reviews into any target language
- **Customizable Translation**: Allows users to specify:
  - Desired tone of the translated message
  - Target language
  - Company name for context
- **Secure API Key Management**: Safely handles OpenAI API keys through environment variables
- **Model Selection**: Choose between different GPT models for translation
- **User-friendly Interface**: Simple and intuitive web interface built with Streamlit

## Create an OpenAI API Key 
To create an OpenAI API key for accessing the GPT-4 model, follow these steps:

1. Create an OpenAI Account: Visit the OpenAI website (https://openai.com) and sign up for an account if you don’t already have one.
1. Log In: Once you have an account, log in to your OpenAI account.
1. Access API Settings: Navigate to the API section. This is usually found in the dashboard under “API” or something similar.
1. Create a New API Key: Look for an option to create a new API key. This might be labeled as “Create API Key” or something similar.
1. Save the Key: Once the API key is generated, be sure to copy and save it in a secure location. You won’t be able to view this key again for security reasons.
1. Set Up Billing: Make sure you have set up payment methods, as using the API usually involves certain costs depending on usage.
1. Start Using the API: With your API key, you can now authenticate requests to access the GPT-4 model.

Always make sure to keep your API key secure and never expose it in client-side code or public repositories.

## Requirements

- Python 3.x
- OpenAI API key
- Required Python packages (listed in requirements.txt)

## Setup

1. Clone this repository
   ```
   git clone <repository-url>
   cd <repository-directory>
   ```
2. Install the required packages:
   ```bash
   pip install -r requirements.txt
   ```
3. Create a `.env` file in the project root and add your OpenAI API key:
   ```
   OPENAI_API_KEY=your_api_key_here
   ```

## Running Application [ intro-chat-prompts.py ]
To run the python application, use the following command:

```
python intro-chat-prompts.py
```

## Running Application [ streamlit_app.py ]
To run the Streamlit application, use the following command:

```
streamlit run streamlit_app.py
```

This will start a local web server and open the application in your default web browser.

### Usage

1. Run the application:
```bash
streamlit run streamlit_app.py
```

2. Open your web browser and navigate to the provided local URL (typically http://localhost:8501)

3. In the sidebar:
   - Enter your OpenAI API key
   - Select your preferred GPT model

4. In the main interface:
   - Enter or paste the customer review
   - Specify the desired tone
   - Choose the target language
   - Enter your company name
   - Click "Translate" to process the review

## Notes

- The application uses GPT-4 by default
- Temperature settings are adjusted for different use cases (deterministic vs. creative responses)
- Environment variables are used for secure API key management

## Alert 

The code presented in this project should be considered as demonstrative and not used in as production quality.

## License

MIT License