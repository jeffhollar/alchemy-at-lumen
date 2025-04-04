import os
import streamlit as st
from dotenv import find_dotenv, load_dotenv
from langchain_openai import ChatOpenAI
from langchain.prompts import ChatPromptTemplate

# Load environment variables from .env file
# This allows secure storage of API keys and other sensitive information
load_dotenv(find_dotenv())
openai_api_key = os.getenv("OPENAI_API_KEY")

# Configure the Streamlit page settings
# This sets the title, icon, and layout of the web application
st.set_page_config(page_title="Text Translation App", page_icon="🌐", layout="centered")

# Main application title and description
st.title("🌐 Text Translation App")
st.markdown(
    """
This app translates customer reviews into different languages with a polite tone.
"""
)

# Sidebar configuration section
# This section handles user authentication and model selection
with st.sidebar:
    st.header("Configuration")
    # Input field for OpenAI API key with password masking for security
    api_key = st.text_input("OpenAI API Key", value=openai_api_key, type="password")
    if not api_key:
        st.warning("Please enter your OpenAI API key in the sidebar.")
        st.stop()

    # Model selection dropdown
    # Users can choose between different GPT models for translation
    model = st.selectbox("Select Model", ["gpt-4o", "gpt-4o-mini"], index=0)

# Main content section for translation settings
st.header("Translation Settings")

# Default customer review text
# This serves as an example for users to understand the input format
default_review = """
Your product is terrible! I don't know how 
you were able to get this to the market.
I don't want this! Actually no one should want this.
Seriously! Give me money now!
"""

# Input fields for translation parameters
# These fields collect all necessary information for the translation process
customer_review = st.text_area("Customer Review", value=default_review, height=150)
tone = st.text_input(
    "Tone", value="Proper British English in a nice, warm, respectful tone"
)
language = st.text_input("Target Language", value="Turkish")
company_name = st.text_input("Company Name", value="Google")

# Translation execution section
# This handles the actual translation process when the button is clicked
if st.button("Translate"):
    # Input validation
    if not customer_review or not language:
        st.error("Please provide both a customer review and target language.")
    else:
        with st.spinner("Translating..."):
            try:
                # Initialize the OpenAI chat model with specified parameters
                chat_model = ChatOpenAI(
                    temperature=0.7,  # Controls randomness in responses
                    model=model,      # Selected model from sidebar
                    openai_api_key=api_key
                )

                # Create a prompt template for the translation task
                # This template structures how the AI will process the input
                template_string = """
                Rewrite the following customer review in a {tone}, and then
                translate the new review message into {language}.
                
                Customer Review: {customer_review}
                Company Name: {company_name}
                """

                prompt_template = ChatPromptTemplate.from_template(template_string)

                # Format the prompt with user inputs
                translation_message = prompt_template.format_messages(
                    customer_review=customer_review,
                    tone=tone,
                    language=language,
                    company_name=company_name,
                )

                # Get the AI's response
                response = chat_model.invoke(translation_message)

                # Display the translation results
                st.success("Translation completed!")
                st.subheader("Translated Review")
                st.write(response.content)

            except Exception as e:
                st.error(f"An error occurred: {str(e)}")

# Footer section
st.markdown("---")
st.markdown("Built with Streamlit and OpenAI")
