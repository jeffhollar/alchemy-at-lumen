import streamlit as st
import os
from apikey import apikey
from langchain_community.llms import OpenAI
from langchain.agents import load_tools, initialize_agent, AgentType

# Set up the page configuration
st.set_page_config(
    page_title="Wikipedia Research Tool",
    page_icon="🔍",
    layout="centered"
)

# Set the OpenAI API key
os.environ["OPENAI_API_KEY"] = apikey

# Initialize the LLM and tools
@st.cache_resource
def get_agent():
    llm = OpenAI(temperature=0.0, max_tokens=50)
    tools = load_tools(['wikipedia', 'llm-math'], llm)
    agent = initialize_agent(
        tools,
        llm,
        agent=AgentType.ZERO_SHOT_REACT_DESCRIPTION,
        verbose=True
    )
    return agent

# Main app interface
st.title("🔍 Wikipedia Research Tool")
st.write("Ask any question and let the AI research it using Wikipedia and perform calculations!")

# User input
user_input = st.text_input(
    "Enter your research question:",
    placeholder="e.g., In what year did the Titanic sink? How many years has it been since?"
)

# Initialize the agent
agent = get_agent()

# Process the input when the user submits
if user_input:
    with st.spinner("Researching..."):
        try:
            response = agent.invoke(user_input)
            st.success("Research complete!")
            st.write("### Answer:")
            st.write(response['output'])
        except Exception as e:
            st.error(f"An error occurred: {str(e)}")

# Add some helpful information
st.markdown("""
### Tips:
- Ask questions that can be answered using Wikipedia and calculations
- You can ask about historical events, scientific facts, or any topic available on Wikipedia
- The tool can perform mathematical calculations based on the information it finds
""") 