# Wikipedia Research Tool

This project provides two implementations of a Wikipedia research tool that leverages LangChain, OpenAI, and Wikipedia to answer questions and perform calculations.

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

## Using the OpenAI API Key
For this application, edit the apikey.py file. Replace <OPENAI_KEY_VALUE> with your new OpenAI API Key value.

```
apikey = "<OPENAI_KEY_VALUE>"
```


## Applications

### 1. Command Line Interface (agents.py)
The `agents.py` file provides a command-line interface for the Wikipedia research tool. It uses LangChain agents to process user queries, search Wikipedia, and perform calculations.

**Purpose:**
- Demonstrates the core functionality of the Wikipedia research tool
- Provides a simple command-line interface for research queries
- Useful for testing and development purposes

**How to Run:**

```bash
cd wikipediatool
pip3 install -r requirements.txt
python agents.py
```

### 2. Web Interface (app.py)
The `app.py` file provides a Streamlit-based web interface for the Wikipedia research tool. It offers a user-friendly way to interact with the research capabilities.

**Purpose:**
- Provides a modern, user-friendly web interface
- Makes the tool accessible to non-technical users
- Includes helpful tips and error handling
- Offers a more interactive experience

**How to Run:**

```bash
cd wikipediatool
pip3 install -r requirements.txt
streamlit run app.py
```

## Features
- Wikipedia search integration
- Mathematical calculations
- Natural language processing
- Error handling
- User-friendly interface (app.py)

## Project Setup
```bash
git clone https://github.com/jeffhollar/alchemy-at-lumen.git
cd wikipediatool
```

## Run Applications

### CLI Agents Application

```
python3 agents.py
```
***TEST PROMPT***

```
What year was Obama elected as the US President?
```

### WebUI Application

```
streamlit run app.py
```
***TEST PROMPT***

```
Can you describe Quantum Computing?
```

## Requirements
All required dependencies are listed in `requirements.txt`. The main dependencies include:
- streamlit
- langchain-community
- openai
- wikipedia
- python-dotenv
- pydantic 

## Alert 

The code presented in this project should be considered as demonstrative and not used in as production quality.

## License

MIT License