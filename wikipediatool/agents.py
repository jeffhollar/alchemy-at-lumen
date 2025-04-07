# Wikipedia Research Tool
#
# The primary function of agents is to leverage the GPT model to determine
# the subsequent action. Essentially, when faced with a problem, they outline
# the necessary steps to find a solution.
#
import os

from apikey import apikey
# from langchain.llms import OpenAI
from langchain_community.llms import OpenAI
from langchain.agents import load_tools, initialize_agent, AgentType

os.environ["OPENAI_API_KEY"] = apikey

# The temperature value influences the creativity of our model.
# A higher temperature setting results in more creativity, while a
# lower setting leads to more factual and objective outputs. A more
# factual response also tends to reduce potential inaccuracies or
# hallucinations from the model. The appropriate temperature setting
# depends on the application. For our current endeavor, which aims
# to generate creative article titles, a higher temperature is perferred.
#
# However, for tasks requiring factual accuracy, such as summarizing a legal
# document, a lower temperature might be more suitable.
#

# Having setup two LLMs, two prompt templates, and two chains, it's time to link them together.
# We'll establish an overarching chain using the SimpleSequentialChain with this code.
# - - - - - - - -- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
llm = OpenAI(temperature=0.0, max_tokens=50)

# Agents require access to specific tools, such as Google, or Wikipedia search capabilities.
#

# Provide Agent with two tools, Wikipedia and llm-math tools
#
tools = load_tools(
	['wikipedia', 'llm-math'],
	llm
)

# AgentType used is the most versatile action agent.
#
agent = initialize_agent(
	tools,
	llm,
	agent=AgentType.ZERO_SHOT_REACT_DESCRIPTION,
	verbose=True
)

prompt = input('Input Wikipedia Research Task\n')
agent.invoke(prompt)

# The agent accepts a user prompt in the Terminal, evaluates it and based on the tools
# provided determines the necessary actions to solve the problem.

# Test prompt used:
# In what year did the Titanic sink? How many years has it been since?
# - - - - - - - - - - - - - - - - - - -
