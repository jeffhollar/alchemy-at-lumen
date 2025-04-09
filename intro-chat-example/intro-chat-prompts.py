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

def transform_text(text, target_language, tone="polite"):
    """
    Transform the given text into the specified language with the given tone.
    
    Args:
        text (str): The text to transform
        target_language (str): The language to translate the text into
        tone (str): The desired tone of the translation (default: "polite")
    
    Returns:
        str: The transformed text
    """
    prompt = f""" 
    Rewrite the following text in a {tone} tone, and then
    please translate the new text into {target_language}.
    
    Text to transform:
    {text}
    """
    
    return get_completion(prompt=prompt)

# Example usage
if __name__ == "__main__":
    # Example text to transform
    customer_review = """
    Your product is terrible!  I don't know how 
    you were able to get this to the market.
    I don't want this! Actually no one should want this.
    Seriously!  Give me money now!
    """
    
    # Transform the text to German
    transformed_text = transform_text(customer_review, "German")
    print("Transformed text:")
    print(transformed_text)
    
    # Example 2: Using LangChain Framework
    chat_model = ChatOpenAI(temperature=0.7, model=llm_model)
    
    # Define a template for translation with placeholders
    template_string = """ 
    Translate the following text {text}
    into {target_language} in a {tone} tone.
    And the company name is {company_name}
    """
    
    # Create a reusable prompt template
    prompt_template = ChatPromptTemplate.from_template(template_string)
    
    # Format the template with actual values
    translation_message = prompt_template.format_messages(
        text=customer_review,
        target_language="Italian",
        tone="polite",
        company_name="Google"
    )
    
    # Get the translation response
    response = chat_model.invoke(translation_message)
    print("\nTransformed text using LangChain:")
    print(response.content)
