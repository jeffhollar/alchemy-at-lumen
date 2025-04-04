"""
This application demonstrates different approaches to using OpenAI's language models for text transformation tasks.
It showcases both direct OpenAI API usage and LangChain integration for handling prompts and responses.

Key Features:
- Text translation and tone modification using OpenAI's GPT models
- Two different implementation approaches:
  1. Direct OpenAI API calls
  2. LangChain framework with prompt templates
- Environment variable management for API key security
"""

import os
import openai
from dotenv import find_dotenv, load_dotenv
from langchain_openai import ChatOpenAI
from langchain.prompts import ChatPromptTemplate

# Load environment variables from .env file
load_dotenv(find_dotenv())
openai.api_key = os.getenv("OPENAI_API_KEY")

# Define the language model to use
llm_model = "gpt-4o"

# OpenAI Completion Endpoint
# This function provides a simple interface to get completions from OpenAI's API
# Parameters:
#   prompt: The text prompt to send to the model
#   model: The specific OpenAI model to use (defaults to llm_model)
def get_completion(prompt, model=llm_model):
    messages = [{"role": "user", "content": prompt}]
    response = openai.chat.completions.create(
        model=model,
        messages=messages,
        temperature=0,  # Low temperature for more deterministic responses
    )
    return response.choices[0].message.content

# Example 1: Direct OpenAI API Usage
# This section demonstrates how to:
# 1. Modify the tone of a customer review
# 2. Translate the modified review into another language
customer_review = """
 Your product is terrible!  I don't know how 
 you were able to get this to the market.
 I don't want this! Actually no one should want this.
 Seriously!  Give me money now!
 
"""
tone = """ Proper British English in a nice, warm, respectful tone """
language = "Turkish"

# Create a prompt that combines tone modification and translation
prompt = f""" 
  Rewrite the following {customer_review} in {tone}, and then
  please translate the new review message into {language}.
"""

rewrite = get_completion(prompt=prompt)

# Example 2: Using LangChain Framework
# This section demonstrates:
# 1. Using LangChain's ChatOpenAI wrapper
# 2. Creating and using prompt templates
# 3. Handling structured prompts with variables
chat_model = ChatOpenAI(temperature=0.7, model=llm_model)  # Higher temperature for more creative responses

# Define a template for translation with placeholders
template_string = """ 
 Translate the following text {customer_review}
 into italiano in a polite tone.
 And the company name is {company_name}
"""

# Create a reusable prompt template
prompt_template = ChatPromptTemplate.from_template(template_string)

# Format the template with actual values
translation_message = prompt_template.format_messages(
    customer_review=customer_review, company_name="Google"
)

# Get the translation response
response = chat_model.invoke(translation_message)
print(response.content)
